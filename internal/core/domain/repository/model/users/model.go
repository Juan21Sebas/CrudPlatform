package users

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImagePath string    `json:"image_path,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUser struct {
	Id string `json:"id"`
}

type UpdateUser struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	ImagePath string `json:"image_path,omitempty"`
}

type DeleteUser struct {
	Id string `json:"id"`
}
