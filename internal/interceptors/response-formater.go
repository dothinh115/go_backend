package interceptors

import "project/internal/errors"

func init() {
	InterceptorRegister(responseFormatter)
}

type Response struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"statusCode"`
}

func HandleError(err error) Response {
	switch e := err.(type) {
	case *errors.BadRequestException:
		{
			return Response{
				Message:    e.Message,
				StatusCode: e.StatusCode,
			}
		}
	case *errors.UnauthorizedException:
		{
			{
				return Response{
					Message:    e.Message,
					StatusCode: e.StatusCode,
				}
			}
		}
	default:
		{
			return Response{
				Message:    e.Error(),
				StatusCode: 500,
			}
		}
	}
}

func responseFormatter(data interface{}) interface{} {
	if err, ok := data.(error); ok {
		data = HandleError(err)
	} else {
		data = Response{
			Message:    "Thành công",
			Data:       data,
			StatusCode: 200,
		}
	}

	return data
}
