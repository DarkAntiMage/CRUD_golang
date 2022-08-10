package model

import (
	"echofw/config"
)

type (
	Users struct {
		ID    int    `json:"id" gorm:"primaryKey"`
		Name  string `json:"nama"`
		Email string `json:"email"`
	}
)

func (user *Users) CreateUser() error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUser(id int) (Users, error) {
	var user Users
	result := config.DB.Where("id = ?", id).First(&user)
	return user, result.Error
}

func (user *Users) UpdateUser(id int) error {
	if err := config.DB.Model(&Users{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) DeleteUser() error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUsers() ([]Users, error) {
	var users []Users
	result := config.DB.Find(&users)

	return users, result.Error
}
