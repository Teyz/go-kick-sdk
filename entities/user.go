package entities

type User struct {
	Email          string `json:"email"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture"`
	UserID         int    `json:"user_id"`
}

type GetUsersResponse struct {
	Data    []*User `json:"data"`
	Message string  `json:"message"`
}
