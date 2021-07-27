package controller

import (
	"dsp/src/dsp_user/model"
	"dsp/src/dsp_user/service"
	"dsp/src/util"
	"dsp/src/util/cusfun"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

/**
 * @author  tianqiang
 * @date  2021/7/19 10:57
 */

/*
查询角色
 */
func RoleList(ctx *gin.Context)  {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	var pojo cusfun.ParamsPOJO
	err = mapstructure.Decode(params,&pojo)
	if err != nil {
		util.Fail(ctx,"参数解析失败","")
		return
	}
	data,err,count := service.RoleFindList(pojo)
	util.MsgLog(ctx,nil,"","","",err,data,count)
}

/*
角色新增
*/
func RoleAdd(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if params == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	var role model.Role
	//参数格式化为model
	err = mapstructure.Decode(params, &role)
	if err != nil {
		util.Fail(ctx, "参数解析错误 ", "")
		return
	}
	//sessionInfo, err := cusfun.GetSessionInfo(c)
	//role.Createuserid = sessionInfo.UserId
	data, err := service.RoleInsert(role)
	util.MsgLog(ctx,nil,"","","",err,data,1)


}

/*
角色修改
*/
func RoleUpdate(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if params == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	var role model.Role
	//参数格式化为model
	err = mapstructure.Decode(params, &role)
	if err != nil {
		util.Fail(ctx, "参数解析错误 ", "")
		return
	}
	//sessionInfo, err := cusfun.GetSessionInfo(c)
	//role.Modifyuserid = sessionInfo.UserId
	data, err, count := service.RoleModify(role)
	util.MsgLog(ctx,nil,"","","",err,data,count)


}

/*
角色删除
*/
func RoleDel(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if params == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	var role model.Role
	//参数格式化为model
	err = mapstructure.Decode(params, &role)
	if err != nil {
		util.Fail(ctx, "参数解析错误 ", "")
		return
	}
	data, count, err := service.RoleDel(role)
	util.MsgLog(ctx,nil,"","","",err,data,count)
}