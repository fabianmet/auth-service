//package repository contains all backend repositories for the authentication service.
// working on redis and an in memory one.
// this package exposes the ways to crud rsa keys for token generation and validation, and user data.
package repository

import (
	"github.com/fabianmet/auth-service/pkg/repository/inmemory"
)

type Repository interface {
	Read(string) ([]byte, error)
	Delete(string) error
	Create(string, []byte) error
	UpdateKey(string, []byte) error
}

// interface checking!
var _ Repository = (*inmemory.InMemoryClient)(nil)
