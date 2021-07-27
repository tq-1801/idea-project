package dao

import (
	"dsp/src/dsp_resource/model"
	"dsp/src/util"
	"dsp/src/util/cusfun"
	"fmt"
	"github.com/pkg/errors"

	"strconv"
	"strings"
)

/**
 * @Author: zhaojia
 * @Description:
 * @Version: 1.0.0
 * @Date: 2021/4/3
 */
//分页查询列表
func ListHostProtocol(pojo cusfun.ParamsPOJO) ([]model.HostProtocolPage, int, error) {
	var count int64
	var prosList []model.HostProtocolPage
	db := util.DbConn.Table("z_res_host_protocol protocol left join z_dictionary dictionary on dictionary.id=protocol.protocol_type").
		Select("protocol.id,protocol.protocol_port,protocol.protocol_type,protocol.res_id,dictionary.dic_name")
	db = cusfun.GetSqlByParams(db, pojo, &count).Find(&prosList)
	return prosList, int(count), nil
}

//添加协议管理
func CreateHostProtocol(resId int, pros []model.HostProtocol) (err error) {
	//开始事务
	tx := util.DbConn.Begin()
	//判断是否重复
	var reaptStr string
	for _, proReapt := range pros {
		reaptStr = reaptStr + strconv.Itoa(proReapt.ProtocolType) + "&" + strconv.Itoa(proReapt.ProtocolPort) + ";"
	}
	for _, proStr := range pros {
		num := strings.Count(reaptStr, strconv.Itoa(proStr.ProtocolType)+"&"+strconv.Itoa(proStr.ProtocolPort))
		if num > 1 {
			return errors.New("协议信息重复")
		}
	}
	//删除之前的
	err = tx.Table("z_res_host_protocol").Where("res_id", resId).Delete(&pros).Error
	//添加新的ip
	for _, pro := range pros {
		pro.ResId = resId
		err = tx.Create(&pro).Error
	}
	if err == nil {
		tx.Commit()
		fmt.Println("添加事务成功")
	} else {
		tx.Rollback()
		fmt.Println("回退事务成功")
	}
	return
}
