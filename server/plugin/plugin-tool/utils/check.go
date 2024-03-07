package utils

import (
	"fmt"
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/system"
	"strconv"
)

// RegisterApis registers API routes in the database.
func RegisterApis(apis ...system.SysApi) {
	var count int64
	var apiPaths []string
	for i := range apis {
		apiPaths = append(apiPaths, apis[i].Path)
	}
	global.GVA_DB.Find(&[]system.SysApi{}, "path in (?)", apiPaths).Count(&count)
	if count > 0 {
		fmt.Println("The plugin is already installed or there are routes with the same name.")
		return
	}
	err := global.GVA_DB.Create(&apis).Error
	if err != nil {
		fmt.Println(err)
	}
}

// RegisterMenus registers menus in the database.
func RegisterMenus(menus ...system.SysBaseMenu) {
	var count int64
	var menuNames []string
	parentMenu := menus[0]
	otherMenus := menus[1:]
	for i := range menus {
		menuNames = append(menuNames, menus[i].Name)
	}
	global.GVA_DB.Find(&[]system.SysBaseMenu{}, "name in (?)", menuNames).Count(&count)
	if count > 0 {
		fmt.Println("The plugin is already installed or there are menus with the same name.")
		return
	}
	parentMenu.ParentId = "0"
	err := global.GVA_DB.Create(&parentMenu).Error
	if err != nil {
		fmt.Println(err)
	}
	for i := range otherMenus {
		pid := strconv.Itoa(int(parentMenu.ID))
		otherMenus[i].ParentId = pid
	}
	err = global.GVA_DB.Create(&otherMenus).Error
	if err != nil {
		fmt.Println(err)
	}
}
