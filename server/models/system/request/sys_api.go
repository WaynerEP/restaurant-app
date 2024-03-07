package request

import (
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/system"
)

// SearchApiParams represents the structure for paginated, conditional queries, and sorting for an API.
type SearchApiParams struct {
	system.SysApi           // Embedded field, including fields from the `SysApi` structure
	request.PageInfo        // Embedded field, including fields from the `PageInfo` structure
	OrderKey         string `json:"orderKey"` // OrderKey is the field used for sorting
	Desc             bool   `json:"desc"`     // Desc is the sorting order: ascending (false by default) or descending (true)
}
