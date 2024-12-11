package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type HashedPassword struct {
	Hash string `bson:"hash"`
	Salt string `bson:"salt"`
}

func GeneratePassword(rawPassword string) (*HashedPassword, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("generate salt failed: %w", err)
	}

	saltedPassword := append([]byte(rawPassword), salt...)

	hash, err := bcrypt.GenerateFromPassword(saltedPassword, bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password failed: %w", err)
	}

	return &HashedPassword{
		Hash: base64.StdEncoding.EncodeToString(hash),
		Salt: base64.StdEncoding.EncodeToString(salt),
	}, nil
}

func VerifyPassword(storedPassword *HashedPassword, inputPassword string) error {
	hash, err := base64.StdEncoding.DecodeString(storedPassword.Hash)
	if err != nil {
		return fmt.Errorf("decode hash failed: %w", err)
	}

	salt, err := base64.StdEncoding.DecodeString(storedPassword.Salt)
	if err != nil {
		return fmt.Errorf("decode salt failed: %w", err)
	}

	saltedPassword := append([]byte(inputPassword), salt...)

	return bcrypt.CompareHashAndPassword(hash, saltedPassword)
}
