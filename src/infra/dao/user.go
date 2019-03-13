package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/snowman-mh/go-sample/src/config"
	"github.com/snowman-mh/go-sample/src/domain/model"
	"github.com/snowman-mh/go-sample/src/domain/repository"
	"github.com/snowman-mh/go-sample/src/infra/entity"
)

type userImpl struct {
	db *gorm.DB
}

func NewUser() repository.User {
	db, err := gorm.Open(config.Get().DB.Driver, config.Get().DB.DSN())
	if err != nil {
		panic(err)
	}
	return userImpl{db: db}
}

func (user userImpl) Add(userModel *model.User) error {
	userEntity := entity.NewUser(userModel)
	if err := user.db.Create(&userEntity).Error; err != nil {
		return err
	}
	userEntity.Write(userModel)
	return nil
}

func (user userImpl) Fetch(userModel *model.User) error {
	userEntity := entity.NewUser(userModel)
	if err := user.db.Find(&userEntity).Error; err != nil {
		return err
	}
	userEntity.Write(userModel)
	return nil
}
