package router

import (
	"github.com/gin-gonic/gin"
	. "dsp/src/dsp_user/controller"
)

func Routers(router *gin.RouterGroup) {

	/*
		用户
	*/
	usercluster := router.Group("user")
	{
		usercluster.POST("/list", UserList)
		usercluster.POST("/create", UserAdd)
		usercluster.POST("/update", UserUpdate)
		usercluster.POST("/lock_update", UserLockUpdate)
		usercluster.POST("/delete", UserDel)
		usercluster.POST("/pwd_update", ModifyPwd)
		usercluster.POST("/pwd_reset", ResetPwd)
		usercluster.POST("/import", Import)
		usercluster.POST("/export", Export)
	}

	/*
		角色
	*/
	rolecluster := router.Group("role")
	{
		rolecluster.POST("/list", RoleList)
		rolecluster.POST("/create", RoleAdd)
		rolecluster.POST("/update", RoleUpdate)
		rolecluster.POST("/delete", RoleDel)
	}

	/*
	   部门
	*/
	departmentluster := router.Group("dep")
	{
		departmentluster.POST("find/list", DepartmentFindList)
		departmentluster.POST("/recursion/list", DepartmentTreeList)
		departmentluster.POST("/list", DepartmentFindByIdList)
		departmentluster.POST("/create", DepartmentAdd)
		departmentluster.POST("/update", DepartmentUpdate)
		departmentluster.POST("/delete", DepartmentDel)
	}
}
