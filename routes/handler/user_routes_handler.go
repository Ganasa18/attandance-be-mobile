package routeshandler

import (
	"database/sql"
	"ganasa18/attandance-be-mobile/controller"
	"ganasa18/attandance-be-mobile/repository"
	"ganasa18/attandance-be-mobile/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
)

func SetupUserHandlerRoute(db *sql.DB, validate *validator.Validate) controller.UserController {

	userHandlerRoute := controller.NewUserController(
		service.NewUserService(
			repository.NewUserRepository(),
			db,
			validate,
		),
	)

	return userHandlerRoute
}
