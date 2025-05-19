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

const (
	KeyAlgorithm             = "RSA"
	SignatureAlgorithmSHA256 = crypto.SHA256
)

// SignSHA256 signs the data using SHA256 with RSA
func SignSHA256(data []byte, privateKeyStr string) (string, error) {
	// Decode base64 private key
	keyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		return "", err
	}

	// Parse private key
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return "", errors.New("failed to parse PEM block containing the private key")
	}

	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	rsaPrivateKey, ok := privKey.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("key is not a valid RSA private key")
	}

	// Hash the data
	hashed := sha256.Sum256(data)

	// Sign the hashed data
	signature, err := rsa.SignPKCS1v15(nil, rsaPrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}

	// Return base64 encoded signature
	return base64.StdEncoding.EncodeToString(signature), nil
}

// Helper function to encode data to base64
func encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Helper function to decode base64 string
func decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

//========

// Verify verifies the signature of the given data using the public key
func Verify(data []byte, publicKeyStr string, sign string) (bool, error) {
	// Decode the base64 encoded public key
	keyBytes, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		return false, fmt.Errorf("failed to decode public key: %v", err)
	}

	// Parse the public key
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return false, errors.New("failed to parse PEM block containing the public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, fmt.Errorf("failed to parse public key: %v", err)
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return false, errors.New("public key is not an RSA key")
	}

	// Decode the signature
	signature, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, fmt.Errorf("failed to decode signature: %v", err)
	}

	// Hash the data
	hasher := sha256.New()
	hasher.Write(data)
	hashed := hasher.Sum(nil)

	// Verify the signature
	err = rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, hashed, signature)
	if err != nil {
		if err == rsa.ErrVerification {
			return false, nil // Signature is invalid
		}
		return false, fmt.Errorf("verification error: %v", err)
	}

	return true, nil
}
