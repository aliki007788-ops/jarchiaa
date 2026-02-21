package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleRegister(c *gin.Context) {
var req struct {
Name string json:"name"
}

text
if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}

c.JSON(http.StatusOK, gin.H{"message": "business registered"})
}


