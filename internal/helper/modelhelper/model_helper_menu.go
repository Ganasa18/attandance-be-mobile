package modelhelper

import (
	"ganasa18/attandance-be-mobile/internal/model/domain"
	"ganasa18/attandance-be-mobile/internal/model/web"
)

func ToMenuResponse(menu domain.MenuModel) web.MenuMasterResponseRequest {
	return web.MenuMasterResponseRequest{
		Id:         menu.Id,
		MenuName:   menu.MenuName,
		TitleMenu:  menu.TitleMenu,
		Path:       menu.Path,
		IsSubmenu:  menu.IsSubmenu,
		ParentName: menu.ParentName,
		CreatedAt:  menu.CreatedAt,
		UpdatedAt:  menu.UpdatedAt,
	}
}

func ToMenuResponses(menus []domain.MenuModel, rowCount int) ([]web.MenuMasterResponseRequest, int) {
	var menuResponse []web.MenuMasterResponseRequest
	for _, menu := range menus {
		menuResponse = append(menuResponse, ToMenuResponse(menu))
	}
	return menuResponse, rowCount
}
