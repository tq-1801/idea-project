package controller

import (
	"dsp/src/dsp_user/dao"
	"dsp/src/dsp_user/model"
	"dsp/src/dsp_user/service"
	"dsp/src/util"
	"dsp/src/util/cusfun"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"strconv"
	"time"
)

/**
 * @author  tianqiang
 * @date  2021/7/14 10:26
 */


/*
查询用户
*/
func UserList(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	var pojo cusfun.ParamsPOJO
	err = mapstructure.Decode(params, &pojo)
	if err != nil {
		util.Fail(ctx, "参数解析tx失败 ", "")
		return
	}
	data,count,err := service.UserFindList(pojo)
	util.MsgLog(ctx,nil,"","","",err,data,count)
}

/*
增加用户
 */
func UserAdd(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if params == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	var user model.User
	//参数格式化为model
	err = mapstructure.Decode(params, &user)
	if err != nil {
		util.Fail(ctx, "参数解析错误 ", "")
		return
	}
	data,err := service.UserInsert(user)

	util.MsgLog(ctx,nil,"","","",err,data,1)

}

/*
修改用户
 */
func UserUpdate(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if params == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	var user model.User
	//参数格式化为model
	err = mapstructure.Decode(params, &user)
	if err != nil {
		util.Fail(ctx, "参数解析错误 ", "")
		return
	}
	data,count,err := service.UserUpdate(user)

	util.MsgLog(ctx,nil,"","","",err,data,count)

}

/*
用户锁定
*/
func UserLockUpdate(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if params == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	var user model.User
	err = mapstructure.Decode(params, &user)
	if err != nil {
		util.Fail(ctx, "参数解析错误 ", "")
		return
	}
	//sessionInfo, err := cusfun.GetSessionInfo(c)
	//user.Modifyuserid = sessionInfo.UserId
	//修改锁定状态，同时重置其锁定时间和密码错误次数
	if user.Status == "1" {
		var locktime int
		locktime = dao.UserPwdList(user.UserPasswordPolicyId).LockDuration
		now := time.Now()
		m, _ := time.ParseDuration(strconv.Itoa(locktime) + "h")
		now = now.Add(m)
		user.Locktime = now.Format("2006-01-02 15:04:05") //重置锁定时间
	}
	if user.Status == "2" {
		user.Locktime = time.Now().Format("2006-01-02 15:04:05")     //重置锁定时间
		user.Pwderrorunm = 0                                         //重置错误次数
		user.Pwderrortime = time.Now().Format("2006-01-02 15:04:05") //重置错误时间
	}
	data, count, err := service.UserUpdate(user)
	
	util.MsgLog(ctx,nil,"","","",err,data,count)
}

/**
修改密码
*/
func ModifyPwd(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if params == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	//var userid string
	//sessionInfo, err := cusfun.GetSessionInfo(c)
	//userid := sessionInfo.UserId
	userid := params["userid"].(string)
	oldPwd := params["password"].(string)
	newPwd := params["newpassword"].(string)
	data, err, count := service.ModifyPwd(userid, oldPwd, newPwd)
	util.MsgLog(ctx,nil,"","","",err,data,count)
}

/*
重置密码
*/
func ResetPwd(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if params == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	userid := params["uid"].(string)
	data, err,count := service.ResetPwd(userid)
	util.MsgLog(ctx,nil,"","","",err,data,count)
}

/*
逻辑删除用户
 */
func UserDel(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	if params == nil {
		util.Fail(ctx, "参数缺失 ", "")
		return
	}
	var user model.User
	//参数格式化为model
	err = mapstructure.Decode(params, &user)
	var userid string
	if err != nil {
		util.Fail(ctx, "参数解析错误 ", "")
		return
	}
	data,count,err := service.UserDel(user,userid)

	util.MsgLog(ctx,nil,"","","",err,data,count)

}

//导入
func Import(ctx *gin.Context) {
	mfrom, err := ctx.MultipartForm()
	if err != nil || mfrom.File == nil {
		util.Fail(ctx, "文件缺失", "")
		return
	}

	files, ok := mfrom.File["file"]
	if !ok || len(files) == 0 {
		util.Fail(ctx, "文件缺失", "")
		return
	}

	reader, err := files[0].Open()
	if err != nil {
		util.Fail(ctx, "文件缺失", "")
		return
	}
	defer reader.Close()

	data, err := service.Import(reader, files[0].Size)

	util.MsgLog(ctx,nil,"","","",err,data,1)
}

//导出
func Export(ctx *gin.Context) {
	var params map[string]interface{}
	err := ctx.ShouldBind(&params)
	var pojo cusfun.ParamsPOJO
	err = mapstructure.Decode(params, &pojo)
	if err != nil {
		util.Fail(ctx, "参数解析失败 ", "")
		return
	}
	data, err := service.Export(pojo)

	util.MsgLog(ctx,nil,"","","",err,data,1)
}