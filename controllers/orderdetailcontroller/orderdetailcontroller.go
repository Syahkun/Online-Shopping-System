package orderdetailcontroller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"main.go/models"
)

func Index(c *gin.Context) {

	var orderdetails []models.OrderDetail

	if err := models.DB.Preload("Order.User").
		Preload("Order.Delivery").Preload("Product.User").
		Preload("Product.Category").Preload(clause.Associations).
		Find(&orderdetails).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"orderdetails": orderdetails})
		return
	}

}

func Show(c *gin.Context) {
	var orderdetail models.OrderDetail
	id := c.Param("id")

	if err := models.DB.First(&orderdetail, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"orderdetail": orderdetail})
}

func Create(c *gin.Context) {

	var orderdetail models.OrderDetail
	var product models.Product

	if err := c.ShouldBindJSON(&orderdetail); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	log.Println("===========================")
	log.Println(orderdetail.ID)
	models.DB.First(&product, orderdetail.ProductID)
	totalprice := int(product.Price) * orderdetail.Amount
	log.Println(totalprice)
	orderdetail.Price = totalprice
	log.Println(orderdetail.Price)

	models.DB.Create(&orderdetail)
	c.JSON(http.StatusOK, gin.H{"orderdetail": orderdetail})
}

func Update(c *gin.Context) {
	var orderdetail models.OrderDetail
	id := c.Param("id")

	if err := c.ShouldBindJSON(&orderdetail); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&orderdetail).Where("id = ?", id).Updates(&orderdetail).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot Update data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}

func Delete(c *gin.Context) {
	var orderdetail models.OrderDetail
	id := c.Param("id")
	if err := models.DB.First(&orderdetail, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
	} else if models.DB.Delete(&orderdetail, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "request failed"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Deleted"})
	}
}
