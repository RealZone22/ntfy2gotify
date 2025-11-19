package api

import (
	"ntfy2gotify/api"
	"ntfy2gotify/api/routes"
	"ntfy2gotify/pkg/utils"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartRestAPI() {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(api.LoggerMiddleware())

	e.POST("/:topic", routes.HandleNtfyRequests)
	e.POST("/", routes.HandleNtfyRequests)

	utils.Logger.Info().Str("host", utils.Config.Api.Host).Int("port", utils.Config.Api.Port).Msg("Starting REST API")

	utils.Logger.Fatal().Err(e.Start(utils.Config.Api.Host + ":" + strconv.Itoa(utils.Config.Api.Port))).Msg("Failed to start API")
}
