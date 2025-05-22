package response

type (
	Res struct {
		Tasks []Task `json:"tasks"`
		Date  string `json:"date"`
	}
	Task struct {
		Id   int64  `json:"id"`
		UserId string `json:"user_id"`
		Name string `json:"name"`
		SortKey *int64 `json:"sort_key"`
		Work Work `json:"work"`
	}
	Work struct {
		Id   int64  `json:"id"`
		State int64 `json:"state"`
	}
)
