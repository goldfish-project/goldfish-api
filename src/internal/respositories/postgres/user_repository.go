package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"goldfish-api/internal/core/domain"
	"goldfish-api/internal/core/ports"
)

type UserRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) ports.UserRepository {
	return &UserRepository{db:db}
}

// get retrieves an user by its user id
func (repo *UserRepository) Get(userId uuid.UUID) (domain.User, error) {
	user := &User{UserId: userId}

	// select user from database by its id
	if err := repo.db.Model(user).WherePK().Select(); err != nil {
		return domain.User{}, err
	}

	return user.domain(), nil
}

// get retrieves an user by its email
func (repo *UserRepository) GetByEmail(email string) (domain.User, error) {
	user := &User{}

	// select user from database by its id
	if err := repo.db.Model(user).Where("? = ?", pg.Ident("email"), email).Select(); err != nil {
		return domain.User{}, err
	}

	return user.domain(), nil
}

// save saves an user to the database
func (repo *UserRepository) Save(user *domain.User) error {
	user.UserId = uuid.New()

	if _, err := repo.db.Model(user).Insert(); err != nil {
		return err
	}

	return nil
}

// update updates the user's information
func (repo *UserRepository) Update(user *domain.User) error {
	return nil
}