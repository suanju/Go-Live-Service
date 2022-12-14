package home

import (
	homeLogic "Go-Live/logic/home"
	"Go-Live/models/home"
	"Go-Live/utils/response"
	"Go-Live/utils/validator"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
}

//GetHomeInfo 获取主页信息
func (C Controllers) GetHomeInfo(ctx *gin.Context) {
	GetHomeInfoReceive := new(home.GetHomeInfoReceiveStruct)
	if err := ctx.ShouldBind(GetHomeInfoReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := homeLogic.GetHomeInfo(GetHomeInfoReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
