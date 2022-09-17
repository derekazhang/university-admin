package tool

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type WxUserInfo struct {
	PurePhoneNumber string `json:"purePhoneNumber"`
	Watermark       struct {
		Appid string `json:"appid"`
	} `json:"watermark"`
}

func DecryptUserInfo(orgAppid, sessionKey, encryptData, iv string) (map[string]interface{}, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, err
	}
	sessionKeyBytes, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	dataBytes, err := AesDecrypt(decodeBytes, sessionKeyBytes, ivBytes)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &m)
	if err != nil {
		return nil, err
	}
	temp := m["watermark"].(map[string]interface{})
	appid := temp["appid"].(string)
	if appid != orgAppid {
		return nil, fmt.Errorf("invalid appid, get %s", appid)
	}
	if err != nil {
		return nil, err
	}
	return m, nil

}

func AesDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	//获取的数据尾端有'/x0e'占位符,去除它
	for i, ch := range origData {
		if ch == '\x0e' {
			origData[i] = ' '
		}
	}
	//{"phoneNumber":"15082726017","purePhoneNumber":"15082726017","countryCode":"86","watermark":{"timestamp":1539657521,"appid":"wx4c6c3ed14736228c"}}//<nil>
	return origData, nil
}

// func DecryptUserInfo() {
// 	src, err := Dncrypt(str, key, iv)
// 	fmt.Println(err)
// 	var s = map[string]interface{}{}
// 	json.Unmarshal([]byte(src), &s)
// 	fmt.Printf("== %+v", src)
// 	fmt.Printf("cc== %+v", s)
// }
