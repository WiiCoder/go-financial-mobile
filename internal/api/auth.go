package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wiiCoder/installment/internal/tools"
	"github.com/wiiCoder/installment/pkg/struct/api"
	"github.com/wiiCoder/installment/pkg/struct/db"
	"gorm.io/gorm"
	"net/http"
)

type AuthApi struct {
	db *gorm.DB
}

func NewAuthApi(db *gorm.DB) AuthApi {
	return AuthApi{db}
}

func (a *AuthApi) Register(ctx *gin.Context) {

}

func (a *AuthApi) Login(ctx *gin.Context) {
	var params api.LoginRequest
	ctx.Bind(&params)

	var user db.AppUser
	tx := a.db.Where("username = ?", params.Username).First(&user)

	if tx.Error != nil {
		ctx.JSON(http.StatusOK, api.ApiError(tx.Error))
		return
	}

	checked := tools.CheckPassword(params.Password+user.Salt, user.Password)
	if !checked {
		ctx.JSON(http.StatusUnauthorized, api.ApiError(errors.New("wrong password")))
		return
	}

	ctx.JSON(http.StatusOK, api.ApiSuccess(user))
}
