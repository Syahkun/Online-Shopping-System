package ordercontroller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/models"
)

func GetAll(c *gin.Context) {
	var orders []models.Order
	log.Println(models.DB.Joins("left join customers on customers.id = orders.customer_id").Find(&orders))
	if err := models.DB.Preload("User").Preload("Delivery").Find(&orders).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"orders": orders})
		return
	}

}

func Create(c *gin.Context) {

	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Preload("User").Preload("Delivery").Create(&order)
	c.JSON(http.StatusOK, gin.H{"order": order})
}

func Show(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := models.DB.Preload("User").Preload("Delivery").First(&order, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func Delete(c *gin.Context) {
	var order models.Order
	id := c.Param("id")
	if err := models.DB.First(&order, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
	} else if models.DB.Delete(&order, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "request failed"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Deleted"})
	}
}

func Update(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&order).Where("id = ?", id).Updates(&order).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot Update data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}
