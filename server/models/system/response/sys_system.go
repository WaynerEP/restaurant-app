package response

import "github.com/WaynerEP/restaurant-app/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
