package entities

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	EMail    string `json:"email"`
	Password string `json:"password"`
}
