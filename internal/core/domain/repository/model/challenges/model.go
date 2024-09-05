package challenges

import "time"

type Challenge struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Difficulty  int       `json:"difficulty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetChallenge struct {
	ID string `json:"id"`
}

type UpdateChallenge struct {
	ID          string `json:"id"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Difficulty  int    `json:"difficulty,omitempty"`
}

type DeleteChallenge struct {
	ID string `json:"id"`
}
