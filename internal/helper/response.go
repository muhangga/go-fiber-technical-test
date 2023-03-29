package helper

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ValidationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type EmptyObject struct{}

func ValidResponse(status, message string) ValidationResponse {
	return ValidationResponse{
		Status:  status,
		Message: message,
	}
}

func BuildResponse(message string, data interface{}) Response {
	return Response{
		Status:  "Success",
		Message: message,
		Data:    data,
	}
}

func BuildErrorResponse(status, message string) Response {
	return Response{
		Status:  status,
		Message: message,
	}
}
