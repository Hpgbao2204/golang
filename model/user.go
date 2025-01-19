package model

import "time"

type User struct {
	UserID    string    `json:"-" db:"user_id, omitempty" json:"userID,omitempty"`
	FullName  string    `json:"fullName,omitempty" db:"full_name, omitempty"`
	Email     string    `db:"email, omitempty" json:"email,omitempty"`
	Password  string    `db:"password, omitempty" json:"password,omitempty"`
	Role      string    `db:"role, omitempty" json:"role,omitempty"`
	CreatedAt time.Time `db:"created_at, omitempty" json:"-"`
	UpdatedAt time.Time `db:"updated_at, omitempty" json:"-"`
	Token     string    `json:"-"`
}
