// User is the struct used to return an existing user
package models

type User struct {
	UUID          string `json:"uuid"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
}

// PostUser is the struct used to create a new user
type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// PostUser is the struct used to update an existing user
type PostUser struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}