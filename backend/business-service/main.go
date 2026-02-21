package main

import (
"log"
"os"
"github.com/gin-gonic/gin"
)

func main() {
router := gin.Default()

text
router.POST("/v1/business/register", handleRegister)
router.GET("/v1/business/profile", handleGetProfile)

port := os.Getenv("PORT")
if port == "" {
    port = "8082"
}

log.Printf("Business service running on port %s", port)
router.Run(":" + port)
}


