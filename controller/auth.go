package controller

import (
	"diary_api/helper"
	"diary_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var vm model.AuthInput

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := model.User{
		Username: vm.Username,
		Password: vm.Password,
	}

	savedUser, err := u.Save()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(ctx *gin.Context) {
	var vm model.AuthInput

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.FindUserByUsername(vm.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(vm.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
