package routeshandler

import (
	"database/sql"
	controller "ganasa18/attandance-be-mobile/internal/controller/crud"
	repository "ganasa18/attandance-be-mobile/internal/repository/crud"
	service "ganasa18/attandance-be-mobile/internal/service/crud"

	"github.com/go-playground/validator"
)

func SetupMenuHandlerRoute(db *sql.DB, validate *validator.Validate) controller.MenuController {

	userHandlerRoute := controller.NewMenuController(
		service.NewMenuService(
			repository.NewMenuRepository(),
			db,
			validate,
		),
	)

	return userHandlerRoute
}
