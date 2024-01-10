package routeshandler

import (
	"database/sql"
	controller "ganasa18/attandance-be-mobile/internal/controller/user"
	repository "ganasa18/attandance-be-mobile/internal/repository/user"
	service "ganasa18/attandance-be-mobile/internal/service/user"

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
