package users

type UsersGetResponse struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	ImagePath string `json:"image_path,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UsersUpdateResponse struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	ImagePath string `json:"image_path,omitempty"`
	UpdatedAt string `json:"updated_at"`
}
