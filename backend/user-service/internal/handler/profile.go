package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{
"id": "123",
"name": "??? ?????",
"phone": "09123456789",
})
}


