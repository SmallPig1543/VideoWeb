package dao

import (
	"VideoWeb2/conf"
	"VideoWeb2/database/db/model"
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitMySQL() {
	dsn := conf.MySqlUser + ":" + conf.MySqlPassword + "@tcp(" + conf.MysqlIP + ")/" + conf.MySqlDataBase + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic(err)
	}
	DB = db
	_ = DB.AutoMigrate(&model.User{})
	_ = DB.AutoMigrate(&model.Message{})
	_ = DB.AutoMigrate(&model.Video{})
	_ = DB.AutoMigrate(&model.Comment{})
	_ = DB.AutoMigrate(&model.Reply{})

}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := DB
	return db.WithContext(ctx)
}
