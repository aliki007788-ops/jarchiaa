package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleWithdraw(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{"message": "withdrawal requested"})
}


