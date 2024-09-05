package challenges

type ChallengeGetResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Difficulty  int    `json:"difficulty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ChallengeUpdateResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Difficulty  int    `json:"difficulty"`
	UpdatedAt   string `json:"updated_at"`
}
