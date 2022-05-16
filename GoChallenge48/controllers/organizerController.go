package controllers

import (
	"GoChallenge48/Go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)


// GetUsers handle /users (GET) - PRIVATE
func GetOrganizer(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetOrganizer())
}