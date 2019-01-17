package account

import "time"

type User struct {
	ID               string    `json:"_id"`
	Name             string    `json:"name"`
	Password         string    `json:"password"`
	PasswordHash     []byte    `json:"passwordhash"`
	Email            string    `json:"email"`
	PhoneNumber      string    `json:"phone_number"`
	AvatarURL        string    `json:"avatar_url"`
	IsActive         bool      `json:"is_active"`
	IsVerified       bool      `json:"is_verified"`
	Token            string    `json:"token"`
	TokenExpiryTime  time.Time `json:"token_expiry_time"`
	VerificationCode string    `json:"verification_code"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
