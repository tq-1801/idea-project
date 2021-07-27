package dao

import (
	"dsp/src/dsp_user/model"
	"dsp/src/util"
	"dsp/src/util/cusfun"
	"errors"
)

/**
 * @author  tianqiang
 * @date  2021/7/19 10:29
 */

/*
查询角色列表
 */
func RoleFindList(pojo cusfun.ParamsPOJO) ( []model.Role,int, error) {
	var count int64
	var roleFindList []model.Role
	db := util.DbConn.Model(&model.Role{})
	db = cusfun.GetSqlByParams(db, pojo, &count).Find(&roleFindList)
	return roleFindList, int(count), nil
}

/*
角色表新增
*/
func RoleInsert(zRole model.Role) (model.Role, error) {
	var role model.Role

	if zRole.Rolename == "" {
		return role, errors.New("角色名称不能为空")
	}
	//查询角色名称是否重复
	util.DbConn.Table("z_role").Where("rolename = ? ", zRole.Rolename).Take(&role)
	if zRole.Rolename == role.Rolename {
		return role,errors.New("角色名称已存在")
	}

	err := util.DbConn.Create(&zRole).Error
	return zRole, err
}


/*
角色修改
*/
func RoleUpdate(zRole model.Role) (model.Role,int, error) {
	var all model.Role
	var role model.Role
	//查询修改之前的角色信息
	util.DbConn.Table("z_role").Where("id = ?",zRole.Id).Take(&all)
	//查询角色名称是否重复
	util.DbConn.Table("z_role").Where("rolename = ? ", zRole.Rolename).Take(&role)
	if all.Rolename != zRole.Rolename && role.Rolename == zRole.Rolename {
		return all,1, errors.New("角色名称已存在")
	}
	if zRole.Rolename == "" {
		return all,1, errors.New("角色名称不能为空")
	}

	count := util.DbConn.Table("z_role").Where("id = ?", zRole.Id).Save(&zRole).RowsAffected
	return zRole,int(count), nil
}

/*
角色删除
*/
func RoleDelById(zRole model.Role) (model.Role, int, error) {

	var user model.User
	var role model.Role
	//查询删除之前的角色信息
	util.DbConn.Table("z_role").Where("id = ?",zRole.Id).Take(&role)
	//有用户的角色不能被删除
	util.DbConn.Table("z_user").Where("roleid = ? and (status = ? or status = ? )", zRole.Id, 1, 2).Take(&user)
	if user.Roleid == zRole.Id {
		return role, 1, errors.New("本角色下包含用户,请先进行删除所属用户操作")
	}
	//默认角色不能删除
	if role.Isinlay == "1" {
		return role, 1, errors.New("默认角色无法删除")
	}
	//删除角色
	count := util.DbConn.Table("z_role").Where("id = ? ", role.Id).Delete(&model.Role{}).RowsAffected

	return role, int(count), nil
}


/*
删除权限
*/
func FuncRelDeleteByGroupId(id int) error {

	err := util.DbConn.Table("z_role_func").Where("roleid = ?", id).Delete(model.UgFunc{}).Error
	return err
}