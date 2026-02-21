package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleStats(c *gin.Context) {
id := c.Param("id")
c.JSON(http.StatusOK, gin.H{
"ad_id": id,
"views": 1234,
"clicks": 234,
})
}


