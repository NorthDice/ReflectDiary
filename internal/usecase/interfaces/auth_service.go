package interfaces

// AuthService defines authentication-related operations,
// including token generation and validation.
type AuthService interface {
	// GenerateToken creates a new authentication token for the given username.
	// Returns the generated token string or an error if token creation fails.
	GenerateToken(userID int) (string, error)

	// ValidateToken verifies the provided token string.
	// Returns the token if valid,
	// or an error if the token is invalid.
	ValidateToken(token string) (string, error)
}
