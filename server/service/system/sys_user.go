package system

import (
	"errors"
	"fmt"
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/system"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"time"

	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

// UserService provides user-related operations.
type UserService struct{}

// Login handles user login.
func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if global.GVA_DB == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	var user system.SysUser
	err = global.GVA_DB.Where("email = ?", u.Email).
		Preload("Authorities").
		Preload("SysAuthority").
		First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("Incorrect password")
		}
		MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}

// Register registers a new user.
func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.GVA_DB.Where("email = ?", u.Email).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("Email already registered")
	}
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())
	err = global.GVA_DB.Create(&u).Error
	return u, err
}

// ChangePassword changes a user's password.
func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (userInter *system.SysUser, err error) {
	var user system.SysUser
	if err = global.GVA_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("Incorrect original password")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.GVA_DB.Save(&user).Error
	return &user, err
}

// GetUserInfoList retrieves user information with pagination.
func (userService *UserService) GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.GetLimit()
	page := info.GetPage()
	offset := limit * (page - 1)

	db := global.GVA_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).
		Offset(offset).
		Preload("Authorities").
		Preload("SysAuthority").
		Find(&userList).Error
	return userList, total, err
}

// SetUserAuthority sets the authority of a user.
func (userService *UserService) SetUserAuthority(id uint, authorityId uint) (err error) {
	assignErr := global.GVA_DB.Where("sys_user_id = ? AND sys_authority_id = ?", id, authorityId).First(&system.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("User does not have this role")
	}
	err = global.GVA_DB.Where("id = ?", id).First(&system.SysUser{}).Update("sys_authority_id", authorityId).Error
	return err
}

// SetUserAuthorities sets the authorities of a user.
func (userService *UserService) SetUserAuthorities(id uint, authorityIds []uint) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []system.SysUserAuthority
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, system.SysUserAuthority{
				SysUserId: id, SysAuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("id = ?", id).First(&system.SysUser{}).Update("sys_authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// Return nil to commit the transaction.
		return nil
	})
}

// DeleteUser deletes a user.
func (userService *UserService) DeleteUser(id int) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&system.SysUser{}, id).Error; err != nil {
			return err
		}
		if err := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
}

// SetUserInfo sets user information.
func (userService *UserService) SetUserInfo(req system.SysUser) error {
	return global.GVA_DB.Model(&system.SysUser{}).
		Select("updated_at", "nick_name", "header_img", "phone", "email", "sideMode", "enable").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  req.NickName,
			"header_img": req.HeaderImg,
			"phone":      req.Phone,
			"email":      req.Email,
			"side_mode":  req.SideMode,
			"enable":     req.Enable,
		}).Error
}

// SetSelfInfo sets user information for self.
func (userService *UserService) SetSelfInfo(req system.SysUser) error {
	return global.GVA_DB.Model(&system.SysUser{}).
		Where("id=?", req.ID).
		Updates(req).Error
}

// GetUserInfo retrieves user information by UUID.
func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user system.SysUser, err error) {
	var reqUser system.SysUser
	err = global.GVA_DB.Preload("Authorities").Preload("SysAuthority").First(&reqUser, "uuid = ?", uuid).Error
	if err != nil {
		return reqUser, err
	}
	MenuServiceApp.UserAuthorityDefaultRouter(&reqUser)
	return reqUser, err
}

// FindUserById retrieves user information by ID.
func (userService *UserService) FindUserById(id int) (user *system.SysUser, err error) {
	var u system.SysUser
	err = global.GVA_DB.Where("id = ?", id).First(&u).Error
	return &u, err
}

// FindUserByUuid retrieves user information by UUID string.
func (userService *UserService) FindUserByUuid(uuid string) (user *system.SysUser, err error) {
	var u system.SysUser
	if err = global.GVA_DB.Where("uuid = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("User does not exist")
	}
	return &u, nil
}

// ResetPassword resets a user's password.
func (userService *UserService) ResetPassword(ID uint) (err error) {
	err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", ID).Update("password", utils.BcryptHash("12345678")).Error
	return err
}
