package dtos

type ValidationError struct {
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}

type UnexpectedError struct {
	Message string `json:"message"`
}

type Error struct {
	Message          string            `json:"message"`
	StatusCode       int               `json:"status_code"`
	ValidationErrors []ValidationError `json:"validation_errors"`
	UnexpectedError  *UnexpectedError  `json:"unexpected_error"`
}

func (e *Error) SetStatusCode(statusCode int) {
	e.StatusCode = statusCode
}

func (e *Error) SetMessage(message string) {
	e.Message = message
}

func (e *Error) SetValidationErrors(validationErrors []ValidationError) {
	e.ValidationErrors = validationErrors
}

func (e *Error) SetUnexpectedError(unexpectedError *UnexpectedError) {
	e.UnexpectedError = unexpectedError
}

func NewValidationError(field string, messages ...string) ValidationError {
	return ValidationError{
		Field:    field,
		Messages: messages,
	}
}

func NewUnexpectedError(message string) UnexpectedError {
	return UnexpectedError{
		Message: message,
	}
}

func NewError(message string, statusCode int) Error {
	return Error{
		Message:    message,
		StatusCode: statusCode,
	}
}
