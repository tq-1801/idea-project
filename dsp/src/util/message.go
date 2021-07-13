package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 *cotroller层通用返回信息函数
 *无需日志记录时，operType和content为""
 *非查询接口，count为0
 */
func MsgLog(c *gin.Context, params map[string]interface{}, operType string, logType string, content string, err error, data interface{}, count int) {
	msg := ""
	code := "1"
	if err != nil {
		msg = "操作失败，" + err.Error()
	} else {
		code = "0"
		msg = "操作成功"
	}

	if operType == "导出" {
		Success(c, data.(string))
		//c.String(200, data.(string))
	} else {
		Msg(c, code, msg, data, count)
	}
}

///**
// *写日志到kafka中
// *返回给前台信息多个参数
// */
//func MsgLogWithKafka(ctx *gin.Context, params map[string]interface{}, data interface{}, count int, err error, logContent string) {
//	returnResult := ""
//	result := ""
//	operateType := ""
//	code := "1"
//	operateTypeMap := map[string]string{
//		"sso":      "单点登录",
//		"list":     "查询",
//		"create":   "新增",
//		"update":   "修改",
//		"delete":   "删除",
//		"login":    "登录",
//		"logout":   "登出",
//		"upload":   "文件上传",
//		"upgrade":  "系统升级",
//		"restore":  "数据恢复",
//		"download": "文件下载",
//		"shutdown": "系统关机",
//		"reboot":   "系统重启",
//		"import":   "文件导入",
//		"export":   "文件导出",
//	}
//	for key, value := range operateTypeMap {
//		if strings.Contains(ctx.Request.RequestURI, key) {
//			operateType = value
//			break
//		} else {
//			operateType = "操作"
//			break
//		}
//	}
//	if err != nil {
//		if err.Error() == "首次登录，需要修改密码" {
//			code = "0"
//		}
//		result = operateType + "失败"
//
//		returnResult = result
//		if containChineseWord(err.Error()) {
//			returnResult = err.Error()
//		}
//	} else {
//		code = "0"
//		result = operateType + "成功"
//
//		returnResult = result
//	}
//	// 继续交由下一个路由处理,并将解析出的信息传递下去
//	if params == nil {
//		params = make(map[string]interface{}, 0)
//	}
//	params["result"] = result
//	params["content"] = logContent
//	if data != nil {
//		dataName := reflect.TypeOf(data)
//		//只有登录才取值
//		if dataName.String() == "cusfun.LoginResp" {
//			dataValue := reflect.ValueOf(data)
//			params["sid"] = dataValue.Field(2).Interface()
//			params["isReplace"] = dataValue.Field(3).Interface()
//			params["clientIp"] = dataValue.Field(4).Interface()
//		}
//	}
//	ctx.Set("params", params)
//
//	/*if err !=nil && err.Error()=="首次登录，需要修改密码"{
//			ctx.JSON(http.StatusPaymentRequired, gin.H{
//				"code":  code,
//				"msg":   result,
//				"data":  data,
//				"count": count,
//			})
//	}else{*/
//	ctx.JSON(http.StatusOK, gin.H{
//		"code":  code,
//		"msg":   returnResult,
//		"data":  data,
//		"count": count,
//	})
//	//}
//
//}
//
//func MsgLogWithCmd(c *gin.Context, params map[string]interface{}, operType, logType, content string, err error, data interface{}, count int, cmdRst string) {
//	msg := ""
//	code := "1"
//	if err != nil {
//		msg = "操作失败"
//	} else {
//		code = "0"
//		msg = "操作成功"
//	}
//	if operType != "" {
//		if logType != "" {
//			fmt.Println("日志记录")
//			insetLog(c, operType, logType, content, msg, params)
//		} else {
//			fmt.Println("日志记录")
//			insetLog(c, operType, "操作日志", content, msg, params)
//		}
//	}
//	Msg(c, code, operType+msg+" "+cmdRst, data, count)
//}
//
func Msg(c *gin.Context, code string, msg string, data interface{}, count int) {
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   msg,
		"data":  data,
		"count": count,
	})
}
//
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": "0",
		"msg":  "成功",
		"data": data,
	})
}
//
func Fail(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": "1",
		"msg":  msg,
		"data": data,
	})
}
//
////func getLog(ctx *gin.Context, params map[string]interface{}) (logInfo SomfLog){
////var resLog ResLog
////var claims cusfun.CustomClaims
////clm, isExists := ctx.Get("claims")
////if isExists {
////	mapstructure.Decode(clm, &claims)
////	logInfo.AccessID = claims.Sid
////	logInfo.UserID = claims.UserId
////	logInfo.UserName = claims.UserName
////	logInfo.UserAddr = claims.ClientIp
////	logInfo.UserDepartment = claims.UserDepartment
////}
////logInfo.ID = uuid.NewV4().String()
////logInfo.LogTime = time.Now().Format("2006-01-02T15:04:05")
////logInfo.LogType = "0x99"
////logInfo.Operate = operateType
////if strings.EqualFold("登录", operateType) || strings.EqualFold("登出", operateType) {
////	logInfo.LogSubtype = "0x01"
////	logInfo.LogSubtypeDesc = "登录日志"
////} else if strings.EqualFold("单点登录", operateType) {
////	logInfo.LogSubtype = "0x02"
////	logInfo.LogSubtypeDesc = "单点登录"
////} else {
////	logInfo.LogSubtype = "0x03"
////	logInfo.LogSubtypeDesc = "操作日志"
////}
////	return
////}
//
//func insertLogToKafka(ctx *gin.Context, params map[string]interface{}, operateType, obj, content, rst string) {
//	//var logs SomfLog
//	//var resLog ResLog
//	//var claims cusfun.CustomClaims
//	//clm, isExists := ctx.Get("claims")
//	//if isExists {
//	//	mapstructure.Decode(clm, &claims)
//	//	logs.AccessID = claims.Sid
//	//	logs.UserID = claims.UserId
//	//	logs.UserName = claims.UserName
//	//	logs.UserAddr = claims.ClientIp
//	//	logs.UserDepartment = claims.UserDepartment
//	//}
//	//logs.ID = uuid.NewV4().String()
//	//logs.LogTime = time.Now().Format("2006-01-02T15:04:05")
//	//logs.LogType = "0x99"
//	//logs.Operate = operateType
//	//if strings.EqualFold("登录", operateType) || strings.EqualFold("登出", operateType) {
//	//	logs.LogSubtype = "0x01"
//	//	logs.LogSubtypeDesc = "登录日志"
//	//} else if strings.EqualFold("单点登录", operateType) {
//	//	logs.LogSubtype = "0x02"
//	//	logs.LogSubtypeDesc = "单点登录"
//	//} else {
//	//	logs.LogSubtype = "0x03"
//	//	logs.LogSubtypeDesc = "操作日志"
//	//}
//	//logs.ResLog = resLog
//	//logs.Object = obj
//	//logs.Result = rst
//	//if content != "" {
//	//	logs.Content = content
//	//} else {
//	//	logs.Content = logs.UserName + "使用账号" + logs.UserID + "于" + logs.LogTime + "通过" + logs.UserAddr + logs.Operate + logs.Object + "，" + logs.Result
//	//}
//	//s := log.Operator + "," + log.Address + "," + content + "," + msg + "," + logType + "," + log.Opertime.Format("2006-01-02 15:04:05")
//	//log.Keywords = s
//	////入本地库
//	//DbConn.Create(&log)
//
//}
//
////func insetLog(c *gin.Context, operType string, logType string, content string, msg string, data map[string]interface{}) {
////	id := int(HashStr32(uuid.NewV4().String()))
////	operator := c.Request.Header.Get("User")
////	address := c.Request.Header.Get("Readdr")
////	if operator == "" {
////		var logdata logdata
////		err := mapstructure.Decode(data, &logdata)
////		if err == nil {
////			if logdata.User != "" {
////				operator = logdata.User
////			}
////			if logdata.Addr != "" {
////				address = logdata.Addr
////			}
////			if logdata.Readdr != "" {
////				address = logdata.Readdr
////			}
////		}
////	}
////	log := model.Syslog{
////		Id:             id,
////		Operator:       operator,
////		Opertype:       operType,
////		Result:         msg,
////		Opertime:       time.Now(),
////		Content:        content,
////		Address:        address,
////		Keywords:       "",
////		Logmillisecond: time.Now().UnixNano() / 1e6,
////		Logtype:        logType,
////	}
////	s := log.Operator + "," + log.Address + "," + content + "," + msg + "," + logType + "," + log.Opertime.Format("2006-01-02 15:04:05")
////	log.Keywords = s
////	//入本地库
////	DbConn.Create(&log)
////
////}
//
//type logdata struct {
//	User   string
//	Addr   string
//	Readdr string
//}
//
//func containChineseWord(str string) bool {
//	r := []rune(str)
//
//	for _, a := range r {
//		if a >= 19968 && a <= 40869 {
//			return true
//		}
//	}
//
//	return false
//}
