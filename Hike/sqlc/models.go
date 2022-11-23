// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package hikedb

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID        int32
	PostId    int32
	Name      string
	Body      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type Route struct {
	ID          int32
	UserId      int32
	Title       string
	Description string
	Location    string
	Destination float64
	Height      float64
	Level       string
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
}

type User struct {
	ID        int32
	FirstName string
	LastName  string
	Age       int32
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
