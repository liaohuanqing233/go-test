package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/policies"
	"goblog/app/requests"
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"

	"gorm.io/gorm"
)

// ArticlesController 文章控制器
type ArticlesController struct {
}

// Index 文章列表首頁
func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	articles, err := article.GetAll()

	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 系统错误")
	} else {
		view.Render(w, view.D{
			"Articles": articles,
		}, "articles.index", "articles._article_meta")
	}
}

// Show 显示文章
func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	article, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 系统错误")
		}
	} else {
		view.Render(w, view.D{
			"Article":          article,
			"CanModifyArticle": policies.CanModifyArticle(article),
		}, "articles.show", "articles._article_meta")
	}
}

// Create 新建文章
func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "articles.create", "articles._form_field")
}

// Store 保存文章
func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	// Form会比PostForm多一些URL中自带的参数
	currentUser := auth.User()
	_article := article.Article{
		Title:  r.PostFormValue("title"),
		Body:   r.PostFormValue("body"),
		UserID: currentUser.ID,
	}
	errors := requests.ValidateArticleForm(_article)

	// 检查是否有错误
	if len(errors) == 0 {
		_article.Create()
		if _article.ID > 0 {
			showURL := route.NameToURL("articles.show", "id", _article.GetStringID())
			http.Redirect(w, r, showURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器错误")
		}
	} else {
		view.Render(w, view.D{
			"Article": _article,
			"Errors":  errors,
		}, "articles.create", "articles._form_field")
	}
}

// Edit 编辑文章
func (*ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_article, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 系统错误")
		}
	} else {
		// 检查权限
		if !policies.CanModifyArticle(_article) {
			flash.Warning("未授权操作！")
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			view.Render(w, view.D{
				"Article": _article,
				"Errors":  view.D{},
			}, "articles.edit", "articles._form_field")
		}

	}
}

// Update 文章更新路由
func (*ArticlesController) Update(w http.ResponseWriter, r *http.Request) {
	// 1 获取路由
	id := route.GetRouteVariable("id", r)

	// 2 获取文章判断是否存在即可
	_article, err := article.Get(id)

	// 3.检查是否有错误
	if err != nil {
		// 3.1 错误处理
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 系统错误")
		}
	} else {
		if !policies.CanModifyArticle(_article) {
			flash.Warning("未授权操作！")
			http.Redirect(w, r, "/", http.StatusForbidden)
		} else {
			// 3.2 获取POST参数
			_article.Title = r.PostFormValue("title")
			_article.Body = r.PostFormValue("body")
			errors := requests.ValidateArticleForm(_article)

			// 3.3 不存在错误进行保存错误返回上一个页面继续进行编辑操作
			if len(errors) == 0 {
				rowsAffected, err := _article.Update()
				if err != nil {
					logger.LogError(err)
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprint(w, "500 系统错误")
					return
				}
				if rowsAffected > 0 {
					showURL := route.NameToURL("articles.show", "id", id)
					http.Redirect(w, r, showURL, http.StatusFound)
				} else {
					fmt.Fprint(w, "您没有做任何更改")
				}
			} else {
				view.Render(w, view.D{
					"Article": _article,
					"Errors":  errors,
				}, "articles.edit", "articles._form_field")
			}
		}

	}
}

// Delete 文章删除路由
func (*ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {
	// 1.获取文章
	id := route.GetRouteVariable("id", r)
	_article, err := article.Get(id)

	// 2.排查错误
	if err != nil {
		logger.LogError(err)
		// 2.1存在错误进行提示
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 系统错误")
		}
	} else {
		// 检查权限
		if !policies.CanModifyArticle(_article) {
			flash.Warning("您没有权限执行此操作！")
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			// 2.2不存在错误进行删除
			rowsAffected, err := _article.Delete()
			// 2.3删除出现错误进行处理未发现错误调整首页
			if err != nil {
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 系统错误")
			} else {
				if rowsAffected > 0 {
					indexURL := route.NameToURL("articles.index")
					http.Redirect(w, r, indexURL, http.StatusFound)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprint(w, "404 文章未找到")
				}
			}
		}

	}
}
