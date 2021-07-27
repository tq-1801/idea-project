package dao

import (
	"bytes"
	"dsp/src/dsp_resource/model"
	"dsp/src/util"
	"dsp/src/util/cusfun"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"golang.org/x/crypto/ssh"
	"net"
	"strconv"
	"strings"
	"time"
)

/**
 * @author  tianqiang
 * @date  2021/7/20 17:44
 */

/*
分页查询账号管理列表
 */
func ListAccount(pojo cusfun.ParamsPOJO) (res []model.AccountPage, total int, err error) {
	var count int64
	db := util.DbConn.Table("(select id,res_id,res_account,res_account_name,res_account_password,res_account_status,res_account_yum,date_format(modify_password_date, '%Y-%m-%d %H:%i:%s') as modify_password_date,date_format(create_date, '%Y-%m-%d %H:%i:%s') as create_date,create_user_id,date_format(modify_date, '%Y-%m-%d %H:%i:%s') as modify_date,modify_user_id,z_desc,res_account_type,su_type,is_admin,is_super,if(is_admin=1,'是','否') as is_admin_name,if(res_account_status=1,'正常','锁定') as res_account_status_name,is_auto from z_res_account) t").Select("id,res_id,res_account,res_account_name,res_account_password,res_account_status,res_account_yum,modify_password_date,create_date,create_user_id,modify_date,modify_user_id,z_desc,res_account_type,su_type,is_admin,is_super,is_admin_name,res_account_status_name,is_auto")
	err = cusfun.GetSqlByParams(db, pojo, &count).Find(&res).Error
	total = int(count)
	return
}

func findProtocol(account model.Account) bool {
	find := false
	var protocolList []model.HostProtocol
	util.DbConn.Table("z_res_host_protocol").Where("res_id=?", account.ResId).Find(&protocolList)
	for _, protocol := range protocolList {
		if protocol.ProtocolType == 2004 {
			find = true
		}
	}
	return find
}

//添加账号
func CreateAccount(account model.Account)(model.Account,int,error){
	count := 1

	//如果是主机设备，判断加的是否是VNC账号
	var res model.Res
	util.DbConn.Table("z_res").Where("id = ? and status = ?",account.ResId,1 ).Take(&res)
	if res.DicCityId == 1301 {
		isFind := findProtocol(account)
		if account.ResAccountType == 2 && !isFind {
			return account, 0, errors.New("未关联vnc协议，不能添加桌面ID")
		}
	}

	//密码加密
	var aes *util.AesEncrypt
	if account.IsAuto == 1 {
		autoPass := cusfun.GeneratePassword(12)
		newPass,_ := aes.Encrypt(autoPass)
		account.ResAccountPassword = newPass
	}else {
		//根据策略判断密码是否符合标准
		passwordPolicyId := res.ResPasswordPolicyId
		passwordPolicy,_ := SelectPasswordPolicyById(passwordPolicyId)
		newPassword := account.ResAccountPassword

		//验证特殊字符等的位数是否满足
		if passwordPolicy != (model.PasswordPolicy{}) {
			if passwordPolicy.PolicyType == "模糊策略" {
				//如果是模糊策略，只验证长度即可
				lengthLess := strconv.Itoa(passwordPolicy.PolicyLengthLess)
				lengthMore := strconv.Itoa(passwordPolicy.PolicyLengthMore)
				if len(newPassword) > passwordPolicy.PolicyLengthMore || len(newPassword) < passwordPolicy.PolicyLengthLess {
					return account,0,errors.New("密码长度必须大于" + lengthLess + "位，小于" + lengthMore + "位")
				}
			}else{
				passLength := strconv.Itoa(passwordPolicy.PasswordLength) // 密码的长度
				if len(newPassword) != passwordPolicy.PasswordLength {
					return account,0,errors.New("密码必须" + passLength + "位数")
				}
				includeChar := passwordPolicy.PolicyIncludeChar // 包含的字符及位数
				charsList := strings.Split(includeChar,",")
				policyMap := make(map[string]string)
				if len(charsList) > 0 {
					for _, chars := range charsList {
						charsList := strings.Split(chars,":")
						policyMap[charsList[1]] = charsList[0]
					}
					fmt.Printf("policyMap:%v",policyMap)
				}
				pwd := newPassword
				wordCount := StrReplaceAllString(pwd)
				//判断数字是否符合
				if policyMap["1"] == "0" && wordCount.Number > 0 {
					return account,0,errors.New("密码中不能有数字")
				} else if policyMap["1"] != "0" {
					num,_ := strconv.Atoi(policyMap["1"])
					if wordCount.Number < num {
						return account,0,errors.New("密码至少有" + policyMap["1"] + "个数字")
					}
				}
				//判断小写是否符合
				if policyMap["a"] == "0" && wordCount.LowercaseLetters > 0 {
					return account,0,errors.New("密码中不能有小写字母")
				} else if policyMap ["a"] != "0" {
					num,_ := strconv.Atoi(policyMap["a"])
					if wordCount.LowercaseLetters < num {
						return account,0,errors.New("密码中至少有" + policyMap["a"] + "个小写字母")

					}
				}
				//判断大写是否符合
				if policyMap["A"] == "0" && wordCount.CapitalLetter > 0 {
					return account,0,errors.New("密码中不能有大写字母")
				} else if policyMap["A"] != "0" {
					num,_ := strconv.Atoi(policyMap["A"])
					if wordCount.CapitalLetter < num {
						return account,0,errors.New("密码至少有" + policyMap["A"] + "个大写字母")
					}
				}
				//判断特殊字符是否符合
				if policyMap["!"] == "0" && wordCount.OtherString > 0 {
					return account,0,errors.New("密码中不能有特殊字符")
				} else if policyMap["!"] != "0" {
					num,_ := strconv.Atoi(policyMap["!"])
					if wordCount.OtherString < num {
						return account,0,errors.New("密码中至少有" + policyMap["!"] + "个特殊字符")

					}
				}
			}
		}
		//修改密码
		newPass,_ := aes.Encrypt(account.ResAccountPassword)
		account.ResAccountPassword = newPass
	}
	err := util.DbConn.Create(&account).Error
	if err != nil {
		count = 0
	}
	return account,count,err

}

