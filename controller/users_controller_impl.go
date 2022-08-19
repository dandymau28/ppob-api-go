package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ppob-api.go/ppob-api/dto"
	"ppob-api.go/ppob-api/entity"
	"ppob-api.go/ppob-api/helper"
	"ppob-api.go/ppob-api/service"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) CheckPhoneNumber(ctx *gin.Context) {
	phone := ctx.Param("phone")

	isPhoneNumberExist := controller.UserService.CheckPhoneNumber(phone)

	if isPhoneNumberExist {
		ctx.JSON(http.StatusConflict, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (controller *UserControllerImpl) Register(ctx *gin.Context) {
	var request dto.RegisterRequest

	if err := ctx.BindJSON(&request); err != nil {
		res := helper.BuildErrorResponse("request invalid", err.Error(), err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	user := entity.User{
		Nohandphone: request.NoHandphone,
		Macaddress:  request.MacAddress,
	}

	if err := controller.UserService.Create(&user); err != nil {
		res := helper.BuildErrorResponse("request invalid", err.Error(), err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	response := dto.RegisterResponse{
		ID:          user.Id.Hex(),
		NoHandphone: user.Nohandphone,
		MacAddress:  user.Macaddress,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	res := helper.BuildResponse(true, "success", response)
	ctx.JSON(http.StatusOK, res)
}

func (controller *UserControllerImpl) GenerateOTP(ctx *gin.Context) {
	phone := ctx.Param("phone")

	err := controller.UserService.GenerateOTP(phone)

	if err != nil {
		res := helper.BuildErrorResponse("request invalid", err.Error(), err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
