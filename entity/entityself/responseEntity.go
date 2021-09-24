package entityself

type Response struct {
	Code   string      `json:"response_code"`
	Refnum string      `json:"response_refnum"`
	ID     string      `json:"response_id"`
	Desc   string      `json:"response_description"`
	Data   interface{} `json:"response_data"`
}

func NewResponse(id string) *Response {
	return &Response{
		ID:   id,
		Code: "XX",
		Desc: "General Error",
		Data: new(struct{}),
	}
}
