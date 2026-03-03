package app

import (
	"ReinaMusic/internal/config"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitReinaRest(config *config.ReinaConfig) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Pong!"})
	})

	if err := r.Run(config.Address + ":" + strconv.Itoa(config.Port)); err != nil {
		log.Fatalf("Failed to run Reina REST server!")
	}

	return nil
}
