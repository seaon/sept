package model

import (
	"backend/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

type PageInfo struct {
	Page     int
	PageSize int
}

func init() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetDatabase().User,
		config.GetDatabase().Password,
		config.GetDatabase().Host,
		config.GetDatabase().Dbname)

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	DB = db

	// 修改 gorm 的数据表名默认组成
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.GetDatabase().TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
}
