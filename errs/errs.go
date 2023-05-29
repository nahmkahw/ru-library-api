package errs

import (
	"net/http"

	"github.com/labstack/echo"
)

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnExpectedError() error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: "Un Expected Error.",
	}
}

func NewBadRequestError(message string) error {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func HandleError(c echo.Context, err error) {
	switch e := err.(type) {
	case AppError:
		c.JSON(http.StatusBadRequest, e.Message)
	case error:
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}
