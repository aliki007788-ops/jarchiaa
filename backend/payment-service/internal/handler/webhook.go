package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleWebhook(c *gin.Context) {
var data map[string]interface{}
c.ShouldBindJSON(&data)

text
c.JSON(http.StatusOK, gin.H{"status": "received"})
}


