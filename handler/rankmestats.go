package handler

import (
	"fmt"

	"github.com/b4cktr4ck5r3/nade404api/database"
	"github.com/b4cktr4ck5r3/nade404api/model"
	"github.com/gofiber/fiber/v2"
)

func GetPlayers(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, steam, name, score, FIND_IN_SET( score, ( SELECT GROUP_CONCAT( score ORDER BY score DESC ) FROM rankme ) ) AS rank, mvp, kills, deaths, ROUND((kills/deaths),2) as ratio, headshots, ROUND((headshots/kills) * 100, 0) as headshots_percent, assists, assist_flash, no_scope, thru_smoke, blind, wallbang FROM rankme")
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	defer rows.Close()
	result := model.Players{}
	for rows.Next() {
		player := model.Player{}
		err := rows.Scan(&player.Id, &player.SteamID, &player.Name, &player.Score, &player.Rank, &player.Mvp, &player.Kills, &player.Deaths, &player.Ratio, &player.Headshots, &player.HeadshotsPercent, &player.Assists, &player.FlashAssists, &player.NoScope, &player.ThruSmoke, &player.Blind, &player.Wallbang)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result.Players = append(result.Players, player)
	}

	if len(result.Players) <= 0 {
		return c.Status(404).JSON(&fiber.Map{
			"succes":  false,
			"message": "No player found",
		})
	}

	response := &fiber.Map{
		"success": true,
		"players": result,
		"message": "All players returned successfully",
	}

	if err := c.JSON(response); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"succes":  false,
			"message": err,
		})
	}

	return c.Status(200).JSON(response)
}

func GetPlayerBySteamID(c *fiber.Ctx) error {
	rows, err := database.DB.Query(fmt.Sprintf("SELECT id, steam, name, score, FIND_IN_SET( score, ( SELECT GROUP_CONCAT( score ORDER BY score DESC ) FROM rankme ) ) AS rank, mvp, kills, deaths, ROUND((kills/deaths),2) as ratio, headshots, ROUND((headshots/kills) * 100, 0) as headshots_percent, assists, assist_flash, no_scope, thru_smoke, blind, wallbang FROM rankme WHERE steam = '%s'", c.Params("steam_id")))
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	defer rows.Close()
	if rows.Next() {
		player := model.Player{}
		err := rows.Scan(&player.Id, &player.SteamID, &player.Name, &player.Score, &player.Rank, &player.Mvp, &player.Kills, &player.Deaths, &player.Ratio, &player.Headshots, &player.HeadshotsPercent, &player.Assists, &player.FlashAssists, &player.NoScope, &player.ThruSmoke, &player.Blind, &player.Wallbang)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result := player
		response := &fiber.Map{
			"success": true,
			"players": result,
			"message": fmt.Sprintf("Player %s returned successfully", result.SteamID),
		}

		if err := c.JSON(response); err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"succes":  false,
				"message": err,
			})
		}
		return c.Status(200).JSON(response)

	}
	return c.Status(404).JSON(&fiber.Map{
		"success": false,
		"message": "Player not found",
	})
}

func GetTop10PlayersByKd(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, steam, name, score, FIND_IN_SET( score, ( SELECT GROUP_CONCAT( score ORDER BY score DESC ) FROM rankme ) ) AS rank, mvp, kills, deaths, ROUND((kills/deaths),2) as ratio, headshots, ROUND((headshots/kills) * 100, 0) as headshots_percent, assists, assist_flash, no_scope, thru_smoke, blind, wallbang FROM `rankme` WHERE kills > 750 ORDER BY ratio DESC LIMIT 10")
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	defer rows.Close()
	result := model.Players{}
	for rows.Next() {
		player := model.Player{}
		err := rows.Scan(&player.Id, &player.SteamID, &player.Name, &player.Score, &player.Rank, &player.Mvp, &player.Kills, &player.Deaths, &player.Ratio, &player.Headshots, &player.HeadshotsPercent, &player.Assists, &player.FlashAssists, &player.NoScope, &player.ThruSmoke, &player.Blind, &player.Wallbang)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result.Players = append(result.Players, player)
	}

	if len(result.Players) <= 0 {
		return c.Status(404).JSON(&fiber.Map{
			"succes":  false,
			"message": "No player found for top 10 by kd",
		})
	}

	response := &fiber.Map{
		"success": true,
		"players": result,
		"message": "All players returned successfully",
	}

	if err := c.JSON(response); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"succes":  false,
			"message": err,
		})
	}

	return c.Status(200).JSON(response)
}

func GetTop10PlayersByHs(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, steam, name, score, FIND_IN_SET( score, ( SELECT GROUP_CONCAT( score ORDER BY score DESC ) FROM rankme ) ) AS rank, mvp, kills, deaths, ROUND((kills/deaths),2) as ratio, headshots, ROUND((headshots/kills) * 100, 0) as headshots_percent, assists, assist_flash, no_scope, thru_smoke, blind, wallbang FROM `rankme` WHERE kills > 750 ORDER BY headshots_percent DESC LIMIT 10")
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	defer rows.Close()
	result := model.Players{}
	for rows.Next() {
		player := model.Player{}
		err := rows.Scan(&player.Id, &player.SteamID, &player.Name, &player.Score, &player.Rank, &player.Mvp, &player.Kills, &player.Deaths, &player.Ratio, &player.Headshots, &player.HeadshotsPercent, &player.Assists, &player.FlashAssists, &player.NoScope, &player.ThruSmoke, &player.Blind, &player.Wallbang)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result.Players = append(result.Players, player)
	}

	if len(result.Players) <= 0 {
		return c.Status(404).JSON(&fiber.Map{
			"succes":  false,
			"message": "No player found for top 10 by hs",
		})
	}

	response := &fiber.Map{
		"success": true,
		"players": result,
		"message": "All players returned successfully",
	}

	if err := c.JSON(response); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"succes":  false,
			"message": err,
		})
	}

	return c.Status(200).JSON(response)
}
