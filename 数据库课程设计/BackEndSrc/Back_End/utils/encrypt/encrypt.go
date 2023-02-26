package encrypt

import (
	"Back_End/conf"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"
)

type KeyPair struct {
	PrivateKey string
	PubKey string
}

var seed int64

func init() {
	seed = time.Now().Unix()
}

func RefreshSeed() {
	seed = time.Now().Unix()
}

func GenerateKeyPair() (result KeyPair,err error) {
	fmt.Println("KeySize is : ",conf.GlobalConfig.AuthConf.KeySize)
	priKey,err := rsa.GenerateKey(rand.Reader,conf.GlobalConfig.AuthConf.KeySize)
	if err != nil {
		return KeyPair{},err
	}
	priBlock := pem.Block{
		Type: "RSA Private Key",
		Bytes: x509.MarshalPKCS1PrivateKey(priKey),
	}

	pubBlock := pem.Block{
		Type: "RSA Public Key",
		Bytes: x509.MarshalPKCS1PublicKey(&priKey.PublicKey),
	}

	priBytes := pem.EncodeToMemory(&priBlock)
	// fmt.Println("Pri:",string(priBytes))
	pubBytes := pem.EncodeToMemory(&pubBlock)
	// fmt.Println("Pub:",string(pubBytes))

	result.PubKey = string(pubBytes)
	result.PrivateKey = string(priBytes)
	return
}

func EncryptString(pubKeyStr, raw string) (ans string,err error) {
	pubKey,err := x509.ParsePKCS1PublicKey([]byte(pubKeyStr))
	if err != nil {
		return "",err
	}
	ansBytes,err := rsa.EncryptPKCS1v15(rand.Reader,pubKey,[]byte(raw))

	if err != nil {
		return "",err
	}

	ans = string(ansBytes)
	return
}

func DecodeString(priKeyStr, origin string) (ans string,err error) {
	priKey,err := x509.ParsePKCS1PrivateKey([]byte(priKeyStr))
	if err != nil {
		return "",err
	}
	ansBytes,err := rsa.DecryptPKCS1v15(rand.Reader,priKey,[]byte(origin))

	if err != nil {
		return "",err
	}

	ans = string(ansBytes)
	return
}