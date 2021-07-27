package dao

import (
	"crypto/md5"
	"dsp/src/dsp_user/model"
	"dsp/src/util"
	"dsp/src/util/cusfun"
	"encoding/hex"
	"errors"
	"io"
	"regexp"
	"time"
)

/**
 * @author  tianqiang
 * @date  2021/7/14 10:22
 */

/*
查询用户列表
 */
func UserFind(pojo cusfun.ParamsPOJO) (rst []model.UserAll, total int, err error) {
	var count int64
	db := util.DbConn.Table(`(
					SELECT z.*,
                           b.rolename As rolename
					from(SELECT
						z_user.*,
						z_department.dep_name 
					FROM
						z_user
						LEFT JOIN z_department ON z_user.department_id = z_department.id
					) z 
					left join z_role b on z.roleid = b.id where isadmin != 1
			)t`)
	err = cusfun.GetSqlByParams(db, pojo, &count).Find(&rst).Error
	total = int(count)
	return
}

/*
增加用户
 */
func UserInsert(user model.User) (model.User,int, error) {
	var zUser model.User
	var authenticactionPolicy model.AuthenticactionPolicy
	var passwordPolicy model.PasswordPolicy
	if user.Uid == "" {
		return user,0,errors.New("用户名不能为空")
	}
	if user.Uname == "" {
		return user,0, errors.New("姓名不能为空")
	}
	if user.DepartmentId == 0 {
		return user,0, errors.New("部门不能为空")
	}
	util.DbConn.Table("z_user").Where("uid = ? and department_id != 3 ", user.Uid).Take(&zUser)
	if zUser.Uid == user.Uid {
		return user,0,errors.New("账户名已存在")
	}
	util.DbConn.Table("z_user").Where("uname = ? and department_id != 3 ", user.Uname).Take(&zUser)
	if zUser.Uname == user.Uname {
		return user,0,errors.New("姓名已存在")
	}
	util.DbConn.Table("z_user").Where("employee_number = ? and department_id != 3", user.EmployeeNumber).Take(&zUser)
	if user.EmployeeNumber != "" && zUser.EmployeeNumber == user.EmployeeNumber {
		return user,0,errors.New("员工号已存在")
	}
	util.DbConn.Table("z_user").Where("mobile = ? and department_id != 3 ", user.Mobile).Take(&zUser)
	if user.Mobile != "" && zUser.Mobile == user.Mobile {
		return user,0,errors.New("手机号已存在")
	}
	util.DbConn.Table("z_user").Where("cardid = ? and department_id != 3", user.Cardid).Take(&zUser)
	if user.Cardid != "" && zUser.Cardid == user.Cardid {
		return user,0,errors.New("身份证号已存在")
	}
	util.DbConn.Table("z_user").Where("mail = ? and department_id != 3", user.Mail).Take(&zUser)
	if user.Mail != "" && zUser.Mail == user.Mail {
		return user,0,errors.New("邮箱已存在")
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
	//如果未选择认证策略，则选择默认认证策略
	if user.UserAuthenPolicyId == 0 {
		user.UserAuthenPolicyId = -1
	}
	util.DbConn.Table("z_authenticaction_policy").Where("id = ?", user.UserAuthenPolicyId).Take(&authenticactionPolicy)
	if authenticactionPolicy.AuthenType == 2 {
		secret := cusfun.NewGoogleAuth().GetSecret()
		_, qrPng, _ := cusfun.NewGoogleAuth().CreateQrcode(user.Uid, user.Mobile, secret)
		user.AuthKey = secret
		user.QrCode = qrPng
	}
	//自动更新密码过期日期
	util.DbConn.Table("z_password_policy").Where("id = ?", user.UserPasswordPolicyId).Take(&passwordPolicy)
	nowTime := time.Now()
	newTime := nowTime.AddDate(0, 0, passwordPolicy.PolicyValidity).Format("2006-01-02")
	user.Expiretime = newTime
	count := util.DbConn.Create(&user).RowsAffected
	return user,int(count),nil
}

//查询密码策略
func UserPwdList(id int) (res model.PasswordPolicy) {
	util.DbConn.Table("z_password_policy").Where("id = ?", id).First(&res)
	return res
}

/*
修改用户
 */
func UserUpdate(user model.User) (model.User,int, error) {

	var zUser model.User
	var User model.User
	var authenticactionPolicy model.AuthenticactionPolicy
	var passwordPolicy model.PasswordPolicy
	//通过用户id查询修改前的用户信息
	util.DbConn.Table("z_user").Where("uid = ?", user.Uid).First(&User)
	if user.Uid == "" {
		return user,0,errors.New("用户名不能为空")
	}
	if user.Uname == "" {
		return user,0, errors.New("姓名不能为空")
	}
	if user.DepartmentId == 0 {
		return user,0, errors.New("部门不能为空")
	}
	util.DbConn.Table("z_user").Where("uname = ? and department_id != 3", user.Uname).Take(&zUser)
	if User.Uname != user.Uname && zUser.Uname == user.Uname {
		return user,0,errors.New("用户姓名已存在")
	}
	util.DbConn.Table("z_user").Where("employee_number = ? and department_id != 3", user.EmployeeNumber).Take(&zUser)
	if User.EmployeeNumber != user.EmployeeNumber {
		if user.EmployeeNumber != "" && zUser.EmployeeNumber == user.EmployeeNumber {
			return user,0,errors.New("员工号已存在")
		}
	}
	util.DbConn.Table("z_user").Where("mobile = ? and department_id != 3", user.Mobile).Take(&zUser)
	if User.Mobile != user.Mobile {
		if user.Mobile != "" && zUser.Mobile == user.Mobile {
			return user,0,errors.New("手机号已存在")
		}
	}
	util.DbConn.Table("z_user").Where("cardid = ? and department_id != 3", user.Cardid).Take(&zUser)
	if User.Cardid != user.Cardid {
		if user.Cardid != "" && zUser.Cardid == user.Cardid {
			return user,0,errors.New("身份证号已存在")
		}
	}
	util.DbConn.Table("z_user").Where("mail = ? and department_id != 3", user.Mail).Take(&zUser)
	if User.Mail != user.Cardid {
		if user.Mail != "" && zUser.Mail == user.Mail {
			return user,0,errors.New("邮箱已存在")
		}
	}
	if user.DepartmentId == 3 {
		return user,0, errors.New("已删除数据无法修改")
	}
	//根据身份日期生成出生日期
	if user.Cardid != "" {
		user.Birthdate = user.Cardid[6:14]
	}
	util.DbConn.Table("z_authenticaction_policy").Where("id = ?", user.UserAuthenPolicyId).Take(&authenticactionPolicy)
	if authenticactionPolicy.AuthenType == 2 {
		secret := cusfun.NewGoogleAuth().GetSecret()
		_, qrPng, _ := cusfun.NewGoogleAuth().CreateQrcode(user.Uid, user.Mobile, secret)
		user.AuthKey = secret
		user.QrCode = qrPng
	} else {
		user.AuthKey = ""
		user.QrCode = ""
	}
	//如果重置密码
	if user.Password != User.Password || user.UserPasswordPolicyId != User.UserPasswordPolicyId {
		//自动更新密码过期日期
		util.DbConn.Table("z_password_policy").Where("id = ?", user.UserPasswordPolicyId).Take(&passwordPolicy)
		nowTime := time.Now()
		newTime := nowTime.AddDate(0, 0, passwordPolicy.PolicyValidity).Format("2006-01-02")
		user.Expiretime = newTime
	}

	count := util.DbConn.Save(&user).RowsAffected
	return user, int(count),nil
}

/*
修改密码
 */
func ModifyPwd(userid, oldPwd, newPwd string)(model.User,int, error) {

	var user model.User
	var passwordPolicy model.PasswordPolicy
	//查询账号信息，有无账号，获得密码
	util.DbConn.Table("z_user").Where("uid = ?", userid).First(&user)
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
		return user,0, errors.New("旧密码错误")
	}
	if oriPwd == temporarynewPwd {
		return user,0, errors.New("新密码和旧密码不能相同")
	}
	if len(newPwd) < 8 {
		return user,0, errors.New("密码长度最少有8位")
	}
	pattern1 := "\\d+" //反斜杠要转义
	res1, _ := regexp.MatchString(pattern1, newPwd)
	if !res1 {
		return user,0, errors.New("密码必须包含数字")
	}
	pattern2 := "[a-zA-Z]"
	res2, _ := regexp.MatchString(pattern2, newPwd)
	if !res2 {
		return user,0, errors.New("密码必须包含字母")
	}
	pattern3 := "\\W+"
	res3, _ := regexp.MatchString(pattern3, newPwd)
	if !res3 {
		return user,0, errors.New("密码必须包含字符")
	}
	user.Password = temporarynewPwd
	//自动更新密码过期日期
	util.DbConn.Table("z_password_policy").Where("id = ?", user.UserPasswordPolicyId).Take(&passwordPolicy)
	nowTime := time.Now()
	newTime := nowTime.AddDate(0, 0, passwordPolicy.PolicyValidity).Format("2006-01-02")
	user.Expiretime = newTime

	util.DbConn.Table("z_user").Where("uid = ?",userid).Update("password",newPwd)
	return user,1,nil
}

