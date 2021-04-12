package crypto

import (
	"github.com/fabianmet/auth-service/pkg/repository"
)

type Crypto interface {
	GeneratePrivateKey(bs int) error
	GeneratePublicKey()
	StorePrivateKey(r repository.Repository) error
}
