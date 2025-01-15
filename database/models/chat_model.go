package models

type Chat struct {
	UserId   string   `db:"user_id"`
	Messages []string `db:"brand"`
}
