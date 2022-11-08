package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keerthi/golang-jwt--project/database"
	"github.com/keerthi/golang-jwt--project/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func New() *UserRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.User{})
	return &UserRepo{DB: db}
}

// create user
func (userRepo *UserRepo) CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.CreateUser(userRepo.DB, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
