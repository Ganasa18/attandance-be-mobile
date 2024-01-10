package routeshandler

import (
	"database/sql"
	controller "ganasa18/attandance-be-mobile/internal/controller/crud"
	repository "ganasa18/attandance-be-mobile/internal/repository/crud"
	service "ganasa18/attandance-be-mobile/internal/service/crud"

	"github.com/go-playground/validator"
)

func SetupRoleHandlerRoute(db *sql.DB, validate *validator.Validate) controller.RoleController {

	userHandlerRoute := controller.NewRoleController(
		service.NewRoleService(
			repository.NewRoleRepository(),
			db,
			validate,
		),
	)

	return userHandlerRoute
}

func SetupRoleAccessHandlerRoute(db *sql.DB, validate *validator.Validate) controller.RoleAccessController {
	userHandlerRoute := controller.NewRoleAccessController(
		service.NewRoleAccessService(
			repository.NewRoleAccessRepository(),
			db,
			validate,
		),
	)

	return userHandlerRoute
}
