package cusfun

import (
	"math/rand"
	"time"
)

var (
	length  int
	charset string
)

func GeneratePassword(length int) string {
	//初始化密码切片
	var password = make([]byte, length, length)
	//源字符串1234567890
	var sourceStr string = "ABCDEFG1234567890HIJKab1234567890--==__==-_-1234567890_==-_==-_==-_-_==_cdefg012345hijklmn6789ABCDEFGHIJK1234567890LMNOP--==__==-_-_==-1234567890_==-_==-_-_==_QRSTUVWX1234567890sdferewe1234567890YZabcdefgh--==__==-_-_==-_==-_==-_-_==_ijklmnopq1234567890rsABCDEFGH--==__==-_-_==-_==-_==-_-_==_IJKLMNOtuvwxyz"

	//遍历，生成一个随机index索引,
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(sourceStr))
		password[i] = sourceStr[index]
	}
	return string(password)
}
