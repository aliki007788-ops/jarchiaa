package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleEitaa(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{"message": "eitaa message sent"})
}


