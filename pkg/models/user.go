package models

type User struct {
	ID       int    `json:"-"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
}
