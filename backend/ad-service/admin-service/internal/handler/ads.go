package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func GetPendingAds(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{
"ads": []map[string]interface{}{
{"id": "1", "title": "???? ????", "status": "pending"},
},
})
}


