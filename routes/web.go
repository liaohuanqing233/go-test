package routes

import (
	"goblog/app/http/controllers"
	"goblog/app/http/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterWebRoutes 注册web路由
func RegisterWebRoutes(r *mux.Router) {
	pages := new(controllers.PagesController)
	// 静态页面
	r.HandleFunc("/about", pages.About).Methods("GET").Name("about")
	r.NotFoundHandler = http.HandlerFunc(pages.NotFound)

	// 文章路由
	articles := new(controllers.ArticlesController)
	r.HandleFunc("/", articles.Index).Methods("GET").Name("home")
	r.HandleFunc("/articles", articles.Index).Methods("GET").Name("articles.index")
	r.HandleFunc("/articles/{id:[0-9]+}", articles.Show).Methods("GET").Name("articles.show")
	r.HandleFunc("/articles/create", middlewares.Auth(articles.Create)).Methods("GET").Name("articles.create")
	r.HandleFunc("/articles", middlewares.Auth(articles.Store)).Methods("POST").Name("articles.store")
	r.HandleFunc("/articles/{id:[0-9]+}/edit", middlewares.Auth(articles.Edit)).Methods("GET").Name("articles.edit")
	r.HandleFunc("/articles/{id:[0-9]+}", middlewares.Auth(articles.Update)).Methods("POST").Name("articles.update")
	r.HandleFunc("/articles/{id:[0-9]+}/delete", middlewares.Auth(articles.Delete)).Methods("POST").Name("articles.delete")

	// 用戶認證
	auc := new(controllers.AuthController)
	r.HandleFunc("/auth/register", middlewares.Guest(auc.Register)).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", middlewares.Guest(auc.DoRegister)).Methods("POST").Name("auth.doregister")
	r.HandleFunc("/auth/login", middlewares.Guest(auc.Login)).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/do-login", middlewares.Guest(auc.DoLogin)).Methods("POST").Name("auth.dologin")
	r.HandleFunc("/auth/logout", middlewares.Auth(auc.Logout)).Methods("POST").Name("auth.logout")

	uc := new(controllers.UserController)
	r.HandleFunc("/users/{id:[0-9]+}", uc.Show).Methods("GET").Name("users.show")

	//静态资源路由
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	// 中间件
	// r.Use(middlewares.ForceHTML)

	//开启会话
	r.Use(middlewares.StartSession)
}
