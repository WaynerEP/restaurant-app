package system

import (
	"errors"
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/system"
	"github.com/WaynerEP/restaurant-app/server/models/system/request"
	"github.com/WaynerEP/restaurant-app/server/models/system/response"
	"gorm.io/gorm"
)

// AuthorityBtnService provides operations related to authority buttons.
type AuthorityBtnService struct{}

// GetAuthorityBtn retrieves authority buttons based on the provided request.
func (a *AuthorityBtnService) GetAuthorityBtn(req request.SysAuthorityBtnReq) (res response.SysAuthorityBtnRes, err error) {
	var authorityBtn []system.SysAuthorityBtn
	err = global.GVA_DB.Find(&authorityBtn, "authority_id = ? and sys_menu_id = ?", req.AuthorityId, req.MenuID).Error
	if err != nil {
		return
	}
	var selected []uint
	for _, v := range authorityBtn {
		selected = append(selected, v.SysBaseMenuBtnID)
	}
	res.Selected = selected
	return res, err
}

// SetAuthorityBtn sets authority buttons based on the provided request.
func (a *AuthorityBtnService) SetAuthorityBtn(req request.SysAuthorityBtnReq) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var authorityBtn []system.SysAuthorityBtn
		err = tx.Delete(&[]system.SysAuthorityBtn{}, "authority_id = ? and sys_menu_id = ?", req.AuthorityId, req.MenuID).Error
		if err != nil {
			return err
		}
		for _, v := range req.Selected {
			authorityBtn = append(authorityBtn, system.SysAuthorityBtn{
				AuthorityId:      req.AuthorityId,
				SysMenuID:        req.MenuID,
				SysBaseMenuBtnID: v,
			})
		}
		if len(authorityBtn) > 0 {
			err = tx.Create(&authorityBtn).Error
		}
		if err != nil {
			return err
		}
		return err
	})
}

// CanRemoveAuthorityBtn checks if an authority button can be removed based on the provided ID.
func (a *AuthorityBtnService) CanRemoveAuthorityBtn(ID string) (err error) {
	fErr := global.GVA_DB.First(&system.SysAuthorityBtn{}, "sys_base_menu_btn_id = ?", ID).Error
	if errors.Is(fErr, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.New("this button is in use and cannot be deleted")
}
