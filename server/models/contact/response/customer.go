package response

import "github.com/WaynerEP/restaurant-app/server/models/contact"

type CustomerResponse struct {
	Customer contact.Customer `json:"contact"`
}
