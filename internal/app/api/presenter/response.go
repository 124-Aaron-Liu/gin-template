package presenter

type Response struct {
	Code    string      `json:"ret_code"`
	Message string      `json:"ret_msg"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *Response) WithData(data interface{}) *Response {
	r.Data = data
	return r
}

func NewReponse(code string, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}

func OK() *Response {
	return NewReponse("0", "ok")
}
