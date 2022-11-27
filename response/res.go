package response

type Res struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}