/*
重置密码
 */
func ResetPwd(userid string) (model.User,int, error) {
	var user model.User
	var passwordPolicy model.PasswordPolicy
	util.DbConn.Table("z_user").Where("uid = ?", userid).First(&user)
	//加密
	user.Password = user.Uid
	var aes *util.AesEncrypt
	user.Password, _ = aes.Encrypt(user.Password)
	//自动更新密码过期日期
	util.DbConn.Table("z_password_policy").Where("id = ?", user.UserPasswordPolicyId).Take(&passwordPolicy)
	nowTime := time.Now()
	newTime := nowTime.AddDate(0, 0, passwordPolicy.PolicyValidity).Format("2006-01-02")
	user.Expiretime = newTime
	user.Modifypwddate = time.Now().Format("2006-01-02 15:04:05")
	user.Lastlogin = time.Now().Format("0000-00-00 00:00:00")
	util.DbConn.Table("z_user").Where("uid = ?",userid).Update("password",userid)
	return user,1,nil
}
//根据用户名查询用户信息
func UserFindById(userid string) (user model.User, err error, rows int) {
	var count int64
	util.DbConn.Table("z_user").Where("uid = ?", userid).First(&user)
	return user, err, int(count)
}


/*
根据用户账号逻辑删除用户
*/
func UserDelById(user model.User,userid string) (model.User,int,error) {
	var User model.User
	//通过用户id查询逻辑删除前的用户信息
	util.DbConn.Table("z_user").Where("uid = ?", user.Uid).First(&User)
	//初始化数据不允许删除
	if User.Islinlay == 1 {
		return User,1,errors.New("初始化数据不可被删除")
	}
	//当前登录账号不允许被删除
	if userid == User.Uid {
		return User,1,errors.New("当前登录用户不允许被删除")
	}
	user.Status = "0"
	user.DepartmentId = 3
	count := util.DbConn.Table("z_user").Where("uid = ?", user.Uid).Save(&user).RowsAffected
	return user,int(count),nil
}
/*
查询导出
*/
func UserFindEx(pojo cusfun.ParamsPOJO) (rst []model.Expdp, total int, err error) {
	var count int64
	db := util.DbConn.Table(`(SELECT z_user.* from z_user where isadmin != 1)t`)
	err = cusfun.GetSqlByParams(db, pojo, &count).Find(&rst).Error
	total = int(count)
	return
}

/*
查角色
*/
func RoleList(id int) (res model.Role) {
	util.DbConn.Table("z_role").Where("id = ?", id).First(&res)
	return res
}

/*
查部门
*/
func DeptList(id int) (res model.Department) {
	util.DbConn.Table("z_department").Where("id = ?", id).First(&res)
	return res
}