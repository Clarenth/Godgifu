package handlers

import "github.com/labstack/echo/v4"

type VideoHandlers interface {
	NewVideo(ctx echo.Context) (err error)
}
