package categorycontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/models"
)

func Index(c *gin.Context) {

	var categories []models.Category

	models.DB.Find(&categories)
	c.JSON(http.StatusOK, gin.H{"categorys": categories})

}

func Show(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := models.DB.First(&category, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"category": category})
}

func Create(c *gin.Context) {

	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&category)
	c.JSON(http.StatusOK, gin.H{"category": category})
}

func Update(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&category).Where("id = ?", id).Updates(&category).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot Update data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}

func Delete(c *gin.Context) {
	var category models.Category
	id := c.Param("id")
	if err := models.DB.First(&category, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
	} else if models.DB.Delete(&category, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "request failed"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Deleted"})
	}
}
