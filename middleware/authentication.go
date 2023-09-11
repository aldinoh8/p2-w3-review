package middleware

import (
	"example/helpers"
	"example/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Auth struct {
	DB *gorm.DB
}

func (a Auth) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// token
		// cek ada apa ga
		token := ctx.Request().Header.Get("token")
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token credentials")
		}

		// decode/verify
		claims, err := helpers.VeriyToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token credentials")
		}

		// cek ke db user
		var loggedinUser model.User
		result := a.DB.First(&loggedinUser, claims["id"])
		if result.Error != nil || result.RowsAffected == 0 {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token credentials")
		}

		// set user ke context (optional?)

		//next
		return next(ctx)
	}
}
