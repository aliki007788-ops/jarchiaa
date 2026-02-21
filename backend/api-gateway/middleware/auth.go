package middleware

import (
"net/http"
"strings"
"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
return func(c *gin.Context) {
token := c.GetHeader("Authorization")
if token == "" {
c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
c.Abort()
return
}
c.Next()
}
}


