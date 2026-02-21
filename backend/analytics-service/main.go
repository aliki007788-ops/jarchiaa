package main

import (
"log"
"os"
"github.com/gin-gonic/gin"
)

func main() {
router := gin.Default()

text
router.GET("/v1/analytics/dashboard", handleDashboard)

port := os.Getenv("PORT")
if port == "" {
    port = "8087"
}

log.Printf("Analytics service running on port %s", port)
router.Run(":" + port)
}


