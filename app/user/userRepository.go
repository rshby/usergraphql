package user

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"reflect"
	"strconv"
	database "usergraphql/db"
)

type UserRepository struct {
	Db *sql.DB
}

// method to insert data user
func (u *UserRepository) Insert(ctx context.Context, entity *User) (*User, error) {
	stmt, err := database.Db.PrepareContext(ctx, "INSERT INTO users(Username, Password) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	log.Println(stmt)
	defer stmt.Close()

	// execute
	result, err := stmt.ExecContext(ctx, entity.Username, entity.Password)
	if err != nil {
		return nil, err
	}

	// get id inserted
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	entity.Id = strconv.Itoa(int(id))
	return entity, nil
}

// method get data by username
func (u *UserRepository) GetByUsername(ctx context.Context, username string) (*User, error) {
	stmt, err := database.Db.PrepareContext(ctx, "SELECT ID, Username, Password FROM users WHERE Username=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// execute
	result, err := stmt.QueryContext(ctx, username)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var user User
	if result.Next() {
		if err := result.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			return nil, err
		}
	}

	if reflect.DeepEqual(user, User{}) {
		return nil, errors.New("user not found")
	}

	// success
	return &user, nil
}

// method get all users
func (u *UserRepository) GetAll(ctx context.Context) ([]User, error) {
	stmt, err := database.Db.PrepareContext(ctx, "SELECT ID, Username, Password from users")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var users []User
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.New("record users not found")
	}

	return users, nil
}
