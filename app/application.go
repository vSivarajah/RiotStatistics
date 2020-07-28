package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	MapUrls()
	log.Printf("Starting application server...")
	router.Run(":8081")
}
