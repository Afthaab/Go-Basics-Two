package models

type User struct {
	Name       string
	Email      string
	Password   string
	Permission map[string]bool
	Hobbies    []string
}
