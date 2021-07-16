package dao

import (
	"dsp/src/dsp_user/model"
	"dsp/src/util"
	"dsp/src/util/cusfun"
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
					SELECT 
						z.*,
						b.rolename As rolename
					from
					(SELECT
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
	count := util.DbConn.Create(&user).RowsAffected
	return user,int(count),nil
}

//根据用户账号查重
func ExistUserFindByUid(uid string)(count int64) {
	util.DbConn.Table("z_user").Where("uid",uid).Count(&count)
	return count
}

//根据用户姓名查重
func ExistUserFindByUname(uname string) (count int64) {
	util.DbConn.Table("z_user").Where("uname ",uname).Count(&count)
	return count

}

//身份证号查重
func ExistUserFindByCardId(cardid string) (rows int) {
	var count int64
	util.DbConn.Table("z_user").Where("cardid = ? ", cardid).Count(&count)
	return int(count)
}

//手机号码查重
func ExistUserFindByMobile(mobile string) (rows int) {
	var count int64
	util.DbConn.Table("z_user").Where("mobile = ? ", mobile).Count(&count)
	return int(count)
}

//员工编号查重
func ExistUserFindByEmpNum(employeeNumber string) (rows int) {
	var count int64
	util.DbConn.Table("z_user").Where("employee_number = ? ", employeeNumber).Count(&count)
	return int(count)
}

//邮箱查重
func ExistUserFindByMail(mail string) (rows int) {
	var count int64
	util.DbConn.Table("z_user").Where("mail = ? ", mail).Count(&count)
	return int(count)
}

//策略类型查询
func PolicyFindById(id int) (AuthenticactionPolicy model.AuthenticactionPolicy) {
	util.DbConn.Table("z_authenticaction_policy").Where("id = ?", id).First(&AuthenticactionPolicy)
	return AuthenticactionPolicy
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
	count := util.DbConn.Save(&user).RowsAffected
	return user, int(count),nil
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
func UserDelById(user model.User) (model.User,int,error) {
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