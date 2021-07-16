package service

import (
	"crypto/md5"
	"dsp/src/dsp_user/dao"
	"dsp/src/dsp_user/model"
	"dsp/src/util"
	"dsp/src/util/cusfun"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/pikanezi/mapslice"
	"github.com/tealeg/xlsx"

	"io"
	"os"
	"regexp"
	"runtime"
	"time"
)

/**
 * @author  tianqiang
 * @date  2021/7/14 10:20
 */


/*
查询用户列表
*/
func UserFindList(pojo cusfun.ParamsPOJO) (interface{}, int, error) {
	res,count,err := dao.UserFind(pojo)
	return res,count,err
}

/*
增加用户
 */
func UserInsert(user model.User) (rst interface{}, count int,err error) {
	//查询账户是否重复
	if dao.ExistUserFindByUid(user.Uid) > 0 {
		err := errors.New("账户名已存在")
		return user,count,err
	}
	if dao.ExistUserFindByUname(user.Uname) > 0 {
		err := errors.New("姓名已存在")
		return user, count, err
	}
	if user.EmployeeNumber != "" && dao.ExistUserFindByEmpNum(user.EmployeeNumber) > 0 {
		err := errors.New("员工号已存在")
		return user,count,err
	}
	if user.Mobile != "" && dao.ExistUserFindByMobile(user.Mobile) > 0 {
		err := errors.New("手机号已存在")
		return user,count,err
	}
	if user.Cardid != "" && dao.ExistUserFindByCardId(user.Cardid) > 0 {
		err := errors.New("身份证号已存在")
		return user,count,err
	}
	if user.Mail != "" && dao.ExistUserFindByMail(user.Mail) > 0 {
		err := errors.New("邮箱已存在")
		return user,count,err
	}
	if user.ZDesc == "" {
		user.ZDesc = "无描述"
	}
	if user.Uid == "" {
		err := errors.New("用户名不能为空")
		return user, count, err
	}
	if user.Uname == "" {
		err := errors.New("姓名不能为空")
		return user, count, err
	}
	if user.DepartmentId == 0 {
		err := errors.New("部门不能为空")
		return user, count, err
	}
	//密码为空，则密码默认为用户id
	if user.Password == "" {
		user.Password = user.Uid
	}
	var aes *util.AesEncrypt
	user.Password, _ = aes.Encrypt(user.Password)
	user.Status = "2"
	//增加用户，如果未选择角色，增加的用户属于默认角色，默认角色id为-1
	if user.Roleid == 0 {
		user.Roleid = -1
	}
	//如果未选密码策略,则选择默认策略
	if user.UserPasswordPolicyId == 0 {
		user.UserPasswordPolicyId = -1
	}
	//根据身份日期生成出生日期
	if user.Cardid != "" {
		user.Birthdate = user.Cardid[6:14]
	}
	if user.UserAuthenPolicyId == 0 {
		user.UserAuthenPolicyId = -1
	}
	if dao.PolicyFindById(user.UserAuthenPolicyId).AuthenType == 2 {
		secret := cusfun.NewGoogleAuth().GetSecret()
		_, qrPng, _ := cusfun.NewGoogleAuth().CreateQrcode(user.Uid, user.Mobile, secret)
		user.AuthKey = secret
		user.QrCode = qrPng
	}
	//自动更新密码过期日期
	a := dao.UserPwdList(user.UserPasswordPolicyId)
	nowTime := time.Now()
	newTime := nowTime.AddDate(0, 0, a.PolicyValidity).Format("2006-01-02")
	user.Expiretime = newTime

	res, count,err := dao.UserInsert(user)

	return res, count,err
}

/*
修改用户
 */
