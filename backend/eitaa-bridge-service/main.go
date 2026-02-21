package main

import (
"log"
"os"
"github.com/gin-gonic/gin"
)

func main() {
router := gin.Default()

text
router.POST("/v1/eitaa/auth", handleAuth)
router.POST("/v1/eitaa/message", handleMessage)

port := os.Getenv("PORT")
if port == "" {
    port = "8088"
}

log.Printf("Eitaa bridge service running on port %s", port)
router.Run(":" + port)
}


