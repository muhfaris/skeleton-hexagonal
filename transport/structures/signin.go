package structures

// LoginRead is wrap data of login
type LoginRead struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type DataLoginRead struct {
	Data *LoginRead `json:"data,omitempty"`
}
