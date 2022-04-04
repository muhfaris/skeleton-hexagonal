package rslvmodel

import (
	"github.com/google/uuid"
	"github.com/graph-gophers/graphql-go"
)

type AccountResolver struct {
	Id          uuid.UUID `json:"id,omitempty"`
	Fullname    string    `json:"full_name,omitempty"`
	Username    string    `json:"username,omitempty"`
	Email       string    `json:"email,omitempty"`
	AccessToken string    `json:"access_token,omitempty"`
}

// ID is to resolve id field
func (account *AccountResolver) ID() *graphql.ID {
	return &graphql.ID(account.Id.String())
}

func (a *AccountResolver) FullName() string {
	return a.Fullname
}

func (account *AccountResolver) AccountResponse() {

}

/*
func (a *AccountResolver) ID() graphql.ID {
	return a.Id
}

func (a *AccountResolver) Username() string {
	return a.Username
}

func (a *AccountResolver) Email() string {
	return a.Email
}

func (a *AccountResolver) AccessToken() string {
	return a.AccessToken
}
*/
