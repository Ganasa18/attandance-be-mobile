package modelhelper

import (
	"ganasa18/attandance-be-mobile/internal/model/domain"
	"ganasa18/attandance-be-mobile/internal/model/web"
)

func ToUserMenuAccessResponse(userMenuAccess domain.UserMenuAccessModel) web.UserMenuAccessResponseRequest {
	return web.UserMenuAccessResponseRequest{
		Id:           userMenuAccess.Id,
		MenuId:       userMenuAccess.MenuId,
		UserId:       userMenuAccess.UserId,
		RoleAccessId: userMenuAccess.RoleAccessId,
		CreatedAt:    userMenuAccess.CreatedAt,
		UpdatedAt:    userMenuAccess.UpdatedAt,
	}
}

func ToUserMenuAccessResponses(userMenuAccesss []domain.UserMenuAccessModel, rowCount int) ([]web.UserMenuAccessResponseRequest, int) {
	var userMenuAccessResponse []web.UserMenuAccessResponseRequest
	for _, userMenuAccess := range userMenuAccesss {
		userMenuAccessResponse = append(userMenuAccessResponse, ToUserMenuAccessResponse(userMenuAccess))
	}
	return userMenuAccessResponse, rowCount
}

func ToUserMenusJoinResponse(userMenuAccess domain.UserMenuAccessJoin) web.UserMenuAccessByUserIdResponseRequest {
	return web.UserMenuAccessByUserIdResponseRequest{
		UmaID:        userMenuAccess.UmaID,
		MenuID:       userMenuAccess.MenuID,
		UserID:       userMenuAccess.UserID,
		RoleAccessID: userMenuAccess.RoleAccessID,
		RoleID:       userMenuAccess.RoleID,
		RoleName:     userMenuAccess.RoleName,
		MenuName:     userMenuAccess.MenuName,
		TitleMenu:    userMenuAccess.TitleMenu,
		Path:         userMenuAccess.Path,
		IsSubmenu:    userMenuAccess.IsSubmenu,
		ParentName:   userMenuAccess.ParentName,
	}
}

func ToUserMenusJoinResponses(userMenus []domain.UserMenuAccessJoin) []web.UserMenuAccessByUserIdResponseRequest {
	var userMenuRespone []web.UserMenuAccessByUserIdResponseRequest
	for _, userMenu := range userMenus {
		userMenuRespone = append(userMenuRespone, ToUserMenusJoinResponse(userMenu))
	}
	return userMenuRespone
}
