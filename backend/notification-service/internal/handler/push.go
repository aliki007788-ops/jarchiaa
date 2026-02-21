package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandlePush(c *gin.Context) {
var req struct {
UserID string json:"user_id"
Title string json:"title"
Body string json:"body"
}

text
if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}

c.JSON(http.StatusOK, gin.H{"message": "notification sent"})
}


