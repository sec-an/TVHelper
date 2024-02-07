package parser

import (
	"TVHelper/global"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"io"
	"regexp"
	"strings"

	"go.uber.org/zap"

	"github.com/DisposaBoy/JsonConfigReader"

	"github.com/tidwall/gjson"
)

func getJson(url string) string {
	key := ""
	if strings.Contains(url, ";") {
		if urlSplit := strings.Split(url, ";"); len(urlSplit) > 2 {
			key = urlSplit[2]
			url = urlSplit[0]
		}
	}
	resp, err := global.ParserClient.R().Get(url)
	if err != nil {
		global.Logger.Error(url, zap.Error(err))
		return ""
	}
	dataWithOutComment := JsonConfigReader.New(strings.NewReader(resp.String()))
	buf := new(strings.Builder)
	_, err = io.Copy(buf, dataWithOutComment)
	if err != nil {
		global.Logger.Error(url+":json注释处理出错", zap.Error(err))
		return ""
	}
	data := buf.String()
	if !gjson.Valid(data) {
		data = resp.String()
		if strings.Contains(data, "**") {
			data = base64ToString(data)
		}
		if strings.HasPrefix(data, "2423") {
			data = cbcDecrypt(data)
		}
		if key != "" {
			data = ecbDecrypt(data, key)
		}
	}
	return strings.Replace(data, "./", url[:strings.LastIndex(url, "/")+1], -1)
}

func ecbDecrypt(data string, key string) string {
	block, err := aes.NewCipher(padEnd(key))
	if err != nil {
		global.Logger.Error(key,
			zap.String("data", data),
			zap.Error(err))
		return ""
	}
	ciphertext := decodeHex(data)
	mode := NewECBDecrypter(block)
	mode.CryptBlocks(ciphertext, ciphertext)
	return string(pkcs5Trimming(ciphertext))
}

func cbcDecrypt(data string) string {
	indexKey := strings.Index(data, "2324") + 4
	key := string(decodeHex(data[:indexKey]))
	key = strings.Replace(key, "$#", "", -1)
	key = strings.Replace(key, "#$", "", -1)
	indexIv := len(data) - 26
	iv := decodeHex(strings.TrimSpace(data[indexIv:]))
	ciphertext := decodeHex(strings.TrimSpace(data[indexKey:indexIv]))
	block, err := aes.NewCipher(padEnd(key))
	if err != nil {
		global.Logger.Error(key,
			zap.String("data", data),
			zap.Error(err))
		return ""
	}
	mode := cipher.NewCBCDecrypter(block, padEnd(string(iv)))
	mode.CryptBlocks(ciphertext, ciphertext)
	return string(pkcs5Trimming(ciphertext))
}

func base64ToString(data string) string {
	extracted := extract(data)
	if extracted == "" {
		return data
	}
	decodeBytes, err := base64.StdEncoding.DecodeString(extracted)
	if err != nil {
		global.Logger.Error(data, zap.Error(err))
	}
	return string(decodeBytes)
}

func extract(data string) string {
	index := regexp.MustCompile(`[A-Za-z0-9]{8}\*\*`).FindStringIndex(data)
	if index != nil {
		return data[index[1]:]
	}
	return ""
}

func padEnd(key string) []byte {
	return []byte(key + "0000000000000000"[len(key):])
}

func decodeHex(s string) []byte {
	data, err := hex.DecodeString(strings.ToUpper(s))
	if err != nil {
		global.Logger.Error(s, zap.Error(err))
		return []byte{}
	}
	return data
}

func pkcs5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

type ECB struct {
	b         cipher.Block
	blockSize int
}

func NewECB(b cipher.Block) *ECB {
	return &ECB{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ECBDecrypter ECB

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ECBDecrypter)(NewECB(b))
}
func (x *ECBDecrypter) BlockSize() int {
	return x.blockSize
}
func (x *ECBDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		global.Logger.Error("crypto/cipher: input not full blocks")
		return
	}
	if len(dst) < len(src) {
		global.Logger.Error("crypto/cipher: output smaller than input")
		return
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
