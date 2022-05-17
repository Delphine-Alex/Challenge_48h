package controllers

import (
	"GoChallenge48/Go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)
  ////////////////////
 //		GET		/////
////////////////////

// GetUsers handle /users (GET) - PRIVATE
func GetCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetCustomers())
}


// GetUserById handle /user/id (GET) - PUBLIC
func GetCustomerById(c *gin.Context) {
	if !isUserExistById(c.Params.ByName("id")) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found with this ID"})
		return
	}

	c.JSON(http.StatusOK, repository.GetCustomerById(c.Params.ByName("id")))
}


// isUserExistById return true if the user exist in the db
func isUserCustomerExistById(uuid string) bool {
	if repository.GetCustomerById(uuid).UUID == "" {
		return false
	} else {
		return true
	}
}

  ////////////////////////
 //		POST		/////
////////////////////////



