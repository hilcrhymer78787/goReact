package request

type (
	Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
