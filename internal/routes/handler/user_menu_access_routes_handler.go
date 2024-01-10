package routeshandler

import (
	"database/sql"
	controller "ganasa18/attandance-be-mobile/internal/controller/crud"
	repository "ganasa18/attandance-be-mobile/internal/repository/crud"
	service "ganasa18/attandance-be-mobile/internal/service/crud"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
)

func SetupUserMenuAccessHandlerRoute(db *sql.DB, validate *validator.Validate) controller.UserMenuAccessController {

	userMenuAccessHandlerRoute := controller.NewUserMenuAccessController(
		service.NewUserMenuAccessService(
			repository.NewUserMenuAccessRepository(),
			db,
			validate,
		),
	)

	return userMenuAccessHandlerRoute
}
