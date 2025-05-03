package token

import "time"

// an interface for managing tokens
type Maker interface {
	// creates a new token for given username and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// checks if token is valid or not
	VerifyToken(token string) (*Payload, error)
}
