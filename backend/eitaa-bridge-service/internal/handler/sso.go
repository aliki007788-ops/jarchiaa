package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleSSO(c *gin.Context) {
var req struct {
Code string json:"code"
}

text
if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}

c.JSON(http.StatusOK, gin.H{
    "token": "jwt-token",
    "user": map[string]interface{}{
        "id": "123",
        "name": "????? ????",
    },
})
}


