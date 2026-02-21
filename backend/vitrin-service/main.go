package main

import (
"log"
"os"
"github.com/gin-gonic/gin"
)

func main() {
router := gin.Default()

text
router.GET("/v1/vitrin", handleList)
router.GET("/v1/vitrin/trending", handleTrending)

port := os.Getenv("PORT")
if port == "" {
    port = "8090"
}

log.Printf("Vitrin service running on port %s", port)
router.Run(":" + port)
}


