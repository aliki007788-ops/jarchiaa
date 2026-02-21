package main

import (
"log"
"os"
"github.com/gin-gonic/gin"
)

func main() {
router := gin.Default()

text
router.GET("/v1/ads/nearby", handleNearby)
router.POST("/v1/ads", handleCreate)

port := os.Getenv("PORT")
if port == "" {
    port = "8083"
}

log.Printf("Ad service running on port %s", port)
router.Run(":" + port)
}


