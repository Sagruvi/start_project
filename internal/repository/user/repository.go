package repository_user

import (
	"context"
	model_user "start/internal/model/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func New(conn *pgxpool.Pool) *userRepository {
	return &userRepository{db: conn}
}

func (u *userRepository) GetUser(ctx context.Context, id int) (model_user.User, error) {
	query := "SELECT id, name, nickname FROM users WHERE id = $1"
	row := u.db.QueryRow(ctx, query, id)

	var user model_user.User
	if err := row.Scan(&user.ID, &user.Name, &user.NickName); err != nil {
		if err.Error() == "no rows in result set" {
			return model_user.User{}, nil // No user found
		}
		return model_user.User{}, err // Other error
	}

	return user, nil
}

func (u *userRepository) CreateUser(ctx context.Context, user model_user.User) (model_user.User, error) {
	query := "INSERT INTO users (name, nickname) VALUES ($1, $2) RETURNING id"
	row := u.db.QueryRow(ctx, query, user.Name, user.NickName)

	var id int
	if err := row.Scan(&id); err != nil {
		return model_user.User{}, err // Error inserting user
	}

	user.ID = id // Set the ID of the newly created user
	return user, nil
}

func (u *userRepository) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := u.db.Exec(ctx, query, id)
	if err != nil {
		return err // Error deleting user
	}
	return nil // User deleted successfully
}

func (u *userRepository) UpdateUser(ctx context.Context, user model_user.User) (model_user.User, error) {
	query := "UPDATE users SET name = $1, nickname = $2 WHERE id = $3"
	_, err := u.db.Exec(ctx, query, user.Name, user.NickName, user.ID)
	if err != nil {
		return model_user.User{}, err // Error updating user
	}
	return user, nil // User updated successfully
}
