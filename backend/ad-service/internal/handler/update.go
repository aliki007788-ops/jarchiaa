package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleUpdate(c *gin.Context) {
id := c.Param("id")
c.JSON(http.StatusOK, gin.H{"message": "ad updated", "id": id})
}


