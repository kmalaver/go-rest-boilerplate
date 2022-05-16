package user

import (
	"rest/pkg/shared/database"
	"rest/pkg/shared/models"

	"github.com/gocraft/dbr/v2"
)

type userRepository struct {
	db *database.Database
}

func newUserRepository() *userRepository {
	db := database.GetDatabase()
	return &userRepository{
		db,
	}
}

func (r *userRepository) List(filters dbr.Builder) ([]models.User, error) {
	var users []models.User
	_, err := r.db.Select("*").From("users").Where(
		dbr.And(
			dbr.Eq("name", "a"),
			dbr.Gt("age", 20),
		),
	).Load(&users)
	return users, err
}