func UserUpdate(user model.User) (rst interface{},count int,err error)  {
	res, err, count := dao.UserFindById(user.Uid)

	if res.Uid != user.Uid {
		err = errors.New("用户账号无法修改")
		return user, count,err
	}
	if res.Uname != user.Uname {
		if dao.ExistUserFindByUname(user.Uname) > 0 {
			err = errors.New("用户姓名已存在")
		}
	}
	if res.Mobile != user.Mobile {
		if user.Mobile != "" && dao.ExistUserFindByMobile(user.Mobile) > 0 {
			err := errors.New("手机号已存在")
			return user,count,err
		}
	}
	if res.EmployeeNumber != user.EmployeeNumber {
		if user.EmployeeNumber != "" && dao.ExistUserFindByEmpNum(user.EmployeeNumber) > 0 {
			err = errors.New("员工号已存在")
			return user, count,err
		}
	}
	if res.Cardid != user.Cardid {
		if user.Cardid != "" && dao.ExistUserFindByCardId(user.Cardid) > 0 {
			err := errors.New("身份证号已存在")
			return user,count,err
		}
	}
	if res.Mail != user.Mail {
		if user.Mail != "" && dao.ExistUserFindByMail(user.Mail) > 0 {
			err := errors.New("邮箱已存在")
			return user,count,err
		}
	}
	if user.Uid == "" {
		err := errors.New("用户名不能为空")
		return user, count, err
	}
	if user.Uname == "" {
		err := errors.New("姓名不能为空")
		return user, count, err
	}
	if user.DepartmentId == 0 {
		err := errors.New("部门不能为空")
		return user, count, err
	}
	if user.DepartmentId == 3 {
		return user,count, errors.New("已删除数据无法修改")
	}
	//根据身份日期生成出生日期
	if user.Cardid != "" {
		user.Birthdate = user.Cardid[6:14]
	}
	if dao.PolicyFindById(user.UserAuthenPolicyId).AuthenType == 2 {
		secret := cusfun.NewGoogleAuth().GetSecret()
		_, qrPng, _ := cusfun.NewGoogleAuth().CreateQrcode(user.Uid, user.Mobile, secret)
		user.AuthKey = secret
		user.QrCode = qrPng
	} else {
		user.AuthKey = ""
		user.QrCode = ""
	}
	//如果重置密码
	if user.Password != res.Password || user.UserPasswordPolicyId != res.UserPasswordPolicyId {
		//自动更新密码过期日期
		a := dao.UserPwdList(user.UserPasswordPolicyId)
		nowTime := time.Now()
		newTime := nowTime.AddDate(0, 0, a.PolicyValidity).Format("2006-01-02")
		user.Expiretime = newTime
	}

	res, count, err = dao.UserUpdate(user)

	return res,count,err
}

/*
修改密码
 */
func ModifyPwd(userid, oldPwd, newPwd string) (data string, err error) {
	//查询账号信息，有无账号，获得密码
	user, err, count := dao.UserFindById(userid)
	if err != nil || count != 1 {
		return "", errors.New("未找到相关用户数据")
	}
	//加密
	w := md5.New()
	io.WriteString(w, oldPwd) //将str写入到w中
	bw := w.Sum(nil)          //w.Sum(nil)将w的hash转成[]byte格式
	oldPwd = hex.EncodeToString(bw)
	//验证密码
	oriPwd := user.Password
	w1 := md5.New()
	io.WriteString(w1, newPwd) //将str写入到w中
	bw1 := w1.Sum(nil)         //w.Sum(nil)将w的hash转成[]byte格式
	temporarynewPwd := hex.EncodeToString(bw1)
	if oriPwd != oldPwd {
		return "", errors.New("旧密码错误")
	}
	if oriPwd == temporarynewPwd {
		return "", errors.New("新密码和旧密码不能相同")
	}
	if len(newPwd) < 8 {
		return "", errors.New("密码长度最少有8位")
	}
	pattern1 := "\\d+" //反斜杠要转义
	res1, _ := regexp.MatchString(pattern1, newPwd)
	if !res1 {
		return "", errors.New("密码必须包含数字")
	}
	pattern2 := "[a-zA-Z]"
	res2, _ := regexp.MatchString(pattern2, newPwd)
	if !res2 {
		return "", errors.New("密码必须包含字母")
	}
	pattern3 := "\\W+"
	res3, _ := regexp.MatchString(pattern3, newPwd)
	if !res3 {
		return "", errors.New("密码必须包含字符")
	}
	user.Password = temporarynewPwd
	//自动更新密码过期日期
	a := dao.UserPwdList(user.UserPasswordPolicyId)
	nowTime := time.Now()
	newTime := nowTime.AddDate(0, 0, a.PolicyValidity).Format("2006-01-02")
	user.Expiretime = newTime
	_,count, err = dao.UserUpdate(user) //更新用户
	if err != nil {
		return "", errors.New("数据库异常")
	}
	return "", nil
}

/*
重置密码
 */
func ResetPwd(userid string) (data string, err error) {
	user, err, count := dao.UserFindById(userid)
	if err != nil || count != 1 {
		return "重置失败", errors.New("未找到相关用户数据")
	}
	//加密
	user.Password = user.Uid
	var aes *util.AesEncrypt
	user.Password, _ = aes.Encrypt(user.Password)
	//自动更新密码过期日期
	a := dao.UserPwdList(user.UserPasswordPolicyId)
	nowTime := time.Now()
	newTime := nowTime.AddDate(0, 0, a.PolicyValidity).Format("2006-01-02")
	user.Expiretime = newTime
	user.Modifypwddate = time.Now().Format("2006-01-02 15:04:05")
	user.Lastlogin = time.Now().Format("0000-00-00 00:00:00")
	_,count, err = dao.UserUpdate(user) //更新用户
	if err != nil {
		return "重置失败", errors.New("数据库异常")
	}
	return "重置成功", nil
}

/*
逻辑删除用户
 */
