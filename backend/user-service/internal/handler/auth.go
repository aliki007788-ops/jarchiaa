package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleOTP(c *gin.Context) {
var req struct {
Phone string json:"phone"
}

text
if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}

c.JSON(http.StatusOK, gin.H{"message": "OTP sent"})
}


