package repository

import (
	"fmt"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

const usersErrorPrefix = "[users_repository]"

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (ur *UsersRepository) Create(user *models.User) (int, error) {
	tx, err := ur.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", usersErrorPrefix, err)
	}
	var userId int
	createUserQuery := fmt.Sprintf("INSERT INTO %s (name, email) VALUES ($1, $2) RETURNING id", usersTable)
	row := tx.QueryRow(createUserQuery, user.Name, user.Email)
	err = row.Scan(&userId)
	if err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("%s: %w", usersErrorPrefix, err)
	}
	return userId, tx.Commit()
}

func (ur *UsersRepository) Get(id int) (*models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	if err := ur.db.Get(&user, query, id); err != nil {
		return nil, fmt.Errorf("%s: %w", usersErrorPrefix, err)
	}
	return &user, nil
}

func (ur *UsersRepository) GetAll() ([]models.User, error) {
	var users []models.User
	query := fmt.Sprintf("SELECT * FROM %s", usersTable)
	if err := ur.db.Select(&users, query); err != nil {
		return nil, fmt.Errorf("%s: %w", usersErrorPrefix, err)
	}
	return users, nil
}

func (ur *UsersRepository) Update(id int, user *models.UserInput) error {
	setValues := make([]string, 0, 2)
	args := make([]interface{}, 0, 2)
	argsId := 1
	if user.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name = $%d", argsId))
		argsId++
		args = append(args, user.Name)
	}
	if user.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email = $%d", argsId))
		argsId++
		args = append(args, user.Email)
	}
	setQuery := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", usersTable, strings.Join(setValues, ", "), argsId)
	args = append(args, id)

	tx, err := ur.db.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", usersErrorPrefix, err)
	}
	_, err = tx.Exec(setQuery, args...)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("%s: %w", usersErrorPrefix, err)
	}
	return tx.Commit()
}

func (ur *UsersRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)
	_, err := ur.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", usersErrorPrefix, err)
	}
	return nil
}
