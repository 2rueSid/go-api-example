// package to implement types
package types

// DTO that describes which data
// should pass while createing new user
type CreateUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
