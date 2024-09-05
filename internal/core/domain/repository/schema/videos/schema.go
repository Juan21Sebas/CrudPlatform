package videos

type VideosGetResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type VideosUpdateResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updated_at"`
}
