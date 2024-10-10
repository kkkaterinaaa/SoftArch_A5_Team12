package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"twitter/user-service/initializers"
	"twitter/user-service/models"
)

func CreateUser(c *gin.Context) {
	var userData struct {
		Username string `json:"username"`
	}

	if err := c.ShouldBind(&userData); err != nil {
		logrus.Errorf("Username is not provided. Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	var user models.User
	if err := initializers.DB.Where("username = ?", userData.Username).First(&user).Error; err == nil {
		logrus.Infof("Username already exists. Username: %s, Error: %v", userData.Username, err)
		logrus.Infof("Logged user. Username: %s, Id: %s", user.Username, user.ID)
		c.JSON(http.StatusCreated, gin.H{"userId": user.ID})
		return
	}

	newUser := models.User{
		Username: userData.Username,
	}

	if err := initializers.DB.Create(&newUser).Error; err != nil {
		logrus.Errorf("Failed to register new user. Username: %s, Error: %v", newUser.Username, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register new user"})
		return
	}

	logrus.Infof("Registered new user. Username: %s, Id: %s", newUser.Username, user.ID)
	c.JSON(http.StatusCreated, gin.H{"userId": newUser.ID})
}

func GetUser(c *gin.Context) {
	userId := c.Param("id")
	userID, _ := strconv.Atoi(userId)
	var user models.User
	if err := initializers.DB.Where("id = ?", uint(userID)).First(&user).Error; err != nil {
		logrus.Errorf("User not found. ID: %s, error: %v", userID, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	logrus.Infof("Retrieved user. Username: %s", user.Username)
	c.JSON(http.StatusOK, user)
}
