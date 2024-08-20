package businese

import (
	"github.com/gin-gonic/gin"
	"steward/models"
	"steward/system/base"
	"steward/system/constant"
)

type UserController struct {
	base.Base
}

func NewUserController() *UserController {
	return new(UserController)
}

func (uc *UserController) Add(ctx *gin.Context) {
	var req *models.UserStu
	if err := ctx.ShouldBindJSON(req); err != nil {
		uc.Failure(ctx, -1, constant.ParameterError)
		return
	}
	if err := NewService().AddUser(req); err != nil {
		uc.Failure(ctx, -1, constant.DataSaveFailure)
		return
	}
	uc.Success(ctx, nil)
}

func (uc *UserController) Delete(ctx *gin.Context) {
	var req *models.UserStu
	if err := ctx.ShouldBindJSON(req); err != nil {
		uc.Failure(ctx, -1, constant.ParameterError)
		return
	}
	if err := NewService().DeleteUser(req.Code); err != nil {
		uc.Failure(ctx, -1, constant.DataDeleteFailure)
		return
	}
	uc.Success(ctx, nil)
}

func (uc *UserController) Find(ctx *gin.Context) {
	var req *models.UserStu
	if err := ctx.ShouldBindQuery(req); err != nil {
		uc.Failure(ctx, -1, constant.ParameterError)
		return
	}
	mUser, err := NewService().FindUser(req)
	if err != nil {
		uc.Failure(ctx, -1, constant.DataQueryFailure)
		return
	}
	uc.Success(ctx, models.Transfer2UserDto(mUser))
}

func (uc *UserController) PageList(ctx *gin.Context) {
	var req *models.UserStu
	if err := ctx.ShouldBindJSON(req); err != nil {
		uc.Failure(ctx, -1, constant.ParameterError)
		return
	}
	mUsers, total, err := NewService().PageListUser(req)
	if err != nil {
		uc.Failure(ctx, -1, constant.DataQueryFailure)
		return
	}
	uc.PageSuccess(ctx, mUsers, total, req.Page)
}

func (uc *UserController) Update(ctx *gin.Context) {
	var req *models.UserStu
	if err := ctx.ShouldBindJSON(req); err != nil {
		uc.Failure(ctx, -1, constant.ParameterError)
		return
	}
	mUser, err := NewService().UpdateUser(models.Transfer2User(req))
	if err != nil {
		uc.Failure(ctx, -1, constant.DataSaveFailure)
		return
	}
	uc.Success(ctx, mUser)
}
