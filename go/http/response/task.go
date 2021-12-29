package response

type (
	Res struct {
		Tasks []Task `json:"tasks"`
		Hoge string `json:"hoge"`
	}
	Task struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}
)
