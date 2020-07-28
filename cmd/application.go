package cmd

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() error {
	MapUrls()
	if err := router.Run(":8081"); err != nil {
		return err
	}
	return nil
}
