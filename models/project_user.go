package models

import "time"

type ProjectUser struct {
	ID      int
	Project   *Project
	ProjectID int
	User    *User
	UserID  int
	Start	time.Time
	End		time.Time
}
