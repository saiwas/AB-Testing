package model

// User table
type User struct {
	ID             uint64
	Email          string
	PasswordDigest string
	AccessToken    string
}
