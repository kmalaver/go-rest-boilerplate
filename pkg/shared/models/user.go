package models

import (
	"strings"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required,min=18,max=99"`
	Other string `json:"other"`
}

// func (u *User) Validate() error {

// }

func (u *User) Sanitize() {
	u.Name = strings.TrimSpace(u.Name)
	u.Name = strings.ToLower(u.Name)

	u.Email = strings.TrimSpace(u.Email)
	u.Email = strings.ToLower(u.Email)
}
