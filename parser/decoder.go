package parser

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"io"
	"log"
	"regexp"
	"strings"

	"github.com/DisposaBoy/JsonConfigReader"

	"github.com/tidwall/gjson"
)

func getJson(url string) string {
	key := ""
	if strings.Contains(url, ";") {
		urlSplit := strings.Split(url, ";")
		key = urlSplit[2]
		url = urlSplit[0]
	}
	resp, err := parserClient.R().Get(url)
	if err != nil {
		//log.Fatal(err)
		return ""
	}
	dataWithOutComment := JsonConfigReader.New(strings.NewReader(resp.String()))
	buf := new(strings.Builder)
	_, _ = io.Copy(buf, dataWithOutComment)
	data := buf.String()
	if !gjson.Valid(data) {
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
	block, _ := aes.NewCipher(padEnd(key))
	ciphertext := decodeHex(data)
	mode := NewECBDecrypter(block)
	mode.CryptBlocks(ciphertext, ciphertext)
	return string(pkcs5Trimming(ciphertext))
}

func cbcDecrypt(data string) string {
	indexKey := strings.Index(data, "2324") + 4
	key := string(decodeHex(data[:indexKey]))
	key = strings.Replace(string(key), "$#", "", -1)
	key = strings.Replace(string(key), "#$", "", -1)
	indexIv := len(data) - 26
	iv := decodeHex(strings.TrimSpace(data[indexIv:]))
	ciphertext := decodeHex(strings.TrimSpace(data[indexKey:indexIv]))
	block, _ := aes.NewCipher(padEnd(key))
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
		log.Fatalln(err)
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
		log.Fatal(err)
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
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
