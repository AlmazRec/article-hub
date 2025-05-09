package errors

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func NewErrorResponse(code int, message string, details ...string) *ErrorResponse {
	if len(details) > 0 {
		return &ErrorResponse{
			Code:    code,
			Message: message,
			Details: details[0],
		}
	}

	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}
