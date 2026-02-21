package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleNearby(c *gin.Context) {
lat := c.Query("lat")
lng := c.Query("lng")

text
c.JSON(http.StatusOK, gin.H{
    "lat": lat,
    "lng": lng,
    "results": []string{},
})
}


