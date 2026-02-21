package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func VerifyBusiness(c *gin.Context) {
id := c.Param("id")
var req struct {
Verified bool json:"verified"
}

text
if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}

c.JSON(http.StatusOK, gin.H{"message": "business verified", "id": id})
}


