package dto

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func ResponseError(message string, err error) ErrorResponse {
	errString := ""
	if err != nil {
		errString = err.Error()
	}
	return ErrorResponse{
		Message: message,
		Error:   errString,
	}
}
