package data

const (
	Guest = 0
	User = 1
	Moderator = 2
	Admin = 3
	Owner = 4
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
