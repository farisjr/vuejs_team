package database

import (
	"vuejs/colabs/config"
	"vuejs/colabs/middlewares"
	"vuejs/colabs/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(users models.User) (interface{}, error) {

	if err := config.DB.Save(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func UpdateUser(users models.User, id int) (interface{}, error) {

	if err := config.DB.Find(&users, "id=?", id).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Save(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func DeleteUser(id int) (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users, "id=?", id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&users, "id=?", id).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id int) (interface{}, error) {
	var users models.User

	if err := config.DB.Find(&users, "id=?", id).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).Error; err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, err
}
