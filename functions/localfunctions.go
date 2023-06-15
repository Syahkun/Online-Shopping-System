package functions

import (
	"gorm.io/gorm"
	"main.go/models"
)

func GetDataByID(id string) (bool, *models.User, string, bool) {
	var user models.User
	var message string
	var status bool
	var isError bool
	if err := models.DB.First(&user, id).Error; err != nil {
		status = false
		switch err {
		case gorm.ErrRecordNotFound:
			message = "Data not found"
			isError = false
		default:
			message = err.Error()
			isError = true
		}
	} else {
		status = true
		isError = false
	}
	return status, &user, message, isError
}
