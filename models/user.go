package models

import "time"

type User struct {
	ID        int64   `db:"id"`
	Name      string  `db:"name"`
	Username  string  `db:"username" gorm:"uniqueIndex"`
	Password  string  `db:"password"`
	Gender    string  `db:"gender"`
	Balance   float64 `db:"balance"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
