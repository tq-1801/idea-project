package controller

import (
	"dsp/src/dsp_resource/model"
	"dsp/src/dsp_resource/service"
	"dsp/src/util"
	"dsp/src/util/cusfun"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"strconv"
)

/**
 * @author  tianqiang
 * @date  2021/7/20 17:35
 */

//分页查询列表
func ListAccount(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	var pojo cusfun.ParamsPOJO
	err = mapstructure.Decode(params, &pojo)
	if err != nil {
		util.Fail(ctx, "参数解析失败 ", "")
		return
	}
	data, err, count := service.ListAccount(pojo)
	util.MsgLog(ctx,nil,"","","",err,data,count)
}

//添加资产账号
func CreateAccount(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if params == nil || params["resId"] == nil || params["resAccount"] == nil || params["resAccountName"] == nil || params["resAccountPassword"] == nil || params["isAdmin"] == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	var account model.Account
	err = mapstructure.Decode(params, &account)
	if err != nil {
		util.Fail(ctx, "参数解析失败 ", "")
		return
	}

	data, err, count := service.CreateAccount(account)
	util.MsgLog(ctx,nil,"","","",err,data,count)
}

//修改资产账号
func UpdateAccount(ctx *gin.Context) {
	//util.CustomLogger.Debug("开始修改----------------------------")
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	util.CustomLogger.Debug(params)
	if params == nil || params["id"] == nil || params["resId"] == nil || params["resAccount"] == nil || params["resAccountName"] == nil || params["resAccountPassword"] == nil || params["resAccountStatus"] == nil || params["isAdmin"] == nil || params["isSuper"] == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	var account model.Account
	err = mapstructure.Decode(params, &account)
	util.CustomLogger.Debug(account)
	if err != nil {
		util.Fail(ctx, "参数解析失败 ", "")
		return
	}
	data, err, count := service.UpdateAccount(params, account)
	//util.CustomLogger.Debug("--------------------修改密码返回控制层开始-------------------")
	fmt.Println(err)
	//util.CustomLogger.Debug("--------------------修改密码返回控制层结束-------------------")
	//if err != nil {
	//	//util.CustomLogger.Debug("--------------------修改密码失败-------------------")
	//	util.Fail(ctx,err.Error(),"")
	//	return
	//}
	util.MsgLog(ctx,nil,"","","",err,data,count)
}

//删除资产账号
func DeleteAccount(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)

	if params == nil || params["id"] == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	var account model.Account
	err = mapstructure.Decode(params, &account)
	if err != nil {
		util.Fail(ctx, "参数解析失败 ", "")
		return
	}
	data, err, count := service.DeleteAccount(account)
	util.MsgLog(ctx,nil,"","","",err,data,count)
}

// 导出安全设备列表
func ExportAccount(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if params == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	var account model.Account
	err = mapstructure.Decode(params, &account)
	if err != nil {
		util.Fail(ctx, "参数解析失败 ", "")
		return
	}
	data, err, count := service.ExportAccount(account)
	var res []interface{}
	for _, info := range data {
		var exportInfo model.ExportAccount
		util.StructAssign(&exportInfo, &info)
		res = append(res, &exportInfo)
	}

	util.MsgLog(ctx,nil,"","","",err,data,count)
}

// 下载账号模板
func DownloadAccount(c *gin.Context) {
	fileName := "account.xlsx"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	//http.ServeFile(c.Writer, c.Request, util.Cfg.Template.Path+fileName)
}

//导入上传的资产账号信息，并入库
func CreateAccountMore(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if ctx.PostForm("res_id") == "" {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	resId, _ := strconv.Atoi(ctx.PostForm("res_id"))
	file, err := ctx.FormFile("file")
	//将上传的文档保存到磁盘上的位置
	filePath := "upload/" + file.Filename
	ctx.SaveUploadedFile(file, filePath)
	data, err, count := service.CreateAccountMore(filePath, resId)
	util.MsgLog(ctx,nil,"","","",err,data,count)
}

//分页查询列表
func ListAccountByResId(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if err != nil {
		util.Fail(ctx, "参数解析失败 ", "")
		return
	}
	str := params["resId"].(string)
	fmt.Println(str)
	data, count, err := service.ListAccountByResId(str)
	util.MsgLog(ctx,nil,"","","",err,data,count)
}