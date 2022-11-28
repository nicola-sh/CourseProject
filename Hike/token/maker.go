package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// check if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
