package dao

import (
	"bufio"
	"bytes"
	"dsp/src/dsp_resource/model"
	"dsp/src/util"
	"dsp/src/util/cusfun"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

/**
 * @author  tianqiang
 * @date  2021/7/20 17:44
 */

//测试连通性
func TelnetRes(params map[string]interface{}) bool {
	resIpV4 := params["resIpV4"].(string)

	bool := isping(resIpV4)

	return bool
}

func isping(ip string) bool {
	var icmp model.ICMP
	//开始填充数据包
	icmp.Type = 8
	icmp.Code = 0
	icmp.Checksum = 0
	icmp.Identifier = 0
	icmp.SequenceNum = 0

	recvBuf := make([]byte, 32)
	var buffer bytes.Buffer

	//先在buffer中写入icmp数据报求去校验和
	binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.Checksum = CheckSum(buffer.Bytes())
	//然后清空buffer并把求完校验和的icmp数据报写入其中准备发送
	buffer.Reset()
	binary.Write(&buffer, binary.BigEndian, icmp)

	Time, _ := time.ParseDuration("2s")
	conn, err := net.DialTimeout("ip4:icmp", ip, Time)
	if err != nil {
		return false
	}
	_, err = conn.Write(buffer.Bytes())
	if err != nil {
		log.Println("conn.Write error:", err)
		return false
	}
	conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	num, err := conn.Read(recvBuf)
	if err != nil {
		log.Println("conn.Read error:", err)
		return false
	}

	conn.SetReadDeadline(time.Time{})

	if string(recvBuf[0:num]) != "" {
		return true
	}
	return false

}

func CheckSum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)

	return uint16(^sum)
}

//通过部门id获取部门名称
func GetDeptName(deptId int) string {
	var departName string
	var departList []model.Department
	util.DbConn.Where("id=?", deptId).Find(&departList)
	for _, depart := range departList {
		departName = depart.DepName
	}
	return departName
}

func AddIpRules(path string, store model.StoreInfo, maxport int) {

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	totLine := 0 //总共的行数

	allcontent := "" //读取的内容
	var buf = bytes.Buffer{}
	buf.WriteString(string(allcontent))
	for {
		content, isPrefix, err := reader.ReadLine()
		buf.WriteString(string(content) + "\n")
		//如果是第6行向里面加入添加内容
		if totLine == 6 {
			buf.WriteString("-A PREROUTING  -p tcp -m tcp --dport " + strconv.Itoa(maxport) + " -j DNAT --to-destination " + store.ResIpV4 + ":" + strconv.Itoa(store.Store.StorePort) + "\n")
			buf.WriteString("-A POSTROUTING  -p tcp -m tcp --dst " + store.ResIpV4 + " --dport " + strconv.Itoa(store.Store.StorePort) + " -j MASQUERADE" + "\n")
		}
		//当单行的内容超过缓冲区时，isPrefix会被置为真；否则为false；
		if !isPrefix {
			totLine++
		}
		if err == io.EOF {
			fmt.Println("一共有", totLine, "行内容")
			break
		}
	}
	WriteWithIoutil(path, buf.String())
}

//写文件内容
func WriteWithIoutil(name, content string) {

	//func WriteFile(filename string, data []byte, perm os.FileMode) error
	err := ioutil.WriteFile(name, []byte(content), 0666)
	fmt.Println(err)
}

func ResIsExists(pojo cusfun.ParamsPOJO) (err error, total int) {
	var count int64
	db := util.DbConn.Table("z_res").Where("status != 0")
	err = cusfun.GetSqlByParams(db, pojo, &count).Error
	//}
	total = int(count)
	return
}

//修改文档中的内容
func UpdateIpRules(path string, store model.StoreInfo, maxport int) {

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	totLine := 0 //总共的行数

	allcontent := "" //读取的内容
	var buf = bytes.Buffer{}
	buf.WriteString(string(allcontent))
	updateLine := -1
	for {
		content, isPrefix, err := reader.ReadLine()
		if strings.Contains(string(content), strconv.Itoa(store.Res.SecurePort)) {
			buf.WriteString("-A PREROUTING  -p tcp -m tcp --dport " + strconv.Itoa(maxport) + " -j DNAT --to-destination " + store.ResIpV4 + ":" + strconv.Itoa(store.Store.StorePort) + "\n")
			updateLine = totLine + 1
		} else if updateLine == totLine {
			buf.WriteString("-A POSTROUTING  -p tcp -m tcp --dst " + store.ResIpV4 + " --dport " + strconv.Itoa(store.Store.StorePort) + " -j MASQUERADE" + "\n")
		} else {
			buf.WriteString(string(content) + "\n")
		}
		//当单行的内容超过缓冲区时，isPrefix会被置为真；否则为false；
		if !isPrefix {
			totLine++
		}
		if err == io.EOF {
			fmt.Println("一共有", totLine, "行内容")
			break
		}
	}
	WriteWithIoutil(path, buf.String())
}

func ResStoreIsExists(ip, name, storetype string) (err error, total int) {
	var count int64
	db := util.DbConn.Table("z_res res left join z_res_store store on res.id=store.res_id ")
	if storetype == "130201" {
		db.Where("res.status != 0 and res.dic_type_id=1302  and res.res_ip_v4=? and store.store_sid=?", ip, name)
	} else {
		db.Where("res.status != 0 and res.dic_type_id=1302  and res.res_ip_v4=? and store.store_name=?", ip, name)
	}
	err = db.Count(&count).Error
	total = int(count)
	return
}

func ResAccountIsExists(resId, resAccount, accountId string) (err error, total int) {
	var count int64
	db := util.DbConn.Table("z_res_account").Where("res_id =? and res_account =? ", resId, resAccount)
	if accountId != "" {
		db.Where("id <> ?", accountId)
	}
	err = db.Count(&count).Error
	total = int(count)
	return
}
