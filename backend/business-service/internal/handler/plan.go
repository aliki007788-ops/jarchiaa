package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func GetPlans(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{
"plans": []map[string]interface{}{
{"id": "silver", "name": "???????", "price": 250000},
{"id": "gold", "name": "?????", "price": 500000},
},
})
}


