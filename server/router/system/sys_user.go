package system

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("admin_register", baseApi.Register)               // Admin register account
		userRouter.POST("changePassword", baseApi.ChangePassword)         // User change password
		userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     // Set user authority
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)               // Delete user
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // Set user information
		userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                // Set self information
		userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // Set user authority group
		userRouter.POST("resetPassword", baseApi.ResetPassword)           // Reset user password
	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // Paginate to get user list
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)  // Get own information
	}
}
