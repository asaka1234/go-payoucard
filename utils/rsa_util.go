package utils

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

func SignSHA256RSA(data []byte, privateKeyStr string) (string, error) {
	// 1. 解码Base64编码的私钥
	keyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 private key: %v", err)
	}

	// 2. 解析PKCS8格式的私钥
	privateKey, err := ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}

	// 3. 计算数据的SHA256哈希
	hashed := sha256.Sum256(data)

	// 4. 使用RSA私钥签名
	signature, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %v", err)
	}

	// 5. 返回Base64编码的签名
	return base64.StdEncoding.EncodeToString(signature), nil
}

func ParsePKCS8PrivateKey(keyBytes []byte) (*rsa.PrivateKey, error) {
	// 尝试直接解析PKCS8格式
	privKey, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err == nil {
		if rsaKey, ok := privKey.(*rsa.PrivateKey); ok {
			return rsaKey, nil
		}
		return nil, errors.New("key is not an RSA private key")
	}

	// 如果不是PKCS8格式，尝试解析PEM格式
	block, _ := pem.Decode(keyBytes)
	if block != nil {
		privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		if rsaKey, ok := privKey.(*rsa.PrivateKey); ok {
			return rsaKey, nil
		}
		return nil, errors.New("key is not an RSA private key")
	}

	return nil, errors.New("failed to parse private key")
}

//===================

// VerifySHA256RSA 使用RSA公钥验证SHA256withRSA签名
func VerifySHA256RSA(data []byte, publicKeyStr string, signStr string) (bool, error) {
	// 1. 解码Base64编码的公钥
	keyBytes, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		return false, fmt.Errorf("failed to decode base64 public key: %v", err)
	}

	// 2. 解析X509格式的公钥
	publicKey, err := ParseX509PublicKey(keyBytes)
	if err != nil {
		return false, fmt.Errorf("failed to parse public key: %v", err)
	}

	// 3. 解码Base64编码的签名
	signature, err := base64.StdEncoding.DecodeString(signStr)
	if err != nil {
		return false, fmt.Errorf("failed to decode base64 signature: %v", err)
	}

	// 4. 计算数据的SHA256哈希
	hashed := sha256.Sum256(data)

	// 5. 验证签名
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		if err == rsa.ErrVerification {
			return false, nil // 签名验证失败
		}
		return false, fmt.Errorf("signature verification error: %v", err)
	}

	return true, nil
}

// ParseX509PublicKey 解析X509格式的公钥
func ParseX509PublicKey(keyBytes []byte) (*rsa.PublicKey, error) {
	// 尝试直接解析X509格式
	pubKey, err := x509.ParsePKIXPublicKey(keyBytes)
	if err == nil {
		if rsaKey, ok := pubKey.(*rsa.PublicKey); ok {
			return rsaKey, nil
		}
		return nil, errors.New("key is not an RSA public key")
	}

	// 如果不是X509格式，尝试解析PEM格式
	block, _ := pem.Decode(keyBytes)
	if block != nil {
		pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		if rsaKey, ok := pubKey.(*rsa.PublicKey); ok {
			return rsaKey, nil
		}
		return nil, errors.New("key is not an RSA public key")
	}

	return nil, errors.New("failed to parse public key")
}
