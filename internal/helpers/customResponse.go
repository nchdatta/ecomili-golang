package helpers

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(success bool, message string, data interface{}) *Response {
	return &Response{
		Success: success,
		Message: message,
		Data:    data,
	}
}
