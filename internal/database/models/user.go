package models

type User struct {
	*Model
	Name     string
	Email    string
	Username string
}
