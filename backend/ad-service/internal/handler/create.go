package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleCreate(c *gin.Context) {
var req struct {
Title string json:"title"
}

text
if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}

c.JSON(http.StatusOK, gin.H{"message": "ad created"})
}


