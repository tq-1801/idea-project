package cusfun

/**
 * @author  tianqiang
 * @date  2021/7/14 9:23
 */
// 增加
//func Create(Ctx gin.Context, model interface{}) *http.Response {
//	err := Ctx.ReadJSON(&model) // 读取Json请求参数
//
//	if err != nil { // 读取Json错误，返回请求参数格式错误
//		return http.ResJsonError()
//	}
//
//	err = wolf.Verify(model) // 校验参数合法性
//	if err != nil {          // 参数有误，返回参数错误信息
//		return http.ResParamError(err.Error())
//	}
//
//	err = wolf.DB().Create(model).Error // 增加数据
//	if err != nil {                     // 增加错误，返回异常错误信息
//		return http.ResError(err.Error())
//	}
//
//	return http.ResData(model) // 返回成功
//}