func UserDel(user model.User,userid string) (rst interface{},count int,err error)  {
	res, err, count := dao.UserFindById(user.Uid)
	//初始化数据不允许删除
	if res.Islinlay == 1 {
		err = errors.New("初始化数据不可被删除")
	}
	//当前登录账号不允许被删除
	if userid == res.Uid {
		err = errors.New("当前登录用户不允许被删除")
	}
	//逻辑删除
	if err == nil {
		user.Status = "0"
		user.DepartmentId = 3
		res, count, err = dao.UserDelById(user)
	}
	return res,count,err
	
}

//导出
func Export(pojo cusfun.ParamsPOJO) (interface{}, error) {
	var expdplist []model.Expdp
	expdplist, _, _ = dao.UserFindEx(pojo)

	//定义单条数据
	var expdp model.ExpdpAll
	//多条数据
	var expdplists []model.ExpdpAll
	//得到数据库全部数据   数组形式
	//expdplist, _ := GetExpdpList()
	//遍历每一行的数据存储到expdplists
	for i := 0; i < len(expdplist); i++ {
		expdp.Uid = expdplist[i].Uid
		expdp.Uname = expdplist[i].Uname
		//expdp.Sex = expdplist[i].Sex
		expdp.Role = dao.RoleList(expdplist[i].Roleid).Rolename
		//expdp.Birthdate = expdplist[i].Birthdate
		//expdp.Cardid = expdplist[i].Cardid
		expdp.Mobile = expdplist[i].Mobile
		//expdp.DicCity = dao.CityList(expdplist[i]).DicName
		expdp.Department = dao.DeptList(expdplist[i].DepartmentId).DepName
		expdp.EmployeeNumber = expdplist[i].EmployeeNumber
		expdp.Mail = expdplist[i].Mail
		//组装
		expdplists = append(expdplists, expdp)
	}
	//取出数据
	uid, _ := mapslice.ToStrings(expdplists, "Uid")
	uname, _ := mapslice.ToStrings(expdplists, "Uname")
	//sex, _ := mapslice.ToStrings(expdplists, "Sex")
	role, _ := mapslice.ToStrings(expdplists, "Role")
	//birthdate, _ := mapslice.ToStrings(expdplists, "Birthdate")
	//cardid, _ := mapslice.ToStrings(expdplists, "Cardid")
	mobile, _ := mapslice.ToStrings(expdplists, "Mobile")
	//diccity, _ := mapslice.ToStrings(expdplists, "DicCity")
	department, _ := mapslice.ToStrings(expdplists, "Department")
	employeeNumber, _ := mapslice.ToStrings(expdplists, "EmployeeNumber")
	mail, _ := mapslice.ToStrings(expdplists, "Mail")

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("用户列表_" + time.Now().Format("20060102_150405") + ".xlsx")
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "用户账号"
	cell = row.AddCell()
	cell.Value = "用户姓名"
	//cell = row.AddCell()
	//cell.Value = "性别"
	cell = row.AddCell()
	cell.Value = "所属角色"
	//cell = row.AddCell()
	//cell.Value = "出生日期"
	//cell = row.AddCell()
	//cell.Value = "身份证号码"
	cell = row.AddCell()
	cell.Value = "手机号码"
	cell = row.AddCell()
	//cell.Value = "城市"
	//cell = row.AddCell()
	cell.Value = "部门"
	cell = row.AddCell()
	cell.Value = "员工编号"
	cell = row.AddCell()
	cell.Value = "邮箱"
	//遍历每行添加
	for i := 0; i < len(uid); i++ {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = uid[i]
		cell = row.AddCell()
		cell.Value = uname[i]
		//cell = row.AddCell()
		//cell.Value = sex[i]
		cell = row.AddCell()
		cell.Value = role[i]
		//cell = row.AddCell()
		//cell.Value = birthdate[i]
		//cell = row.AddCell()
		//cell.Value = cardid[i]
		cell = row.AddCell()
		cell.Value = mobile[i]
		//cell = row.AddCell()
		//cell.Value = diccity[i]
		cell = row.AddCell()
		cell.Value = department[i]
		cell = row.AddCell()
		cell.Value = employeeNumber[i]
		cell = row.AddCell()
		cell.Value = mail[i]
	}
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" { //开发模式
		if err := os.MkdirAll("./userinfo", 0777); err != nil {
			fmt.Println(err.Error())
		}
		file.Save("./userinfo/用户列表_" + time.Now().Format("20060102_150405") + ".xlsx")
		return "/api/sf-user/users/download/" + "用户列表_" + time.Now().Format("20060102_150405") + ".xlsx", nil
	} else { //部署模式
		if err := os.MkdirAll("/tmp/userinfo", 0777); err != nil {
			fmt.Println(err.Error())
		}
		file.Save("/tmp/userinfo/用户列表_" + time.Now().Format("20060102_150405") + ".xlsx")
		return "/api/sf-user/users/download/" + "用户列表_" + time.Now().Format("20060102_150405") + ".xlsx", nil
	}
}