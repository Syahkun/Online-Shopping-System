package main

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers/authcontroller"
	"main.go/controllers/categorycontroller"
	"main.go/controllers/deliverycontroller"
	"main.go/controllers/ordercontroller"
	"main.go/controllers/orderdetailcontroller"
	"main.go/controllers/productcontroller"
	"main.go/controllers/usercontroller"
	"main.go/middlewares"
	"main.go/models"
)

func main() {
	r := gin.Default()
	models.Conn()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	public := r.Group("/api")
	public.POST("/user", usercontroller.Create)
	public.GET("/login", authcontroller.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/users", usercontroller.Index)
	protected.GET("/api/user/:id", usercontroller.Show)
	protected.PUT("/api/user/:id", usercontroller.Update)
	protected.DELETE("/api/user/:id", usercontroller.Delete)
	r.GET("/api/deliveries", deliverycontroller.Index)
	r.GET("/api/delivery/:id", deliverycontroller.Show)
	r.POST("/api/delivery", deliverycontroller.Create)
	r.PUT("/api/delivery/:id", deliverycontroller.Update)
	r.DELETE("/api/delivery/:id", deliverycontroller.Delete)
	r.GET("api/orders", ordercontroller.GetAll)
	r.POST("api/order", ordercontroller.Create)
	r.GET("api/order/:id", ordercontroller.Show)
	r.PUT("api/order/:id", ordercontroller.Update)
	r.DELETE("api/order/:id", ordercontroller.Delete)
	r.GET("/api/categories", categorycontroller.Index)
	r.POST("api/category", categorycontroller.Create)
	r.GET("api/category/:id", categorycontroller.Show)
	r.PUT("api/category/:id", categorycontroller.Update)
	r.DELETE("api/category/:id", categorycontroller.Delete)
	r.GET("/api/products", productcontroller.Index)
	r.POST("api/product", productcontroller.Create)
	r.GET("api/product/:id", productcontroller.Show)
	r.PUT("api/product/:id", productcontroller.Update)
	r.DELETE("api/product/:id", productcontroller.Delete)
	r.GET("/api/orderdetails", orderdetailcontroller.Index)
	r.POST("api/orderdetail", orderdetailcontroller.Create)
	r.GET("api/orderdetail/:id", orderdetailcontroller.Show)
	r.PUT("api/orderdetail/:id", orderdetailcontroller.Update)
	r.DELETE("api/orderdetail/:id", orderdetailcontroller.Delete)
	// r.GET("/api/login", authcontroller.Login)
	r.Run() // listen and serve on 0.0.0.0:8080
}
