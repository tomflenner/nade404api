package router

import (
	"github.com/b4cktr4ck5r3/nade404api/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/players", handler.GetPlayers)
	api.Get("/players/:steam_id", handler.GetPlayerBySteamID)
	api.Get("/top10kd", handler.GetTop10PlayersByKd)
	api.Get("/top10hs", handler.GetTop10PlayersByHs)
	api.Get("/get5config/:config_id", handler.Get5Config)
	api.Post("/get5config", handler.CreateGet5Config)
	api.Post("/get5config/:match_id/log", handler.HandleGet5ConfigLogs)
	api.Get("/ptero/server", handler.GetPteroServerByIpAndPort)
}
