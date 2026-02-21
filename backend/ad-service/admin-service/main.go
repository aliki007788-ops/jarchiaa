package main

import (
"log"
"os"
"github.com/gin-gonic/gin"
)

func main() {
router := gin.Default()

text
router.GET("/v1/admin/stats", handleStats)
router.GET("/v1/admin/modules", handleModules)

port := os.Getenv("PORT")
if port == "" {
    port = "8089"
}

log.Printf("Admin service running on port %s", port)
router.Run(":" + port)
}


