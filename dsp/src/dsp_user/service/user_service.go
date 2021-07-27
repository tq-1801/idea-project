package service

import (
	"dsp/src/dsp_user/dao"
	"dsp/src/dsp_user/model"
	"dsp/src/util/cusfun"
	"errors"
	"fmt"
	"github.com/pikanezi/mapslice"
	"github.com/tealeg/xlsx"
	"strconv"
	"strings"

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
func UserInsert(user model.User) (model.User, error) {

	res, _,err := dao.UserInsert(user)
	return res,err
}

/*
修改用户
 */
func UserUpdate(user model.User) (rst interface{},count int,err error)  {

	user, count, err = dao.UserUpdate(user)
	return user,count,err
}

/*
修改密码
 */
func ModifyPwd(userid, oldPwd, newPwd string) (rst interface{},err error,count int) {

	rst,count, err = dao.ModifyPwd(userid, oldPwd, newPwd ) //更新密码
	return rst,err,count
}

/*
重置密码
 */
func ResetPwd(userid string) (rst interface{},err error,count int) {

	rst,count, err = dao.ResetPwd(userid) //重置密码
	return rst,err,count
}

/*
逻辑删除用户
 */
func UserDel(user model.User,userid string) (zUser model.User,err error, count int)  {

	res, count, err := dao.UserDelById(user,userid)
	return res,err,count
	
}

/*
导入
*/
func Import(r io.ReaderAt, size int64) (int, error) {
	file, err := xlsx.OpenReaderAt(r, size)
	if err != nil {
		return 0, err
	}

	if len(file.Sheets) == 0 {
		return 0, errors.New("无数据")
	}
	//cells为一行的工作格
	//判断是否为空行
	isEmptyRow := func(cells []*xlsx.Cell) bool {
		cellsLen := len(cells)
		if cellsLen > 5 {
			cellsLen = 5
		}
		for i := 0; i < cellsLen; i++ {
			if len(cells[i].Value) > 0 {
				return false
			}
		}
		return true
	}

	//插入数据
	//防止下面插入数据时 i 瞎写
	getRowValue := func(cells []*xlsx.Cell, i int) string {
		if len(cells) <= i {
			return ""
		}
		return cells[i].Value
	}
	var count int
	//row 行
	//row.cells   列
	for i, row := range file.Sheets[0].Rows {
		if i == 0 && len(row.Cells) != 5 {
			err = errors.New("导入模板不符合要求")
			return count, err
		}else if i < 1 || isEmptyRow(row.Cells) {
			continue
		}else {
			pattern1 := "^[a-zA-Z0-9][a-zA-Z\\d\\-_/]+$" //反斜杠要转义
			res1, _ := regexp.MatchString(pattern1, getRowValue(row.Cells, 0))
			zifu := "./-_!@#$%^&*()~`\\|"
			if !res1 || strings.ContainsAny(getRowValue(row.Cells,0),zifu)|| IsNum(getRowValue(row.Cells,0)){
				return count, errors.New("用户账号"+getRowValue(row.Cells,0)+"的信息有误:"+"账户名由字母、数字或下划线或中横线组成,不能以特殊字符开头,不能纯数字")
			}
			pattern2 := "^[a-zA-Z\u4E00-\u9FA50-9][a-z\u4E00-\u9FA5A-Z\\d\\-_/·.]+$"
			res2, _ := regexp.MatchString(pattern2, getRowValue(row.Cells, 1))
			if !res2 ||strings.ContainsAny(getRowValue(row.Cells,0),zifu)|| IsNum(getRowValue(row.Cells,0)){
				return count, errors.New("用户账号"+getRowValue(row.Cells,0)+"的信息有误:"+"用户名由中英文、数字或者-_.·/等特殊字符组成,不能以特殊字符开头,不能纯数字")
			}
			pattern3 := "^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$"
			if len(getRowValue(row.Cells, 2)) != 0 {
				res3, _ := regexp.MatchString(pattern3, getRowValue(row.Cells, 2))
				if !res3 {
					return count, errors.New("用户账号"+getRowValue(row.Cells,0)+"的信息有误:"+"手机号码格式错误")
				}
			}
			if len(getRowValue(row.Cells, 4)) != 0 {
				pattern4 := "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
				res4, _ := regexp.MatchString(pattern4, getRowValue(row.Cells, 4))
				if !res4 {
					return count, errors.New("用户账号"+getRowValue(row.Cells,0)+"的信息有误:"+"不是有效的邮箱")
				}
			}
			_, err = UserInsert(model.User{
				Uid:            getRowValue(row.Cells, 0),
				Uname:          getRowValue(row.Cells, 1),
				DepartmentId:   1,
				Mobile:    		getRowValue(row.Cells, 2),
				EmployeeNumber: getRowValue(row.Cells, 3),
				Mail:           getRowValue(row.Cells, 4),
			})
			if err != nil {
				return count, errors.New("用户账号"+getRowValue(row.Cells,0)+"的信息有误:"+err.Error())
			}
			count++
		}
	}
	return count, nil
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
		expdp.Role = dao.RoleList(expdplist[i].Roleid).Rolename
		expdp.Mobile = expdplist[i].Mobile
		expdp.Department = dao.DeptList(expdplist[i].DepartmentId).DepName
		expdp.EmployeeNumber = expdplist[i].EmployeeNumber
		expdp.Mail = expdplist[i].Mail
		//组装
		expdplists = append(expdplists, expdp)
	}
	//取出数据
	uid, _ := mapslice.ToStrings(expdplists, "Uid")
	uname, _ := mapslice.ToStrings(expdplists, "Uname")
	role, _ := mapslice.ToStrings(expdplists, "Role")
	mobile, _ := mapslice.ToStrings(expdplists, "Mobile")
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
	cell = row.AddCell()
	cell.Value = "所属角色"
	cell = row.AddCell()
	cell.Value = "手机号码"
	cell = row.AddCell()
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
		cell = row.AddCell()
		cell.Value = role[i]
		cell = row.AddCell()
		cell.Value = mobile[i]
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

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}