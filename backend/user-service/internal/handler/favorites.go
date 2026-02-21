package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func GetFavorites(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{"favorites": []string{}})
}


