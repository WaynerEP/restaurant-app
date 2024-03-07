package request

const (
	MaxPageSize = 100
	MinPageSize = 10
)

// PageInfo is the structure for common paging input parameters
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // Page number
	PageSize int    `json:"pageSize" form:"pageSize"` // Page size
	Keyword  string `json:"keyword" form:"keyword"`   // Keyword
}

// Get get pagination struct.
func (p *PageInfo) Get() *PageInfo {
	return p
}

// GetPage get page value from pagination struct.
func (p *PageInfo) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

// GetLimit get limit value from pagination struct.
func (p *PageInfo) GetLimit() int {
	if p.PageSize > 100 {
		p.PageSize = MaxPageSize
	}
	if p.PageSize <= 0 {
		p.PageSize = MinPageSize
	}
	return p.PageSize
}

// GetById is the structure for finding by ID
type GetById struct {
	ID int `json:"id" form:"id"` // Primary key ID
}

// IdsReq is the structure for batch IDs request
type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId is the structure for getting role by ID
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId" validate:"required"` // Role ID
}

// Empty is an empty structure
type Empty struct{}
