package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseByteArray(byteString string) []byte {

	byteStrArr := strings.Split(byteString, ",")
	byteValues := make([]byte, len(byteStrArr))

	for i, byteStr := range byteStrArr {
		byteVal, err := strconv.Atoi(byteStr)
		if err != nil {
			fmt.Println("Error parsing byte value:", err)
			continue
		}
		byteValues[i] = byte(byteVal)
	}

	return byteValues
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Encrypt(text string) (string, error) {
	mySecretKey := os.Getenv("ENCRYPT_SECRET_KEY")
	myIVKey := os.Getenv("ENCRYPT_IV_KEY")
	bytes := ParseByteArray(myIVKey)

	block, err := aes.NewCipher([]byte(mySecretKey))
	if err != nil {
		return "", err
	}

	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func Decrypt(text string) (string, error) {
	mySecretKey := os.Getenv("ENCRYPT_SECRET_KEY")
	myIVKey := os.Getenv("ENCRYPT_IV_KEY")
	bytes := ParseByteArray(myIVKey)

	block, err := aes.NewCipher([]byte(mySecretKey))
	if err != nil {
	 return "", err
	}

	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
 }