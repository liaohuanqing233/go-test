package httpclient

import (
	"io"
	"net/http"
	"time"

	"goblog/pkg/logger"

	"go.uber.org/zap"
)

// Client 带 curl 异步日志的 HTTP 客户端
type Client struct {
	http *http.Client
}

// Default 默认客户端
var Default = &Client{
	http: &http.Client{Timeout: 30 * time.Second},
}

// Do 执行请求并异步记录 curl 日志
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := c.http.Do(req)

	fields := []zap.Field{
		zap.String("method", req.Method),
		zap.String("url", req.URL.String()),
		zap.Duration("elapsed", time.Since(start)),
	}

	if err != nil {
		fields = append(fields, zap.Error(err))
		logger.LogCurl(fields...)
		return resp, err
	}

	fields = append(fields, zap.Int("status", resp.StatusCode))
	logger.LogCurl(fields...)
	return resp, err
}

// Get 发起 GET 请求
func (c *Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

// Post 发起 POST 请求
func (c *Client) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req)
}
