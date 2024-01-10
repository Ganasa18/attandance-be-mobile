package modelhelper

import (
	"ganasa18/attandance-be-mobile/internal/model/domain"
	"ganasa18/attandance-be-mobile/internal/model/web"
)

func ToRoleResponse(role domain.RoleModel) web.RoleMasterResponseRequest {
	return web.RoleMasterResponseRequest{
		Id:        role.Id,
		Rolename:  role.Rolename,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func ToRoleResponses(roles []domain.RoleModel, rowCount int) ([]web.RoleMasterResponseRequest, int) {
	var roleResponse []web.RoleMasterResponseRequest
	for _, role := range roles {
		roleResponse = append(roleResponse, ToRoleResponse(role))
	}
	return roleResponse, rowCount
}

func ToRoleAccessResponse(roleAccess domain.RoleAccessModel) web.RoleAccessMasterResponseRequest {
	return web.RoleAccessMasterResponseRequest{
		Id:        roleAccess.Id,
		RoleId:    roleAccess.RoleId,
		Create:    roleAccess.Create,
		Read:      roleAccess.Read,
		Update:    roleAccess.Update,
		Delete:    roleAccess.Delete,
		CreatedAt: roleAccess.CreatedAt,
		UpdatedAt: roleAccess.UpdatedAt,
	}
}

func ToRoleAccessResponses(rolesAccess []domain.RoleAccessModel, rowCount int) ([]web.RoleAccessMasterResponseRequest, int) {
	var roleAccessResponse []web.RoleAccessMasterResponseRequest
	for _, roleAcc := range rolesAccess {
		roleAccessResponse = append(roleAccessResponse, ToRoleAccessResponse(roleAcc))
	}
	return roleAccessResponse, rowCount
}
