package initialize

import (
	_ "github.com/WaynerEP/restaurant-app/server/source/contact"
	_ "github.com/WaynerEP/restaurant-app/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
