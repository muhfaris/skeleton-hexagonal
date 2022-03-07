package svcmodel

import "github.com/google/uuid"

type AccountResponse struct {
	ID          uuid.UUID `json:"id,omitempty"`
	FullName    string    `json:"full_name,omitempty"`
	Username    string    `json:"username,omitempty"`
	Email       string    `json:"email,omitempty"`
	Role        string    `json:"role,omitempty"`
	AccessToken string    `json:"access_token,omitempty"`
}
