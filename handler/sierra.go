package handler

import "ru-library-api/service"

type sierraHandler struct {
	sierraService service.SierraServiceInterface
}

func NewSierraHandler(sierraService service.SierraServiceInterface) sierraHandler {
	return sierraHandler{sierraService: sierraService}
}
