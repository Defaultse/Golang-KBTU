package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"user-api/internal/models"
	"user-api/internal/store"
)

func (db *DB) User() store.UserRepository {
	if db.users == nil {
		db.users = NewUsersRepository(db.conn)
	}

	return db.users
}

type UsersRepository struct {
	conn *sqlx.DB
}

func NewUsersRepository(conn *sqlx.DB) store.UserRepository {
	return &UsersRepository{conn: conn}
}

// CRUD

func (u UsersRepository) Create(ctx context.Context, user *models.User) error {
	_, err := u.conn.Exec("INSERT INTO users VALUES ($1, $2, $3)", user.Username, user.Email, user.PhoneNumber)
	if err != nil {
		return err
	}

	return nil
}

func (u UsersRepository) All(ctx context.Context) ([]*models.User, error) {
	users := make([]*models.User, 0)
	if err := u.conn.Select(&users, "SELECT * FROM users"); err != nil {
		return nil, err
	}

	return users, nil
}

func (u UsersRepository) ByID(ctx context.Context, id int) (*models.User, error) {
	user := new(models.User)
	if err := u.conn.Get(user, "SELECT id, name FROM users WHERE id=$1", id); err != nil {
		return nil, err
	}

	return user, nil
 }

func (u UsersRepository) Update(ctx context.Context, user *models.User) error {
	_, err := u.conn.Exec("UPDATE users SET username = $1 WHERE id=$2", user.Username, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u UsersRepository) Delete(ctx context.Context, id int) error {
	_, err := u.conn.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}



