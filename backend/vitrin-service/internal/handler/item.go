package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func GetItem(c *gin.Context) {
id := c.Param("id")
c.JSON(http.StatusOK, gin.H{
"id": id,
"title": "???? ?????",
"views": 1234,
})
}


