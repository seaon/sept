package service

import (
	"backend/model"
	"errors"
)

func Add(article model.Article) (err error, id int) {
	res := model.DB.Create(&article)

	if res.Error != nil {
		return errors.New("创建失败"), 0
	}

	return err, article.ID
}
