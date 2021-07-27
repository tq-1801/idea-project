package service

import (
	"dsp/src/dsp_user/dao"
	"dsp/src/dsp_user/model"
	"dsp/src/util/cusfun"
)

/**
 * @author  tianqiang
 * @date  2021/7/19 10:57
 */
func RoleFindList(pojo cusfun.ParamsPOJO) (rst interface{}, err error, count int) {

	res, count,err := dao.RoleFindList(pojo)
	return res, err, count
}

/*
新增角色
*/
func RoleInsert(role model.Role) (model.Role, error) {

	role, err := dao.RoleInsert(model.Role{
		Id:           role.Id,
		Rolename:     role.Rolename,
		ZDesc:        role.ZDesc,
		Isinlay:      "0",
		Createuserid: role.Createuserid,
	})
	return role, err
}


/*
修改
*/
func RoleModify(zRole model.Role) ( role model.Role,err error,count int) {

	role, count, err = dao.RoleUpdate(zRole)
	return role, err,count
}

/**
删除角色
*/
func RoleDel(role model.Role) (rst interface{}, count int, err error) {

	//删除角色
	role, count, err = dao.RoleDelById(role)
	//删除权限
	if err == nil {
		err = dao.FuncRelDeleteByGroupId(role.Id)
	}
	return role, count, err
}