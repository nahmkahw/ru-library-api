package handler

import (
	"net/http"
	"ru-library-api/entity"

	"github.com/labstack/echo"
)

func (h *sierraHandler) Patron(c echo.Context) error {
	id := c.Param("id")

	patrons, err := h.sierraService.PatronById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, entity.Response{
			Code:         "patron-error",
			Data:         err.Error(),
			HttpCode:     http.StatusNotFound,
			ErrorMessage: "Patron By Id Failed",
		})
	}

	return c.JSON(http.StatusOK, entity.Response{
		Code:         "patron",
		Data:         patrons,
		HttpCode:     200,
		ErrorMessage: "",
	})

}
