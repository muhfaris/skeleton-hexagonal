package model

import (
	"fmt"

	"github.com/google/uuid"
	svcmodel "github.com/muhfaris/skeleton-hexagonal/core/entities/account/service"
	"github.com/muhfaris/skeleton-hexagonal/transport/structures"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID           uuid.UUID `json:"id,omitempty"`
	FullName     string    `json:"full_name,omitempty"`
	Username     string    `json:"username,omitempty"`
	Email        string    `json:"email,omitempty"`
	Password     string    `json:"password,omitempty"`
	HashPassword string    `json:"hash_password,omitempty"`
	Role         string    `json:"role,omitempty"`
}

func CreateAccount(params *structures.SignUpRead) *Account {
	return &Account{
		ID:       uuid.New(),
		FullName: params.Fullname,
		Username: params.Username,
		Email:    params.Email,
		Password: params.Password,
		Role:     params.Role,
	}
}

func (a *Account) GenerateHashPassword() error {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	a.HashPassword = string(hashedPassword)
	return nil
}

func (a *Account) ComparePassword(password string) error {
	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)); err != nil {
		return fmt.Errorf("error: compare password, %v", err)
	}

	return nil
}

func (a *Account) Response() *svcmodel.AccountResponse {
	return &svcmodel.AccountResponse{
		ID:       a.ID,
		FullName: a.FullName,
		Username: a.Username,
		Email:    a.Email,
		Role:     a.Role,
	}
}
