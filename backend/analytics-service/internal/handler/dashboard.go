package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleDashboard(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{
"total_users": 12540,
"active_today": 2345,
"total_revenue": 12500000,
})
}


