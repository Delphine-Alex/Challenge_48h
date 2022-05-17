package controllers

import (
	"GoChallenge48/Go/repository"
	"GoChallenge48/Go/models"
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

// PostUser handle /user for creating a new user (POST) - PUBLIC
func PostCustomer(c *gin.Context) {
	// Validate input
	var input models.CreateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(repository.GetUserByEmail(input.Email)) != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this mail already exist"})
		return
	}

	if len(repository.GetUserByUsername(input.Username)) != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "This username is already taken by another user"})
		return
	}

	repository.PostUser(input)

	c.JSON(http.StatusCreated, input)
}

