package logger

import (
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"goblog/pkg/config"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	channels   = make(map[string]*zap.Logger)
	asyncWriters []*asyncWriter
	initialized  bool
	mu           sync.RWMutex
)

var channelNames = []string{"app", "error", "sql", "curl"}

// Initialize 根据配置初始化所有日志 channel
func Initialize() {
	mu.Lock()
	defer mu.Unlock()

	if initialized {
		return
	}

	storage := config.GetString("log.storage", "storage/logs")
	bufferSize := config.GetInt("log.async_buffer_size", 1024)
	debug := config.GetBool("app.debug")

	for _, name := range channelNames {
		prefix := "log.channels." + name
		if !config.Viper.IsSet(prefix + ".path") {
			continue
		}

		channels[name] = newChannelLogger(channelOptions{
			name:       name,
			storage:    storage,
			subPath:    config.GetString(prefix + ".path"),
			level:      config.GetString(prefix + ".level", "info"),
			maxAge:     config.GetInt(prefix + ".max_age", 30),
			async:      config.GetBool(prefix + ".async"),
			bufferSize: bufferSize,
			debug:      debug,
		})
	}

	if _, ok := channels["error"]; !ok {
		channels["error"] = zap.NewNop()
	}

	initialized = true
}

type channelOptions struct {
	name       string
	storage    string
	subPath    string
	level      string
	maxAge     int
	async      bool
	bufferSize int
	debug      bool
}

func newChannelLogger(opts channelOptions) *zap.Logger {
	dir := opts.storage + "/" + opts.subPath
	_ = os.MkdirAll(dir, 0755)

	pattern := dir + "/" + opts.name + "-%Y-%m-%d.log"
	writer, err := rotatelogs.New(
		pattern,
		rotatelogs.WithMaxAge(time.Duration(opts.maxAge)*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic("logger: create rotatelogs for " + opts.name + ": " + err.Error())
	}

	var output zapcore.WriteSyncer
	if opts.async {
		aw := newAsyncWriter(writer, opts.bufferSize)
		asyncWriters = append(asyncWriters, aw)
		output = aw
	} else {
		output = zapcore.AddSync(writer)
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	cores := []zapcore.Core{
		zapcore.NewCore(encoder, output, parseLevel(opts.level)),
	}

	if opts.debug {
		consoleCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			zapcore.AddSync(os.Stdout),
			parseLevel(opts.level),
		)
		cores = append(cores, consoleCore)
	}

	return zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(1))
}

func parseLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn", "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

// Channel 获取指定 channel 的 logger
func Channel(name string) *zap.Logger {
	mu.RLock()
	defer mu.RUnlock()

	if lg, ok := channels[name]; ok {
		return lg
	}
	return zap.NewNop()
}

// LogError 记录项目错误（同步写入 error channel）
func LogError(err error) {
	if err == nil {
		return
	}
	Channel("error").Error(err.Error(), zap.Error(err))
}

// LogCurl 记录 HTTP 请求日志（异步写入 curl channel）
func LogCurl(fields ...zap.Field) {
	Channel("curl").Info("http request", fields...)
}

// Sync 刷盘所有日志，程序退出前调用
func Sync() {
	mu.RLock()
	defer mu.RUnlock()

	for _, aw := range asyncWriters {
		aw.close()
	}

	for _, lg := range channels {
		_ = lg.Sync()
	}
}

// asyncWriter 异步日志写入器
type asyncWriter struct {
	ch     chan []byte
	writer io.Writer
	wg     sync.WaitGroup
}

func newAsyncWriter(w io.Writer, bufferSize int) *asyncWriter {
	aw := &asyncWriter{
		ch:     make(chan []byte, bufferSize),
		writer: w,
	}
	aw.wg.Add(1)
	go aw.run()
	return aw
}

func (aw *asyncWriter) run() {
	defer aw.wg.Done()
	for p := range aw.ch {
		_, _ = aw.writer.Write(p)
	}
}

func (aw *asyncWriter) Write(p []byte) (int, error) {
	buf := make([]byte, len(p))
	copy(buf, p)

	select {
	case aw.ch <- buf:
		return len(p), nil
	default:
		_, err := aw.writer.Write(p)
		return len(p), err
	}
}

func (aw *asyncWriter) Sync() error {
	return nil
}

func (aw *asyncWriter) close() {
	close(aw.ch)
	aw.wg.Wait()
}
