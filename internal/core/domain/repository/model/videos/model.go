package videos

import "time"

type Videos struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetVideo struct {
	ID string `json:"id"`
}

type UpdateVideo struct {
	ID          string `json:"id"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type DeleteVideo struct {
	ID string `json:"id"`
}
