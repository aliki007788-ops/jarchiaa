package main

import (
"log"
"os"
"github.com/gin-gonic/gin"
)

func main() {
router := gin.Default()

text
router.GET("/v1/geo/nearby", handleNearby)

port := os.Getenv("PORT")
if port == "" {
    port = "8084"
}

log.Printf("Geo service running on port %s", port)
router.Run(":" + port)
}


