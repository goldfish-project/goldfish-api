package postgres

import (
	"github.com/go-pg/pg/v10"
	"goldfish-api/internal/core/domain"
	"goldfish-api/internal/core/ports"
)

type UserRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) ports.UserRepository {
	return &UserRepository{db:db}
}

// get retrieves an user by its email
func (repo *UserRepository) Get(email string) (domain.User, error) {
	return domain.User{}, nil
}

// save saves an user to the database
func (repo *UserRepository) Save(user *domain.User) error {
	if _, err := repo.db.Model(user).Insert(); err != nil {
		return err
	}

	return nil
}

// update updates the user's information
func (repo *UserRepository) Update(user *domain.User) error {
	return nil
}