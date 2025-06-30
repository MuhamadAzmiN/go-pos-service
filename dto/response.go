package dto

type Response[T any] struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func CreateResponseData(message string) Response[string] {
	return Response[string]{
		Code:    "99",
		Message: message,
		Data:    "",
	}
}

func CreateResponseErrorData(message string) Response[string] {
	return Response[string]{
		Code:    "99",
		Message: message,
		Data:    "",
	}
}





func CreateResponseSuccessData(message string, data any) Response[any] {
	return Response[any]{
		Code:    "00",
		Message: message,
		Data:    data,
	}
}

func CreateResponseError(message string) Response[string] {
	return Response[string]{
		Code:    "99",
		Message: message,
		Data:    "",
	}
}
