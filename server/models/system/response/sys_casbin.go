package response

import "github.com/WaynerEP/restaurant-app/server/models/system/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
