package controller

import (
	"example/helpers"
	"example/model"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func (u User) Register(ctx echo.Context) error {
	var rq UserRequestAuth
	db := u.DB

	if err := ctx.Bind(&rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newUser := model.User{Username: rq.Username, Password: rq.Password}
	if err := db.Create(&newUser).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success register user",
	})
}

func (u User) Login(ctx echo.Context) error {
	var rq UserRequestAuth
	db := u.DB

	if err := ctx.Bind(&rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var loggedinUser model.User
	result := db.Where("username = ?", rq.Username).First(&loggedinUser)
	if result.Error != nil || result.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Username/Password Invalid")
	}

	validPassword := helpers.ComparePassword(rq.Password, loggedinUser.Password)
	if !validPassword {
		return echo.NewHTTPError(http.StatusUnauthorized, "Username/Password Invalid")
	}

	token := helpers.GenerateToken(jwt.MapClaims{
		"id": loggedinUser.ID,
	})

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    token,
	})
}
