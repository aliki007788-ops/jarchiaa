package main

import (
"log"
"os"
"github.com/gin-gonic/gin"
)

func main() {
router := gin.Default()

text
router.GET("/v1/notifications", handleList)
router.POST("/v1/notifications", handleSend)

port := os.Getenv("PORT")
if port == "" {
    port = "8086"
}

log.Printf("Notification service running on port %s", port)
router.Run(":" + port)
}


