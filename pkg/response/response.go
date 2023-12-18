package response

type Response struct {
	Code    Code        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(code Code, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func OK(data interface{}) *Response {
	return NewResponse(CodeOK, "OK", data)
}

func Created(data interface{}) *Response {
	return NewResponse(CodeCreated, "Created", data)
}

func Fail(code Code, message string) *Response {
	return NewResponse(code, message, nil)
}
