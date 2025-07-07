package helper

import (
	dbconnection "curd/DBCONNECTION"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var req Employ

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := dbconnection.DBConnect()
	if err != nil {
		panic(err)
	}
	//defer db.close()

	if err := db.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"StatusOk": 200})
	}

}

func GetUser(c *gin.Context) {
	db, err := dbconnection.DBConnect()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	id := c.Param("id")

	fmt.Println("id ******", id)

	var employ Employ
	if err := db.First(&employ, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"Error": "User not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
		}
		return
	}
	c.JSON(http.StatusFound, employ)
}

func ListOfUser(c *gin.Context) {
	var UserList []Employ
	db, err := dbconnection.DBConnect()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	if err := db.Find(&UserList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No records",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
		}
	}

	c.JSON(http.StatusFound, UserList)

}

func UpdateUser(c *gin.Context) {
	var PreData Employ
	id := c.Param("id")
	db, err := dbconnection.DBConnect()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	if err := db.Find(&PreData, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No records",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
		}
	}

	var InputData Employ

	if err := c.ShouldBindJSON(&InputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	db.Model(&PreData).Updates(InputData)

	if err := db.Model(&PreData).Updates(InputData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No records",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
		}
	}
	c.JSON(http.StatusOK, PreData)

}
