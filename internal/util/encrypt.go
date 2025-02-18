package util

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/square/go-jose.v2"
)

func HashMessage(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ValidateHash(hashedPassword, password string) error {
	// 验证密码
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func EncryptMessage(message string, secretKey []byte) (string, error) {
	// 创建 JWE 加密器
	encrypter, err := jose.NewEncrypter(
		jose.A256GCM,
		jose.Recipient{
			Algorithm: jose.DIRECT,
			Key:       secretKey,
		},
		nil,
	)
	if err != nil {
		return "", err
	}

	jwe, err := encrypter.Encrypt([]byte(message))
	if err != nil {
		return "", err
	}

	serializedJWE := jwe.FullSerialize()
	return serializedJWE, nil
}

// 解密
func DecryptMessage(encryptedJWT string, secretKey []byte) (string, error) {
	// 解析 JWE
	jwe, err := jose.ParseEncrypted(encryptedJWT)
	if err != nil {
		return "", err
	}

	// 解密 JWE
	decryptedBytes, err := jwe.Decrypt(secretKey)
	if err != nil {
		return "", err
	}
	return string(decryptedBytes), nil
}
