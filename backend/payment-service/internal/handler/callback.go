package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleCallback(c *gin.Context) {
authority := c.Query("Authority")
status := c.Query("Status")

text
c.JSON(http.StatusOK, gin.H{
    "authority": authority,
    "status": status,
})
}