// 修改账号
func UpdateAccount(params map[string]interface{},account model.Account) (model.Account,int,error) {
	var count int64
	var newPassword string
	isEditPass := params["isEditPass"].(bool)
	if params["newPassword"] != nil {
		newPassword = params["newPassword"].(string)
	}
	var err error
	//AES工具类
	var aes *util.AesEncrypt
	var passwordPolicy model.PasswordPolicy
	//是否修改密码
	if isEditPass {
		//如果资产的密码和账号的密码同步选择的是，则要修改主机的密码，调用修改密码的接口
		var res model.Res

		util.DbConn.Table("z_res").Where("id = ? and status = 1",account.ResId).Take(&res)

		passwordPolicy,_ = SelectPasswordPolicyById(res.ResPasswordPolicyId)
		//验证特殊字符等的位数是否满足
		if passwordPolicy != (model.PasswordPolicy{}) {
			if passwordPolicy.PolicyType == "模糊策略" {
				//如果是模糊策略，只验证长度
				lengthLess := strconv.Itoa(passwordPolicy.PolicyLengthLess)
				lengthMore := strconv.Itoa(passwordPolicy.PolicyLengthMore)
				if len(newPassword) > passwordPolicy.PolicyLengthMore || len(newPassword) < passwordPolicy.PolicyLengthLess {
					return account,0,errors.New("密码长度必须大于" + lengthLess + "位，小于" + lengthMore + "位")
				}
			} else {
				passLength := strconv.Itoa(passwordPolicy.PasswordLength) //密码的长度
				if len(newPassword) != passwordPolicy.PasswordLength {
					return account,0,errors.New("密码必须" + passLength + "位数")
				}
				includeChar := passwordPolicy.PolicyIncludeChar // 包含的字符及位数
				charsList := strings.Split(includeChar,",")
				policyMap := make(map[string]string)
				if len(charsList) > 0 {
					for _, chars := range charsList {
						charsList := strings.Split(chars,":")
						policyMap[charsList[1]] = charsList[0]
					}
					fmt.Printf("policyMap:%v",policyMap)
				}
				pwd := newPassword
				wordCount := StrReplaceAllString(pwd)
				//判断数字是否符合
				if policyMap["1"] == "0" && wordCount.Number > 0 {
					return account,0,errors.New("密码中不能有数字")
				} else if policyMap["1"] != "0" {
					num,_ := strconv.Atoi(policyMap["1"])
					if wordCount.Number < num {
						return account,0,errors.New("密码中至少有" + policyMap["1"] + "个数字")
					}
				}
				//判断小写是否符合
				if policyMap["a"] == "0" && wordCount.LowercaseLetters > 0 {
					return account,0,errors.New("密码中不能有小写字母")
				} else if policyMap["a"] != "0" {
					num,_ := strconv.Atoi(policyMap["a"])
					if wordCount.LowercaseLetters < num {
						return account,0,errors.New("密码中至少有" + policyMap["a"] + "个小写字母")
					}
				}
				//判断大写是否符合
				if policyMap["A"] == "0" && wordCount.CapitalLetter > 0 {
					return account,0,errors.New("密码中不能有大写字母")
				} else if policyMap["A"] != "0" {
					num,_ := strconv.Atoi(policyMap["A"])
					if wordCount.CapitalLetter < num {
						return account,0,errors.New("密码中至少有" + policyMap["A"] + "个大写字母")
					}
				}
				//判断特殊字符是否符合
				if policyMap["!"] == "0" && wordCount.OtherString > 0 {
					return account,0,errors.New("密码中不能有特殊字符")
				} else if policyMap["!"] != "0" {
					num,_ := strconv.Atoi(policyMap["!"])
					if wordCount.OtherString < num {
						return account,0,errors.New("密码中至少有" + policyMap["!"] + "个特殊字符")
					}
				}
			}
		}
		if res.DicConnectId == 1 && account.ResAccountType == 1 {
			//判断系统的操作系统类型
			var icmp model.ICMP
			//开始填充数据包
			icmp.Type = 8
			icmp.Code = 0
			icmp.Checksum = 0
			icmp.Identifier = 0
			icmp.SequenceNum = 0

			recvBuf := make([]byte,32)
			var buffer bytes.Buffer

			//先在buffer中写入icmp数据报求去校验和
			binary.Write(&buffer,binary.BigEndian,icmp)

			Time,_ := time.ParseDuration("2s")
			conn,err := net.DialTimeout("ip4:icmp",res.ResIpV4,Time)

			defer conn.Close()
			if err != nil {
				return account,0,errors.New("连接出错")
			}
			_,err = conn.Write(buffer.Bytes())
			if err != nil {
				return account,0,errors.New("连接出错")
			}
			conn.SetReadDeadline(time.Now().Add(time.Second * 2))
			_,err = conn.Read(recvBuf)
			if err != nil  {
				return account,0,errors.New("连接出错")
			}
			if recvBuf[8] > 100 {
				return account,0,errors.New("不是Linux系统，赞不支持同步主机密码")
			}
			proto,_,_ := ListHostPortForSSH(res.Id)
			if proto.ProtocolPort == 0 {
				proto.ProtocolPort = 22
			}
			err2 := updateHostPassword(params,proto.ProtocolPort,res)
			if err2 != nil {
				return account,0,err2
			}
		}
		newpass,_ := aes.Encrypt(newPassword)
		account.ResAccountPassword = newpass
		err = util.DbConn.Save(&account).Error
		if err == nil {
			count = 1
		} else {
			count = 0
		}
	} else {
		err = util.DbConn.Save(&count).Error
		if err == nil {
			count = 1
		} else {
			count = 0
		}
	}
	return account,int(count),err
}

