package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func GetModules(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{
"modules": []map[string]interface{}{
{"id": "user_management", "name": "?????? ???????", "is_active": true},
{"id": "business_management", "name": "?????? ????????", "is_active": true},
},
})
}


