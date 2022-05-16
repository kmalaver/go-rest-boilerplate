package test

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

type User struct {
	Name string
}

func TestA(t *testing.T) {
	testza.AssertTrue(t, true)
	testza.AssertEqual(t, User{Name: "Marvin"}, User{Name: "Marvin"})
}
