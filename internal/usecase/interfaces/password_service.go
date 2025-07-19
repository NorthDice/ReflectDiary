package interfaces

// PasswordService defines operations related to password hashing and verification.
type PasswordService interface {
	// HashPassword takes a plaintext password and returns its hashed representation.
	// Returns the hashed password or an error if hashing fails.
	HashPassword(password string) (string, error)

	// ComparePassword compares a hashed password with a plaintext password.
	// Returns true if they match, false otherwise, along with any error encountered.
	ComparePassword(hashedPassword string, password string) (bool, error)
}
