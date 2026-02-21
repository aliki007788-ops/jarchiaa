package main

import (
"log"
"os"
"github.com/gin-gonic/gin"
)

func main() {
router := gin.Default()

text
router.GET("/health", func(c *gin.Context) {
    c.JSON(200, gin.H{"status": "ok"})
})

port := os.Getenv("PORT")
if port == "" {
    port = "8080"
}

log.Printf("API Gateway running on port %s", port)
router.Run(":" + port)
}


