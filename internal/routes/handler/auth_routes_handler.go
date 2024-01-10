package routeshandler

import (
	"database/sql"
	controller "ganasa18/attandance-be-mobile/internal/controller/auth"
	repository "ganasa18/attandance-be-mobile/internal/repository/auth"
	service "ganasa18/attandance-be-mobile/internal/service/auth"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
)

func SetupAuthHandlerRoute(db *sql.DB, validate *validator.Validate) controller.AuthController {

	authHandlerRoute := controller.NewAuthController(
		service.NewAuthService(
			repository.NewAuthRepository(),
			db,
			validate,
		),
	)

	return authHandlerRoute
}
