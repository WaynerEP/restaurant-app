package system

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	"github.com/WaynerEP/restaurant-app/server/models/system"
	"strconv"
	"time"

	systemReq "github.com/WaynerEP/restaurant-app/server/models/system/request"
	systemRes "github.com/WaynerEP/restaurant-app/server/models/system/response"
	"github.com/WaynerEP/restaurant-app/server/utils"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// Login .
func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	//key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(l)
	if err != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}

	// Check if captcha is enabled
	//openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // Whether to enable captcha
	//openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // Cache timeout
	//v, ok := global.BlackCache.Get(key)
	//if !ok {
	//	global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	//}

	//ar oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	//if !oc || (l.CaptchaId != "" && l.Captcha != "" && store.Verify(l.CaptchaId, l.Captcha, true)) {
	u := &system.SysUser{Email: l.Email, Password: l.Password}
	user, err := userService.Login(u)
	if err != nil {
		global.GVA_LOG.Error("Login failed! Username does not exist or password is incorrect!", zap.Error(err))
		// Increment captcha attempts
		//global.BlackCache.Increment(key, 1)
		response.FailWithMessage("Username does not exist or password is incorrect", c)
		return
	}
	if user.Enable != 1 {
		global.GVA_LOG.Error("Login failed! User is prohibited from logging in!")
		// Increment captcha attempts
		//global.BlackCache.Increment(key, 1)
		response.FailWithMessage("User is prohibited from logging in", c)
		return
	}
	b.TokenNext(c, *user)
	return
	//}
	// Increment captcha attempts
	//global.BlackCache.Increment(key, 1)
	//response.FailWithMessage("Incorrect captcha", c)
}

// TokenNext generates JWT after login
func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // Unique signature
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		CompanyId:   user.CompanyID,
		EmployeeId:  user.EmployeeId,
		AuthorityId: user.SysAuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("Failed to get token!", zap.Error(err))
		response.FailWithMessage("Failed to get token", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			UserData:  user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "Login successful", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("Failed to set login status!", zap.Error(err))
			response.FailWithMessage("Failed to set login status", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			UserData:  user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "Login successful", c)
	} else if err != nil {
		global.GVA_LOG.Error("Failed to set login status!", zap.Error(err))
		response.FailWithMessage("Failed to set login status", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("JWT invalidation failed", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("Failed to set login status", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			UserData:  user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "Login successful", c)
	}
}

// Register .
func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(r)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			ID: v,
		})
	}
	defaultCompanyId := utils.GetCompanyId(c)
	user := &system.SysUser{
		Username:       r.Username,
		NickName:       r.NickName,
		Password:       r.Password,
		HeaderImg:      r.HeaderImg,
		SysAuthorityId: r.AuthorityId,
		Authorities:    authorities,
		EmployeeId:     r.EmployeeId,
		CompanyID:      defaultCompanyId,
		Enable:         r.Enable,
		Phone:          r.Phone,
		Email:          r.Email,
	}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("Registration failed!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "Registration failed", c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "Registration successful", c)
}

// ChangePassword .
func (b *BaseApi) ChangePassword(c *gin.Context) {
	var req systemReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(req)
	if err != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	uid := utils.GetUserID(c)
	u := &system.SysUser{ModelId: common.ModelId{ID: uid}, Password: req.Password}
	_, err = userService.ChangePassword(u, req.NewPassword)
	if err != nil {
		global.GVA_LOG.Error("Modification failed!", zap.Error(err))
		response.FailWithMessage("Modification failed, the original password does not match the current account", c)
		return
	}
	response.OkWithMessage("Modification successful", c)
}

// GetUserList .
func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userService.GetUserInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Failed to fetch!", zap.Error(err))
		response.FailWithMessage("Failed to fetch", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Fetch successful", c)
}

// SetUserAuthority .
func (b *BaseApi) SetUserAuthority(c *gin.Context) {
	var sua systemReq.SetUserAuth
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if UserVerifyErr := utils.Verify(sua); UserVerifyErr != nil {
		response.FailWithValidationErrors(UserVerifyErr, c)
		return
	}
	userID := utils.GetUserID(c)
	err = userService.SetUserAuthority(userID, sua.AuthorityId)
	if err != nil {
		global.GVA_LOG.Error("Modification failed!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	claims := utils.GetUserInfo(c)
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // Unique signature
	claims.AuthorityId = sua.AuthorityId
	if token, err := j.CreateToken(*claims); err != nil {
		global.GVA_LOG.Error("Modification failed!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		c.Header("new-token", token)
		c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt.Unix(), 10))
		utils.SetToken(c, token, int((claims.ExpiresAt.Unix()-time.Now().Unix())/60))
		response.OkWithMessage("Modification successful", c)
	}
}

// SetUserAuthorities .
func (b *BaseApi) SetUserAuthorities(c *gin.Context) {
	var sua systemReq.SetUserAuthorities
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.SetUserAuthorities(sua.ID, sua.AuthorityIds)
	if err != nil {
		global.GVA_LOG.Error("Modification failed!", zap.Error(err))
		response.FailWithMessage("Modification failed", c)
		return
	}
	response.OkWithMessage("Modification successful", c)
}

// DeleteUser .
func (b *BaseApi) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reqId.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	jwtId := utils.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage("Deletion failed, auto-delete attempt failed", c)
		return
	}
	err = userService.DeleteUser(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

// SetUserInfo .
func (b *BaseApi) SetUserInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(user)
	if err != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}

	if len(user.AuthorityIds) != 0 {
		err = userService.SetUserAuthorities(user.ID, user.AuthorityIds)
		if err != nil {
			global.GVA_LOG.Error("Setting failed!", zap.Error(err))
			response.FailWithMessage("Setting failed", c)
			return
		}
	}
	err = userService.SetUserInfo(system.SysUser{
		ModelId: common.ModelId{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		SideMode:  user.SideMode,
		Enable:    user.Enable,
	})
	if err != nil {
		global.GVA_LOG.Error("Setting failed!", zap.Error(err))
		response.FailWithMessage("Setting failed", c)
		return
	}
	response.OkWithMessage("Setting successful", c)
}

// SetSelfInfo .
func (b *BaseApi) SetSelfInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user.ID = utils.GetUserID(c)
	err = userService.SetSelfInfo(system.SysUser{
		ModelId: common.ModelId{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		SideMode:  user.SideMode,
		Enable:    user.Enable,
	})
	if err != nil {
		global.GVA_LOG.Error("Setting failed!", zap.Error(err))
		response.FailWithMessage("Setting failed", c)
		return
	}
	response.OkWithMessage("Setting successful", c)
}

// GetUserInfo .
func (b *BaseApi) GetUserInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	ReqUser, err := userService.GetUserInfo(uuid)
	if err != nil {
		global.GVA_LOG.Error("Failed to fetch!", zap.Error(err))
		response.FailWithMessage("Failed to fetch", c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "Fetch successful", c)
}

// ResetPassword .
func (b *BaseApi) ResetPassword(c *gin.Context) {
	var user system.SysUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.ResetPassword(user.ID)
	if err != nil {
		global.GVA_LOG.Error("Reset failed!", zap.Error(err))
		response.FailWithMessage("Reset failed"+err.Error(), c)
		return
	}
	response.OkWithMessage("Reset successful", c)
}
