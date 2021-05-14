package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/fabianmet/auth-service/pkg/types"
	"github.com/pascaldekloe/jwt"
)

// RsaKey is the instantiated key to generate JWT tokens with.
// Important method is GenerateJWT
type RsaKey struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

//NewRsaKey generates a new RSA key based on a 2048 byte length
func NewRsaKey() *RsaKey {
	key := generateRsaPrivateKey(2048)

	r := &RsaKey{
		privateKey: key,
		PublicKey:  &key.PublicKey,
	}
	return r
}

// generateRsaPrivateKey generates a private RSA key based on input bytes
func generateRsaPrivateKey(bs int) *rsa.PrivateKey {
	reader := rand.Reader

	// 2048 is a good one
	key, err := rsa.GenerateKey(reader, bs)
	if err != nil {
		panic(err)
	}

	return key
}

// PrintPublicPem returns the public key as []byte in pem format.
func (r *RsaKey) PrintPublicPem() []byte {
	pubKeyBytes := x509.MarshalPKCS1PublicKey(r.PublicKey)
	pubKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubKeyBytes,
		},
	)
	return pubKeyPem
}

//GenerateJWT uses the RSAKey and user input to generate a JWT token. It returns the token in []byte form
func (r *RsaKey) GenerateJWT(u *types.User) []byte {
	var claims jwt.Claims

	claimMap := make(map[string]interface{})

	// Add Subject.
	claims.Subject = u.Subject

	claimMap["email_verified"] = u.EmailVerified
	claimMap["given_name"] = u.GiveName
	claimMap["family_name"] = u.FamilyName
	claimMap["picture"] = u.Picture.String()
	claimMap["preferred_username"] = u.PreferredUserName
	claimMap["email"] = u.Email
	claimMap["customurlpathinc"] = u.Groups

	// Add common claims.
	claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
	claims.Issuer = "auth-service yo"
	claims.Set = claimMap
	// issue a JWT
	token, err := claims.RSASign(jwt.RS256, r.privateKey)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return token

}
