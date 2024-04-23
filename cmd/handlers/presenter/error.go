package presenter

type ErrorType struct {
	Message string `json:"errorMessage"`
}

func Error(err error) ErrorType {
	return ErrorType{
		Message: err.Error(),
	}
}
