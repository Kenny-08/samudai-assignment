package controllers

import (
	"fmt"
	"login-signup-api/config"
	"login-signup-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ManageAccessInput struct {
	AdminID    int `json:"admin_id"`
	UserID int `json:"user_id"`
	Role string	`json:"role"`
}

func GrantRole(c *gin.Context){
	var input ManageAccessInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	

	var access models.ManageAccess
	access.UserID = input.UserID
	access.AdminID = input.AdminID
	access.Role = input.Role

	_, err2 := access.SaveAccess()

	models.UpdateUserRole(access.UserID, access.Role)

	if err2 != nil {
		c.JSON(400, gin.H{"error": err2.Error()})
		return
	}

	c.JSON(200, gin.H{"data": input})
} 

func ChangeRole(c *gin.Context){
	var access models.ManageAccess
	config.DB.Where("id = ?", c.Param("id")).Find(&access)
	c.BindJSON(&access)
	config.DB.Save(&access)
	c.JSON(200, &access)

}

