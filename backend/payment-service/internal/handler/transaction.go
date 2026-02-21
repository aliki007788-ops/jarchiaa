package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{
"transactions": []map[string]interface{}{
{"id": "1", "amount": 500000, "type": "deposit"},
},
})
}


