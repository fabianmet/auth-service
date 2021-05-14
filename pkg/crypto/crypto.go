package crypto

import (
	"github.com/fabianmet/auth-service/pkg/repository"
)

type JwtService interface {
	GeneratePrivateKey(bs int) error
	GeneratePublicKey()
	StorePrivateKey(r repository.Repository) error
}
