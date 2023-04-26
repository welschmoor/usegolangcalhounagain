package models

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	FirstName    string
	LastName     string
	PasswordHash string
	Dob          time.Time
}

type UserService struct {
	DB *pgx.Conn
}

func (us *UserService) CreateUser(email,
	first_name,
	last_name,
	password string,
	dob time.Time,
) (*User, error) {

	email = strings.ToLower(email)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("create error: %w", err)
	}
	passwordHash := string(hashedBytes)

	row := us.DB.QueryRow(
		context.Background(),
		`INSERT INTO users (
			email,
			first_name,
			last_name,
			password_hash,
			dob,
			activated,
			banned
		)
		VALUES ( $1, $2, $3, $4, $5, $6, $7)
		RETURNING id;`,
		email,
		first_name,
		last_name,
		passwordHash,
		dob,
		false,
		false,
	)

	user := User{
		Email:     email,
		FirstName: first_name,
		LastName:  last_name,
		Dob:       dob,
	}

	err = row.Scan(
		&user.ID,
	)
	if err != nil {
		return nil, fmt.Errorf("save to db err: %w", err)
	}

	return &user, nil
}

func (us *UserService) UpdateUser(email,
	first_name,
	last_name,
	password_hash string,
	dob time.Time,
	activated,
	banned bool) (*User, error) {

	return &User{}, nil
}
