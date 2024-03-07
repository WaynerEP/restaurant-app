package system

import "github.com/WaynerEP/restaurant-app/server/config"

// System represents the configuration file structure.
type System struct {
	Config config.Server `json:"config"`
}
