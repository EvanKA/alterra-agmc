package database

import (
	"day2/config"
	"day2/models"
)

func GetUsers() (interface{}, error) {
	var users []models.Users

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}
