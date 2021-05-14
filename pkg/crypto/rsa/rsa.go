package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"

	"github.com/fabianmet/auth-service/pkg/repository"
)

type JwtRSA struct {
	privateKey *rsa.PrivateKey
}

func (j *JwtRSA) GeneratePrivateKey(bs int) error {
	reader := rand.Reader

	key, err := rsa.GenerateKey(reader, bs)
	if err != nil {
		return err
	}
	j.privateKey = key

	byteKey := x509.MarshalPKCS1PrivateKey(key)
	fmt.Println(byteKey)
}

func (j *JwtRSA) GeneratePublicKey() {

}

func (j *JwtRSA) StorePrivateKey(r repository.Repository) error {

}
