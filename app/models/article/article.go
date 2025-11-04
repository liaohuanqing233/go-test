package article

import (
	"goblog/app/models/user"
	"goblog/pkg/model"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"strconv"
)

// Article 文章模型
type Article struct {
	model.BaseModel
	Title  string
	Body   string
	UserID uint64 `gorm:"not null;index"`
	User   user.User
}

// Get 根据ID获取单篇文章
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)
	if err := model.DB.Preload("User").First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

// GetAll 獲取所有文章
func GetAll() ([]Article, error) {
	var articles []Article

	if err := model.DB.Preload("User").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

// GetByUserID 获取全部文章
func GetByUserID(uid string) ([]Article, error) {
	var articles []Article
	if err := model.DB.Where("user_id = ?", uid).Preload("User").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

// Link 文章访问连接
func (a Article) Link() string {
	return route.NameToURL("articles.show", "id", strconv.FormatUint(a.ID, 10))
}

// CreatedAtDate 创建日期
func (a Article) CreatedAtDate() string {
	return a.CreatedAt.Format("2006-01-02")
}
