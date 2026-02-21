package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleDelete(c *gin.Context) {
id := c.Param("id")
c.JSON(http.StatusOK, gin.H{"message": "ad deleted", "id": id})
}


