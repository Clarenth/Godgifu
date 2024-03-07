package handlers

import (
	"godgifu/modules/video/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	VideoService services.VideoService
}

func NewVideoHandlers(service services.VideoService) VideoHandlers {
	return &handler{
		VideoService: service,
	}
}

func (handler *handler) NewVideo(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!\n")
}
