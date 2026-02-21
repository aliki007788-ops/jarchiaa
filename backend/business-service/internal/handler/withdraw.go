package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleWithdraw(c *gin.Context) {
var req struct {
Amount int64 json:"amount"
}

text
if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}

c.JSON(http.StatusOK, gin.H{"message": "withdrawal requested"})
}


