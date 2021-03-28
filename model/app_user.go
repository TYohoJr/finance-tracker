package model

import (
	"finance-tracker/auth"
	"fmt"
	"time"
)

type AppUser struct {
	ID          *int       `db:"id" json:"id"`
	FirstName   *string    `db:"first_name" json:"first_name"`
	Email       *string    `db:"email" json:"email"`
	Password    *string    `db:"password" json:"password"`
	CreatedDate *time.Time `db:"created_date" json:"created_date"`
}

func (db *DB) CreateAppUser(u *AppUser) (*AppUser, error) {
	hashByte, err := auth.Hash(*u.Password)
	if err != nil {
		return nil, err
	}
	hashPass := string(hashByte)
	u.Password = &hashPass
	now := time.Now()
	u.CreatedDate = &now
	rows, err := db.Query(
		`INSERT INTO app_user(
			first_name, email, password, created_date
		) VALUES (
			$1, $2, $3, $4
		) RETURNING id`, u.FirstName, u.Email, u.Password, u.CreatedDate)
	if err != nil {
		return nil, fmt.Errorf("failed to insert new app_user into database: %v", err)
	}
	if rows.Next() {
		rows.Scan(&u.ID)
	}
	return u, nil
}
