package user_entity

import "time"

type User struct {
	Id        *string
	username  string
	password  string
	email     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
