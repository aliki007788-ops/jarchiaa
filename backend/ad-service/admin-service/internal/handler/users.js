package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{
"users": []map[string]interface{}{
{"id": "1", "name": "???", "phone": "09123456789"},
},
})
}


