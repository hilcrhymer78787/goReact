package response

type (
	Res struct {
		Id       int    `json:"id"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Token    string `json:"token"`
		User_img string `json:"user_img"`
	}
)
