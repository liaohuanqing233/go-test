package article

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
)

// Create 保存文章到数据库
func (article *Article) Create() (err error) {
	if err = model.DB.Create(&article).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

// Update 更新文章
func (article *Article) Update() (rowsAffected int64, err error) {
	reuslt := model.DB.Save(&article)
	if err = reuslt.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return reuslt.RowsAffected, nil
}

// Delete 删除文章
func (article Article) Delete() (rowsAffected int64, err error) {
	reuslt := model.DB.Delete(&article)
	if err = reuslt.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return reuslt.RowsAffected, nil
}
