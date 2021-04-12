package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"

	"github.com/fabianmet/auth-service/pkg/repository"
)

func (us *userService) ensureRSAKey() error {
	//Check if RSA key exists in redis
	n, err := us.rc.Exists(rsaKey).Result()
	if err != nil {
		return err
	}
	if n != 1 {
		// For now also stores the key.
		us.generateRSAKey(rsaKeySize)
	}
	byteKey, err := us.rc.Get(rsaKey).Bytes()
	if err != nil {
		fmt.Println(err)
	}
	us.RSAKey, err = x509.ParsePKCS1PrivateKey(byteKey)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (us *userService) generateRSAKey(bs int) error {

	return nil, nil
}

type JwtRSA struct {
	privateKey *rsa.PrivateKey
}

func (j *JwtRSA) GeneratePrivateKey(bs int) error {
	reader := rand.Reader

	key, err := rsa.GenerateKey(reader, bs)
	if err != nil {
		return nil, err
	}
	us.RSAKey = key

	// TODO: should this be in the ensure function?
	byteKey := x509.MarshalPKCS1PrivateKey(key)

	us.rc.Set(rsaKey, byteKey, 0)
}

func (j *JwtRSA) GeneratePublicKey() {

}

func (j *JwtRSA) StorePrivateKey(r repository.Repository) error {

}
