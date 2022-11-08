package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keerthi/golang-jwt--project/controllers"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	public := r.Group("/api")

	userRepo := controllers.New()

	public.POST("/user", userRepo.CreateUser)
	return r
}
