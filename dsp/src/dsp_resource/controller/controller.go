package controller

import (
	"dsp/src/dsp_resource/service"
	"dsp/src/util"
	"dsp/src/util/cusfun"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

/**
 * @author  tianqiang
 * @date  2021/7/20 17:43
 */

/**
测试连通性接口
*/
func TelnetRes(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if params == nil || params["resIpV4"] == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	if err != nil {
		util.Fail(ctx, "参数解析失败 ", "")
		return
	}
	bool := service.TelnetRes(params)
	if !bool {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "1",
			"msg":  "连接失败",
			"data": "",
		})
	} else {
		util.Success(ctx, "连接成功 ")
	}
}

func ResIsExists(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	var pojo cusfun.ParamsPOJO
	err = mapstructure.Decode(params, &pojo)
	if err != nil {
		util.Fail(ctx, "参数解析失败 ", "")
		return
	}
	err, count := service.ResIsExists(pojo)
	if count > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "0",
			"msg":  "1",
			"data": "",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "0",
			"msg":  "0",
			"data": "",
		})
	}
}

func ResStoreIsExists(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	var pojo cusfun.ParamsPOJO
	err = mapstructure.Decode(params, &pojo)
	if err != nil {
		util.Fail(ctx, "参数解析失败 ", "")
		return
	}
	err, count := service.ResStoreIsExists(params["ip"].(string), params["name"].(string), params["storetype"].(string))
	if count > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "0",
			"msg":  "1",
			"data": "",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "0",
			"msg":  "0",
			"data": "",
		})
	}
}

func ResAccountIsExists(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if err != nil {
		util.Fail(ctx, "参数解析失败 ", "")
		return
	}
	accountId := ""
	if params["accountId"] != nil {
		accountId = params["accountId"].(string)
	}
	err, count := service.ResAccountIsExists(params["resId"].(string), params["resAccount"].(string), accountId)
	if count > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "0",
			"msg":  "1",
			"data": "",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "0",
			"msg":  "0",
			"data": "",
		})
	}
}