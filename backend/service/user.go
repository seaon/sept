package service

import (
	"backend/model"
	"backend/utils"
	"errors"
	"gorm.io/gorm"
)

func Register(u model.User) (e error, userInter model.User) {
	var user model.User
	if !errors.Is(model.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), user
	}

	// 附加uuid 密码md5简单加密 注册
	u.Salt = utils.RandomString(6)
	u.Password = utils.Mdv(u.Password, u.Salt)

	err := model.DB.Create(&u).Error
	return err, u
}

func Login(u model.User) (err error, userInter model.User) {
	var user model.User
	err = model.DB.Where("username = ?", u.Username).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户不存在"), user
	}

	if err != nil {
		return err, user
	}

	u.Password = utils.Mdv(u.Password, user.Salt)

	if u.Password != user.Password {
		return errors.New("密码错误"), user
	}

	return err, user
}
