package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func GetBusinesses(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{
"businesses": []map[string]interface{}{
{"id": "1", "name": "???????", "verified": false},
},
})
}


