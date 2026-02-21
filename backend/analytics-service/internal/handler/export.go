package handler

import (
"net/http"
"github.com/gin-gonic/gin"
)

func HandleExport(c *gin.Context) {
c.Header("Content-Disposition", "attachment; filename=report.csv")
c.String(http.StatusOK, "date,views\n2024-01-01,100")
}


