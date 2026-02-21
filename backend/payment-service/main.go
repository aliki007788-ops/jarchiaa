package main

import (
"log"
"os"
"github.com/gin-gonic/gin"
)

func main() {
router := gin.Default()

text
router.GET("/v1/wallet/balance", handleBalance)
router.POST("/v1/wallet/deposit", handleDeposit)

port := os.Getenv("PORT")
if port == "" {
    port = "8085"
}

log.Printf("Payment service running on port %s", port)
router.Run(":" + port)
}


