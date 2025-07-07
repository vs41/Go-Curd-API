package main

import (
	dbconnection "curd/DBCONNECTION"
	"curd/helper"

	"github.com/gin-gonic/gin"
)

// func init() {
// 	err := godotenv.Load()

// }

func main() {

	// _, err := gorm.Open("postgres", "testDB")
	// if err != nil {
	// 	println(err)
	// 	panic(err)
	// }
	// DB.run()

	db, err := dbconnection.DBConnect()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&helper.Employ{})

	router := gin.Default()

	router.POST("/Create-User", helper.CreateUser)
	router.GET("/Get-User/:id", helper.GetUser)
	router.GET("/Get-User-List", helper.ListOfUser)
	router.PUT("Update-User/:id", helper.UpdateUser)
	router.Run(":8000")
}