/*
删除账号
 */
func DeleteAccount(account model.Account) (model.Account,int,error)  {
	var count int
	//开始事务
	tx := util.DbConn.Begin()
	var customErr error
	var accErr error
	//删除账号的时候，如果有授权和自定义关联的一并删除
	var customList []ResAcc
	//删除关联的自定义数据
	tx.Table("z_res_custom_account").Where("acc_id = ?",account.Id).Find(&customList)
	for _,custom := range customList {
		customErr = tx.Table("z_res_custom_account").Delete(&custom).Error
	}
	//删除关联的授权信息
	var accList [] ResAcc
	tx.Table("z_auth_res_acc").Where("acc_id = ?",account.Id).Find(&accList)
	for _,acc := range accList {
		accErr = tx.Table("z_auth_res_acc").Delete(&acc).Error
	}
	fmt.Println(accErr)
	accountErr := tx.Delete(&account).Error
	if customErr == nil && accErr == nil && accountErr == nil {
		count = 1
		tx.Commit()
		fmt.Println("删除事务成功")
	} else {
		count = 0
		tx.Rollback()
		fmt.Println("删除事务回退")
	}
	return account,count,nil

}

/*
导出账号
 */
func ExportAccount(account model.Account) ([]model.ExportAccount,error,int) {
	var count int64
	var exportList [] model.ExportAccount
	db := util.DbConn.Table("z_res_account account left join z_dictionary d on d.id = account.res_account_type").
		Select("account.res_account,account.res_account_name,,if(account.res_account_status = 1,'正常','锁定') " +
			"as res_account_status,account.res_account_yum,DATE_FORMAT(account.create_date,'%Y-%m-%d') as create_date,d.dic_name as res_account_type_name, z_desc,su_type,if(account.is_admin=1,'是','否') as is_admin")
	if account.ResId != 0 {
		db = db.Where("res_id = ? ",account.ResId)
	}
	db.Count(&count).Order("account.id desc").Find(&exportList)
	return exportList,nil,int(count)

}

