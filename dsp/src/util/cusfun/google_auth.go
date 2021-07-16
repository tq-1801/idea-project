package cusfun

import (
	"encoding/base64"
	"github.com/skip2/go-qrcode"
)
import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

type GoogleAuth struct {
}

func NewGoogleAuth() *GoogleAuth {
	return &GoogleAuth{}
}

func (this *GoogleAuth) un() int64 {
	//fmt.Println(time.Now().UnixNano()/1000/30)
	return time.Now().UnixNano() / 1000 / 30
	//return time.Now().Unix()/30
}

func (this *GoogleAuth) hmacSha1(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	if total := len(data); total > 0 {
		h.Write(data)
	}
	return h.Sum(nil)
}

func (this *GoogleAuth) base32encode(src []byte) string {
	return base32.StdEncoding.EncodeToString(src)
}

func (this *GoogleAuth) base32decode(s string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(s)
}

func (this *GoogleAuth) toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func (this *GoogleAuth) toUint32(bts []byte) uint32 {
	return (uint32(bts[0]) << 24) + (uint32(bts[1]) << 16) +
		(uint32(bts[2]) << 8) + uint32(bts[3])
}

func (this *GoogleAuth) oneTimePassword(key []byte, data []byte) uint32 {
	hash := this.hmacSha1(key, data)
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := this.toUint32(hashParts)
	return number % 1000000
}

func (this *GoogleAuth) GetSecret() string {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, this.un())
	return strings.ToUpper(this.base32encode(this.hmacSha1(buf.Bytes(), nil)))
}

func (this *GoogleAuth) GetCode(secret string) ([]string, error) {
	codeArray := make([]string, 0)
	secretUpper := strings.ToUpper(secret)
	secretKey, err := this.base32decode(secretUpper)
	if err != nil {
		return codeArray, err
	}
	number1 := this.oneTimePassword(secretKey, this.toBytes((time.Now().Unix()-120)/60))
	codeArray = append(codeArray, fmt.Sprintf("%06d", number1))
	number2 := this.oneTimePassword(secretKey, this.toBytes((time.Now().Unix()-60)/60))
	codeArray = append(codeArray, fmt.Sprintf("%06d", number2))
	number3 := this.oneTimePassword(secretKey, this.toBytes(time.Now().Unix()/60))
	fmt.Println(time.Now().Unix() / 60)
	codeArray = append(codeArray, fmt.Sprintf("%06d", number3))
	number4 := this.oneTimePassword(secretKey, this.toBytes((time.Now().Unix()+60)/60))
	codeArray = append(codeArray, fmt.Sprintf("%06d", number4))
	number5 := this.oneTimePassword(secretKey, this.toBytes((time.Now().Unix()+120)/60))
	codeArray = append(codeArray, fmt.Sprintf("%06d", number5))

	fmt.Println("********************begin**********************")
	fmt.Println(codeArray)

	fmt.Println("********************end**********************")
	fmt.Println("")
	return codeArray, nil
}

func (this *GoogleAuth) GetQrcode(user, secret string) string {
	return fmt.Sprintf("otpauth://totp/%s?secret=%s", user, secret)
}

func (this *GoogleAuth) GetQrcodeUrl(user, secret string) string {
	qrcode := this.GetQrcode(user, secret)
	return fmt.Sprintf("http://www.google.com/chart?chs=200x200&chld=M%%7C0&cht=qr&chl=%s", qrcode)
}

func (this *GoogleAuth) VerifyCode(secret, code string) (bool, error) {
	rst := false
	_code, err := this.GetCode(secret)
	fmt.Println(_code, code, err)
	if err != nil {
		return false, err
	}
	for _, value := range _code {
		if strings.EqualFold(value, code) {
			rst = true
			break
		}
	}
	return rst, nil
	//return _code == code, nil
}

//创建二维码图片并返回base64编码字符串
func (this *GoogleAuth) CreateQrcode(user, mobile, secret string) (string, string, error) {

	//url := "otpauth://totp/liyinda.com?secret=" + secret + "&issuer=" + user+ "&mobile=" + mobile
	url := "otpauth://totp/somf?secret=" + secret + "&issuer=" + user + "&mobile=" + mobile

	//qrcode.WriteFile(url, qrcode.Medium, 256, "/Users/mxp/boco/somf/code_png/" + user + ".png")
	data, err := qrcode.Encode(url, qrcode.Medium, 256)
	//将存入的图片转换为base64格式
	//file, err := os.Open("/Users/mxp/boco/somf/code_png/" + user + ".png")
	if err != nil {
		fmt.Println("无法打开二维码图片")
	}
	//data, err := ioutil.ReadAll(file)
	//return secret, "data:image/png;base64," + string(Base64Encode(data)), err
	return secret, "data:image/png;base64," + base64.StdEncoding.EncodeToString(data), err

}
