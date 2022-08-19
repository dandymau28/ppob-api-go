package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	CheckPhoneNumber(ctx *gin.Context)
	Register(ctx *gin.Context)
	GenerateOTP(ctx *gin.Context)
}
