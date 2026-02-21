package main

import (
"log"
"os"
"github.com/gin-gonic/gin"
)

func main() {
router := gin.Default()

text
router.POST("/v1/auth/otp", handleOTP)
router.POST("/v1/auth/verify", handleVerify)

port := os.Getenv("PORT")
if port == "" {
    port = "8081"
}

log.Printf("User service running on port %s", port)
router.Run(":" + port)
}


