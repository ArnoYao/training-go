package dao

import (
	"fmt"

	model "github.com/arnoyao/training-go/02week/model"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func (db *UserDB) Get(id int64) (*model.User, error) {
	var u model.User

	err := db.DB.Where("id = ?", id).First(&u).Error
	if err != nil {
		return nil, fmt.Errorf("get user by id %v : %w", id, err)
	}
	return &u, nil
}
