package router

import (
	"github.com/WaynerEP/restaurant-app/server/router/contact"
	"github.com/WaynerEP/restaurant-app/server/router/inventory"
	"github.com/WaynerEP/restaurant-app/server/router/menu"
	"github.com/WaynerEP/restaurant-app/server/router/system"
)

type Group struct {
	System    system.RouterGroup
	Contact   contact.RouterGroup
	Inventory inventory.RouterGroup
	Menu      menu.RouterGroup
}

var AppRouterGroup = new(Group)