//导入资产账号
func CreateAccountMore(url string, resId int) ([]model.ImportErrorLogs, error, int) {
	var errorLogs = make([]model.ImportErrorLogs, 0)
	number := 0

	//取出所有的设备名称和IP
	var accounts = ""
	var names = ""
	var allAccount []model.Account
	util.DbConn.Table("z_res_account").Select("*").Where("res_id = ?", resId).Find(&allAccount).Order("id desc")
	if len(allAccount) > 0 {
		for i := 0; i < len(allAccount); i++ {
			accounts = accounts + allAccount[i].ResAccount + "@"
			names = names + allAccount[i].ResAccountName + "@"
		}
	}

	//取出主表最大值，赋值给number，用于生成资产的主键
	util.DbConn.Table("z_res_account").Select("max(id)").Find(&number)
	f, err := excelize.OpenFile(url)
	if err != nil {
		fmt.Println(err)
		return errorLogs, err, 0
	}

	var accountOne = new(model.Account)
	var account = make([]model.Account, 0)

	rows, err := f.GetRows("资产账号模板")
	//从1开始  第一行不要,for结束以后rows就是校验完的
	for i := 1; i < len(rows); i++ {
		can := true
		var logs = new(model.ImportErrorLogs)
		logs.Name = rows[i][0]
		logs.SecondMessage = rows[i][1]
		//必有字段有无校验
		if len(rows[i]) == 0 {
			can = false
			logs.ErrorLog += "空行 "
		} else {
			if rows[i][0] == "" {
				can = false
				logs.ErrorLog += "账号不能为空 "
			} else {
				if strings.Index(names, rows[i][0]) > -1 {
					can = false
					logs.ErrorLog += "账号重复 "
				} else {
					accounts = accounts + rows[i][0] + "@"
				}
			}
			if rows[i][1] == "" {
				can = false
				logs.ErrorLog += "账号名称不能为空 "
			} else {
				if strings.Index(names, rows[i][1]) > -1 {
					can = false
					logs.ErrorLog += "账号名称重复 "
				} else {
					names = names + rows[i][1] + "@"
				}
			}
			if rows[i][2] == "" {
				can = false
				logs.ErrorLog += "账号密码不能为空 "
			}
			if rows[i][3] == "" {
				can = false
				logs.ErrorLog += "账号状态不能为空 "
			}
			if rows[i][5] == "" {
				can = false
				logs.ErrorLog += "账号类型不能为空 "
			}
			if rows[i][8] == "" {
				can = false
				logs.ErrorLog += "是否管理员不能为空 "
			}
		}
		if can {
			number++
			accountOne.Id = number
			accountOne.ResId = resId
			accountOne.ResAccount = rows[i][0]
			accountOne.ResAccountName = rows[i][1]
			accountOne.ResAccountPassword = rows[i][2]
			if rows[i][3] == "锁定" {
				accountOne.ResAccountStatus = 0
			} else {
				accountOne.ResAccountStatus = 1
			}
			accountOne.ResAccountYum = rows[i][4]
			accountOne.ResAccountType, _ = strconv.Atoi(Substr(rows[i][5], strings.Index(rows[i][5], "_")))
			accountOne.ZDesc = rows[i][6]
			accountOne.SuType = rows[i][7]
			if rows[i][8] == "是" {
				accountOne.IsAdmin = 1
			} else {
				accountOne.IsAdmin = 0
			}
			accountOne.IsSuper = 0 //是否超级账号给了个默认值0
			account = append(account, *accountOne)
		}
		if logs.ErrorLog != "" {
			errorLogs = append(errorLogs, *logs)
		}
	}
	//保存信息
	count := util.DbConn.Table("z_res_account").Create(&account).RowsAffected
	return errorLogs, err, int(count)
}

