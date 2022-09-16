package lib

import (
	mysql "day-2/config"
	"day-2/models"
)

func CreateUser(user *models.Users) error {
	if err := mysql.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(id string, user *models.Users) error {
	if err := mysql.DB.Model(user).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *models.Users) error {
	if err := mysql.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUserById(id string) (models.Users, error) {
	var user models.Users
	result := mysql.DB.Where("id = ?", id).First(&user)
	return user, result.Error
}

func GetOneUserByEmail(email string) (models.Users, error) {
	var user models.Users
	result := mysql.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

func GetAllUser(q string, limit int, offset int) ([]models.Users, error) {
	var users []models.Users
	result := mysql.DB.Where("email LIKE ? OR name LIKE ?", "%"+q+"%", "%"+q+"%").Find(&users)

	return users, result.Error
}
