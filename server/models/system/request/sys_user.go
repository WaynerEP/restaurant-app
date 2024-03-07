package request

import "github.com/WaynerEP/restaurant-app/server/models/system"

// Register User registration structure
type Register struct {
	ID           uint   `json:"id"` //!
	Username     string `json:"userName" example:"Username" validate:"required"`
	Password     string `json:"passWord" example:"Password" validate:"required"`
	NickName     string `json:"nickName" example:"Nickname"`
	HeaderImg    string `json:"headerImg" example:"Avatar URL"`
	AuthorityId  uint   `json:"authorityId" swaggertype:"string" example:"int Role ID" validate:"required"`
	Enable       int    `json:"enable" swaggertype:"string" example:"int Enabled"`
	EmployeeId   uint   `json:"employeeId" validate:"required"`
	CompanyId    uint   `json:"companyId"`
	AuthorityIds []uint `json:"authorityIds" swaggertype:"string" example:"[]uint Role IDs"`
	Phone        string `json:"phone" example:"Phone number"`
	Email        string `json:"email" example:"Email" validate:"required,unique_db=email:sys_users"`
}

// Login User login structure
type Login struct {
	Email    string `json:"email" validate:"required"`    // Username
	Password string `json:"password" validate:"required"` // Password
	//Captcha   string `json:"captcha"`   // Captcha
	//	CaptchaId string `json:"captchaId"` // Captcha ID
}

// ChangePasswordReq Modify password structure
type ChangePasswordReq struct {
	ID              uint   `json:"-"`                                     // Extract user ID from JWT to avoid unauthorized access
	Password        string `json:"password" validate:"required,min=8"`    // Password
	NewPassword     string `json:"newPassword" validate:"required,min=8"` // New password
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8,eqfield=NewPassword"`
}

// SetUserAuth Modify user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId" validate:"required"` // Role ID
}

// SetUserAuthorities Modify user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // Role IDs
}

type ChangeUserInfo struct {
	ID           uint                  `gorm:"primarykey"`                                                                                  // Primary key ID
	NickName     string                `json:"nickName" gorm:"default:System User;comment:User nickname"`                                   // User nickname
	Phone        string                `json:"phone"  gorm:"comment:User phone number"`                                                     // User phone number
	AuthorityIds []uint                `json:"authorityIds" gorm:"-"`                                                                       // Role IDs
	Email        string                `json:"email"  gorm:"comment:User email" validate:"required"`                                        // User email
	HeaderImg    string                `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:User avatar"` // User avatar
	SideMode     string                `json:"sideMode"  gorm:"comment:User side theme"`                                                    // User side theme
	Enable       int                   `json:"enable" gorm:"comment:Freeze user"`                                                           // Freeze user
	Authorities  []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}
