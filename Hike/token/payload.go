package token

import (
	"time"

	"github.com/google/uuid"
)

// contains the payload data of the token
type Payload struct {
	ID        uuid.UUID
	Username  string
	IssuedAt  time.Time
	ExpiredAt time.Time
}

// creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}
