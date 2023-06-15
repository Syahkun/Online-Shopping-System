package deliverycontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/models"
)

func Index(c *gin.Context) {

	var deliveries []models.Delivery

	models.DB.Find(&deliveries)
	c.JSON(http.StatusOK, gin.H{"deliverys": deliveries})

}

func Show(c *gin.Context) {
	var delivery models.Delivery
	id := c.Param("id")

	if err := models.DB.First(&delivery, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"delivery": delivery})
}

func Create(c *gin.Context) {

	var delivery models.Delivery

	if err := c.ShouldBindJSON(&delivery); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&delivery)
	c.JSON(http.StatusOK, gin.H{"delivery": delivery})
}

func Update(c *gin.Context) {
	var delivery models.Delivery
	id := c.Param("id")

	if err := c.ShouldBindJSON(&delivery); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&delivery).Where("id = ?", id).Updates(&delivery).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot Update data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}

func Delete(c *gin.Context) {
	var delivery models.Delivery
	id := c.Param("id")
	if err := models.DB.First(&delivery, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
	} else if models.DB.Delete(&delivery, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "request failed"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Deleted"})
	}
}
