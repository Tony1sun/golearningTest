package logic

import (
	"context"

	"hello/internal/models"
	"hello/internal/svc"
	"hello/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.UserOptReq) (*types.UserOptResp, error) {
	user := models.User{
		Mobile: req.Mobile,
		Passwd: req.Passwd,
	}
	result := l.svcCtx.DbEngin.Create(&user)
	return &types.UserOptResp{
		Id: user.ID,
	}, result.Error
}
