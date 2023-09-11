package controller

import (
	"example/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Player struct {
	DB *gorm.DB
}

func (p Player) Index(ctx echo.Context) error {
	var players []model.Player
	if err := p.DB.Find(&players).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, players)
}

func (p Player) Create(ctx echo.Context) error {
	var reqBody PlayerRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newPlayer := model.Player{
		Username: reqBody.Username,
		TeamName: reqBody.TeamName,
		Ranking:  reqBody.Ranking,
		Score:    reqBody.Score,
	}

	if err := p.DB.Create(&newPlayer).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message":   "success create player",
		"newPlayer": newPlayer,
	})
}

func (p Player) Update(ctx echo.Context) error {
	return nil
}
