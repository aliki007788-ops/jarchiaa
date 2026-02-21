package main

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
v1 := r.Group("/v1")

text
v1.POST("/auth/otp", proxyToUserService)
v1.POST("/auth/verify", proxyToUserService)

v1.GET("/ads/nearby", proxyToAdService)
v1.GET("/vitrin", proxyToVitrinService)
}


