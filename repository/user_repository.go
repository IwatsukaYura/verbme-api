package repository

import (
	"verbme-api/model"

	"gorm.io/gorm"
)

//DBと直接のやり取りを行うインターフェース
type IUserRepository interface {
	GetUserByName(user *model.User, email string) error
	CreateUser(user *model.User) error
}


type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}


//インターフェースを満たすためのメソッド
func (ur *userRepository) GetUserByName(user *model.User, name string) error {
	if err := ur.db.Where("name = ?", name).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
