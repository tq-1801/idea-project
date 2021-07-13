package util

/********************************************************
 * res加解密
 * 设计：薛彬
 * 创建时间：2017-08-30
 ********************************************************/
import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

type AesEncrypt struct {
}

func (this *AesEncrypt) getKey() []byte {
	strKey := "fsferghywedwHYTJHYT"
	keyLen := len(strKey)
	if keyLen < 16 {
		panic("res key 长度不能小于16")
	}
	arrKey := []byte(strKey)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	//取前16个字节
	return arrKey[:16]
}

//加密字符串
func (this *AesEncrypt) Encrypt(strMesg string) (string, error) {
	log.Println("---加密前密码：" + strMesg)
	key := this.getKey()
	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return "", err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMesg))
	encodeString := base64.StdEncoding.EncodeToString(encrypted)
	log.Println("---加密后密码：" + encodeString)
	return encodeString, nil
}

//解密字符串
func (this *AesEncrypt) Decrypt(encodeString string) (strDesc string, err error) {
	src, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Println(err, "")
	}

	key := this.getKey()
	var iv = []byte(key)[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}
