package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/gingonic2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		log.Println(err)
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Delivery{})
	db.Migrator().CreateConstraint(&Delivery{}, "fk_delivery_order")
	db.Migrator().CreateConstraint(&User{}, "fk_user_order")
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&Category{})
	db.Migrator().CreateConstraint(&Category{}, "fk_category_product")
	db.Migrator().CreateConstraint(&User{}, "fk_user_product")
	db.AutoMigrate(&Product{})
	db.Migrator().CreateConstraint(&Order{}, "fk_order_orderdetail")
	db.Migrator().CreateConstraint(&User{}, "fk_user_orderdetail")
	db.AutoMigrate(&OrderDetail{})

	DB = db

}
