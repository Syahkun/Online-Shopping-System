package usercontroller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"main.go/functions"
	"main.go/models"
)

func Index(c *gin.Context) {

	var users []models.User

	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})

}

func Show(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := models.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

type ResponCreateUser struct {
	ID          uint
	Name        string
	Username    string
	Email       *string
	Address     string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func Create(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	user.Password = string(hashedPassword)
	models.DB.Create(&user)

	var ResUser ResponCreateUser

	ResUser.Name = user.Name
	ResUser.Email = user.Email
	ResUser.Address = user.Address
	ResUser.Username = user.Username
	ResUser.PhoneNumber = user.PhoneNumber

	c.JSON(http.StatusOK, gin.H{"user": ResUser})
}

func Update(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot Update data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}

func Delete(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	exist, _, message, isError := functions.GetDataByID(id)
	if !exist {
		if isError {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": message})
			return
		} else if !isError {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": message})
			return
		}
	}
	if models.DB.Delete(&user, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "request failed"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Deleted"})
	}
}
