package models

const (
	RoleAdmin = Role("ADMIN")
	RoleAny   = Role("ANY")
)

type Role string
