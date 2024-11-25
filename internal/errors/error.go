package errors

type BadRequestException struct {
	Message    string
	StatusCode int
}

func (e *BadRequestException) Error() string {
	return e.Message
}

func NewBadRequestException(message string) *BadRequestException {
	if message == "" {
		message = "Bad Request"
	}
	return &BadRequestException{
		Message:    message,
		StatusCode: 400,
	}
}

type UnauthorizedException struct {
	Message    string
	StatusCode int
}

func (e *UnauthorizedException) Error() string {
	return e.Message
}

func NewUnauthorizedException() *UnauthorizedException {
	return &UnauthorizedException{
		Message:    "Unauthorized",
		StatusCode: 401,
	}
}
