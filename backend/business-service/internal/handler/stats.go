package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func GetStats(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{
"views": 1234,
"clicks": 234,
"likes": 45,
})
}


