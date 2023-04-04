package models

import "time"

type User struct {
	Username   string    `json:"username"`
	First_Name string    `json:"first_name"`
	Last_Name  string    `json:"last_name"`
	Password   string    `json:"password"`
	Created_At time.Time `json:"created_at"`
}
