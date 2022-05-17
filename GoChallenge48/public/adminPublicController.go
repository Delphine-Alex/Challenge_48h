package public

import (
	//"GoChallenge48/Go/models"
	"GoChallenge48/Go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNewsletters handle /newsletters (GET)
func GetAdmin(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetAdmin())
}
