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
 * @date  2021/7/13 17:05
 */

/*
查询部门列表
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

//递归部门查询
func DepartmentTreeList(ctx *gin.Context){
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	var zDepartment model.Department
	//参数格式化为model
	err = mapstructure.Decode(params, &zDepartment)

	supId := params["supId"].(float64)
	//isleaf := params["isleaf"].(float64)
	if params == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	data, err, count := service.DepTreeList(int(supId))
	util.MsgLog(ctx,nil,"","","",err,data,count)
}

//通过当前部门id查询下级部门
func DepartmentFindByIdList(ctx *gin.Context) {
	var params map[string] interface{}
	err := ctx.ShouldBind(&params)
	var zDepartment model.Department
	//参数格式化为model
	err = mapstructure.Decode(params,&zDepartment)

	id := (params["id"]).(float64)
	if params == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return

	}
	data, err, count := service.DepFindByIdList(int(id))
	util.MsgLog(ctx,nil,"","","",err,data,count)

}

/*
部门新增
*/
func DepartmentAdd(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if params == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	var zDepartment model.Department
	//参数格式化为model
	err = mapstructure.Decode(params, &zDepartment)
	zDepartment.CreateUserId = ctx.Request.Header.Get("Department")
	if err != nil {
		util.Fail(ctx, "参数解析错误 ", "")
		return
	}
	data, err, count := service.DepInsert(zDepartment)

		util.MsgLog(ctx,nil,"","","",err,data,count)


}

/*
部门更新
*/
func DepartmentUpdate(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if params == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	var zDepartment model.Department
	//参数格式化为model
	err = mapstructure.Decode(params, &zDepartment)
	zDepartment.CreateUserId = ctx.Request.Header.Get("Department")
	if err != nil {
		util.Fail(ctx, "参数解析错误 ", "")
		return
	}
	data, err, count := service.DepartmentUpdate(zDepartment)

		util.MsgLog(ctx,nil,"","","",err,data,count)


}

/*
部门删除
*/
func DepartmentDel(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if params == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	var zDepartment model.Department

	//参数格式化为model
	err = mapstructure.Decode(params, &zDepartment)
	if err != nil {
		util.Fail(ctx, "参数解析错误 ", "")
		return
	}

	data, err, count := service.DepartmentDel(zDepartment)

	util.MsgLog(ctx,nil,"","","",err,data,count)


}



