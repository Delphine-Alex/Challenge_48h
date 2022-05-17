package models

type Organizer struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Referent string `json:"referent"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
