package handler

import (
	"net/http"
	"ru-library-api/entity"

	"github.com/labstack/echo"
)

func (h *sierraHandler) Fine(c echo.Context) error {
	id := c.Param("id")

	fines, err := h.sierraService.FineById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, entity.Response{
			Code:         "fine-error",
			Data:         err.Error(),
			HttpCode:     http.StatusNotFound,
			ErrorMessage: "Fine By Id Failed",
		})
	}
	return c.JSON(http.StatusOK, entity.Response{
		Code:         "fine",
		Data:         fines,
		HttpCode:     200,
		ErrorMessage: "",
	})

}
