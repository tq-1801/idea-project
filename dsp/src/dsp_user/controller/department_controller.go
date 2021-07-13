package controller

import (
	"dsp/src/dsp_user/service"
	"dsp/src/util"
	"dsp/src/util/cusfun"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

/**
 * @author  tianqiang
 * @date  2021/7/13 17:05
 */
/*
查询部门
*/
func DepartmentFindList(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	var pojo cusfun.ParamsPOJO
	err = mapstructure.Decode(params, &pojo)
	if err != nil {
		util.Fail(ctx, "参数解析失败 ", "")
		return
	}
	data, err, count := service.DepartmentFindList(pojo)
	util.MsgLog(ctx,nil,"","","",err,data,count)
}

//测试查询部门
// 信息查询
//func (c *UserController) PostBy(id string) *http.Response {
//	return db.QueryBy(id, &model.User{})

//func DerartmentList(ctx *gin.Context)  {
//	var count int
//	var params map[string]interface{}
//	err := ctx.ShouldBind(&params)
//	var pojo cusfun.ParamsPOJO
//	err = mapstructure.Decode(params, &pojo)
//	if err != nil {
//		util.Fail(ctx, "参数解析失败 ", "")
//		return
//	}
//	db := util.DbConn.Model(&model.Department{})
//
//
//
//	util.MsgLog(ctx,nil,"","","",err,db,count)
//}