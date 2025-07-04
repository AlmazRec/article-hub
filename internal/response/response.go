package response

type SuccessResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Error   string      `json:"error,omitempty"`
}