//分页查询列表
func ListAccountByResId(resId string) (res []model.AccountPage, total int, err error) {
	var count int64
	db := util.DbConn.Table("(select id,res_id,res_account,res_account_name,res_account_password,res_account_status,res_account_yum,modify_password_date,date_format(create_date, '%Y-%m-%d %H:%i:%s') as create_date,create_user_id,date_format(modify_date, '%Y-%m-%d %H:%i:%s') as modify_date,modify_user_id,z_desc,res_account_type,su_type,is_admin,is_super,if(is_admin=1,'是','否') as is_admin_name,if(res_account_status=1,'正常','锁定') as res_account_status_name,is_auto from z_res_account) t").Select("id,res_id,res_account,res_account_name,res_account_password,res_account_status,res_account_yum,modify_password_date,create_date,create_user_id,modify_date,modify_user_id,z_desc,res_account_type,su_type,is_admin,is_super,is_admin_name,res_account_status_name,is_auto")
	err = db.Where("res_id=?", resId).Find(&res).Error
	total = int(count)
	return
}

//修改密码的接口
func updateHostPassword(params map[string]interface{}, port int, res model.Res) error {
	var err error
	newPassword := params["newPassword"].(string)
	resId := res.Id
	resAccount := params["resAccount"].(string)
	//AES工具类
	var aes *util.AesEncrypt
	//修改密码
	var account model.Account
	db := util.DbConn.Table("z_res_account").Where("res_id=? and res_account=? ", resId, resAccount)
	db.Take(&account)
	if (account == model.Account{}) {
		return errors.New("此账号不存在")
	}
	pass, err := aes.Decrypt(account.ResAccountPassword)
	if err != nil {
		return errors.New("密码解析失败")
	}
	addr := res.ResIpV4 + ":" + strconv.Itoa(port)
	// 建立SSH客户端连接
	client, err := ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User:            account.ResAccount,
		Auth:            []ssh.AuthMethod{ssh.Password(pass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		util.CustomLogger.Debug("主机连接失败")
		return errors.New("主机连接失败")
	}

	// 建立新会话
	session, err := client.NewSession()
	if err != nil {
		return errors.New("会话创建失败")
	}

	defer session.Close()

	var b bytes.Buffer
	passWordCmd := ""
	session.Stdout = &b
	if account.IsAdmin == 1 {
		passWordCmd = "bash -c \"echo -e '" + newPassword + "\n" + newPassword + "' | passwd\""
	} else {
		passWordCmd = "bash -c \"echo -e '" + pass + "\n" + newPassword + "\n" + newPassword + "' | passwd\""
	}
	util.CustomLogger.Debug(passWordCmd)
	var out bytes.Buffer
	var stderr bytes.Buffer
	session.Stdout = &out
	session.Stderr = &stderr
	if err := session.Run(passWordCmd); err != nil {
		util.CustomLogger.Debug("Stdout: %s\n", out.String())
		util.CustomLogger.Debug("Stderr: %s\n", stderr.String())
		if strings.Contains(stderr.String(), "simplistic") || strings.Contains(stderr.String(), "palindrome") {
			err = errors.New("新密码复杂度过低")
		}
		if strings.Contains(stderr.String(), "similar") {
			err = errors.New("新旧密码相似度太高")
		}
		if strings.Contains(stderr.String(), "shorter") {
			err = errors.New("新密码长度太短")
		}
		return err
	}

	return nil
}

type ResAcc struct {
	Id      int    `gorm:"primary_key" json:"id"`
	ResId   int    `json:"resId"`
	AccId   int    `json:"accId"`
	AccName string `json:"accName"`
	UserId  string `json:"userId"`
}

/*
根据密码策略主键查询密码策略信息
*/
func SelectPasswordPolicyById(policyId int) (passPolicy model.PasswordPolicy, err error) {
	err = util.DbConn.Where("id=?", policyId).Take(&passPolicy).Error
	return
}

//查询ssh协议的端口
func ListHostPortForSSH(resId int) (hostProtocol model.HostProtocol, total int, err error) {
	var count int64
	err = util.DbConn.Where("res_id = ?", resId).Where("protocol_type = ?", 2001).First(&hostProtocol).Count(&count).Error
	total = int(count)
	return
}

type StrReplaceStruct struct {
	CapitalLetter    int `json:"capital_letter"`
	LowercaseLetters int `json:"lowercase_letters"`
	Number           int `json:"number"`
	OtherString      int `json:"other_string"`
}

func StrReplaceAllString(s2 string) (strReplace StrReplaceStruct) {
	for i := strReplace.OtherString; i < len(s2); i++ {
		switch {
		case 64 < s2[i] && s2[i] < 91:
			strReplace.CapitalLetter += 1
		case 96 < s2[i] && s2[i] < 123:
			strReplace.LowercaseLetters += 1
		case 47 < s2[i] && s2[i] < 58:
			strReplace.Number += 1
		default:
			strReplace.OtherString += 1
		}
	}
	return strReplace
}