package request

// CasbinInfo Casbin info structure
type CasbinInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

// CasbinInReceive Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId uint         `json:"authorityId" validate:"required,gt=0"`
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/menu/getMenu", Method: "POST"},
		{Path: "/jwt/jsonInBlacklist", Method: "POST"},
		{Path: "/base/login", Method: "POST"},
		{Path: "/user/admin_register", Method: "POST"},
		{Path: "/user/changePassword", Method: "POST"},
		{Path: "/user/setUserAuthority", Method: "POST"},
		{Path: "/user/setUserInfo", Method: "PUT"},
		{Path: "/user/getUserInfo", Method: "GET"},
	}
}
