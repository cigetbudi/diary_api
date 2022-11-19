package controller

import (
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
