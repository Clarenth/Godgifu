package services

import "github.com/labstack/echo/v4"

type VideoService interface {
	CreateVideo(ctx echo.Context) (err error)
}
