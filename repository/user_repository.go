package repository

import (
	"SaveMate/models/user"
	"database/sql"
)

type UserRepository interface {
	CreateUser(user *user.User) (*user.User, error)
	// FindByUserId(UserId int) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *user.User) (*user.User, error) {
	query := "INSERT INTO USERS (id, user_id, username, email, password, role, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?)"

	result, err := r.db.Exec(query, user.Id, user.UserId, user.Username, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return user, err
	}

	userId, _ := result.LastInsertId()
	user.Id = int(userId)

	return user, nil
}

// func (r *userRepository) FindByUserId(Id int) (*user.User, error) {
// 	query := "SELECT id, user_id, username, email, role FROM USERS WHERE id = ? "

// 	row := r.db.QueryRow(query, Id)

// 	user := &user.User{}
// 	err := row.Scan(&user.Id, &user.UserId, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

func (r *userRepository) FindByEmail(email string) (*user.User, error) {
	query := "SELECT id, user_id, username, email, role FROM users WHERE email = ? "

	row := r.db.QueryRow(query, email)

	user := &user.User{}
	err := row.Scan(&user.Id, &user.UserId, &user.Username, &user.Email, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return user, err
	}
	return user, nil
}
