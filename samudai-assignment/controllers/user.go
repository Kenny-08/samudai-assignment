package controllers

import (
	"fmt"
	"login-signup-api/config"
	"login-signup-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct{
	Name string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func GetUsers(c *gin.Context){
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func Register(c *gin.Context){
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	user := models.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	
	_, err := user.SaveUser()
	fmt.Println("saved user")
	if err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"registration success"})

	// c.BindJSON(&user)
	// config.DB.Create(&user)
	// c.JSON(200, &user)
}

func DeleteUser(c *gin.Context){
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context){
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Find(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}