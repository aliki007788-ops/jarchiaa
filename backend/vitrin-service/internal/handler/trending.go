package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleTrending(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{
"items": []map[string]interface{}{
{"id": "1", "title": "????? ?"},
},
})
}


