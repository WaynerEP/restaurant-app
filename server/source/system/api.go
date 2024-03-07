package system

import (
	"context"
	sysModel "github.com/WaynerEP/restaurant-app/server/models/system"
	"github.com/WaynerEP/restaurant-app/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initApi struct{}

const initOrderApi = initOrderCompany + 1

// auto run
func init() {
	system.RegisterInit(initOrderApi, &initApi{})
}

func (i initApi) InitializerName() string {
	return sysModel.SysApi{}.TableName()
}

func (i *initApi) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysApi{})
}

func (i *initApi) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysApi{})
}

func (i *initApi) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysApi{
		{ApiGroup: "jwt", Method: "POST", Path: "/jwt/jsonInBlacklist", Description: "Add JWT to blacklist (logout, required)"},

		{ApiGroup: "System User", Method: "DELETE", Path: "/user/deleteUser", Description: "Delete user"},
		{ApiGroup: "System User", Method: "POST", Path: "/user/admin_register", Description: "User registration"},
		{ApiGroup: "System User", Method: "POST", Path: "/user/getUserList", Description: "Get user list"},
		{ApiGroup: "System User", Method: "PUT", Path: "/user/setUserInfo", Description: "Set user information"},
		{ApiGroup: "System User", Method: "PUT", Path: "/user/setSelfInfo", Description: "Set self-information (required)"},
		{ApiGroup: "System User", Method: "GET", Path: "/user/getUserInfo", Description: "Get self-information (required)"},
		{ApiGroup: "System User", Method: "POST", Path: "/user/setUserAuthorities", Description: "Set user authorities"},
		{ApiGroup: "System User", Method: "POST", Path: "/user/changePassword", Description: "Change password (recommended)"},
		{ApiGroup: "System User", Method: "POST", Path: "/user/setUserAuthority", Description: "Modify user role (required)"},
		{ApiGroup: "System User", Method: "POST", Path: "/user/resetPassword", Description: "Reset user password"},

		{ApiGroup: "api", Method: "POST", Path: "/api/createApi", Description: "Create API"},
		{ApiGroup: "api", Method: "POST", Path: "/api/deleteApi", Description: "Delete API"},
		{ApiGroup: "api", Method: "POST", Path: "/api/updateApi", Description: "Update API"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiList", Description: "Get API list"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getAllApis", Description: "Get all APIs"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiById", Description: "Get API details"},
		{ApiGroup: "api", Method: "DELETE", Path: "/api/deleteApisByIds", Description: "Batch delete APIs"},

		{ApiGroup: "Role", Method: "POST", Path: "/authority/copyAuthority", Description: "Copy role"},
		{ApiGroup: "Role", Method: "POST", Path: "/authority/createAuthority", Description: "Create role"},
		{ApiGroup: "Role", Method: "POST", Path: "/authority/deleteAuthority", Description: "Delete role"},
		{ApiGroup: "Role", Method: "PUT", Path: "/authority/updateAuthority", Description: "Update role information"},
		{ApiGroup: "Role", Method: "POST", Path: "/authority/getAuthorityList", Description: "Get role list"},
		{ApiGroup: "Role", Method: "POST", Path: "/authority/setDataAuthority", Description: "Set role resource permissions"},

		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/updateCasbin", Description: "Change role API permissions"},
		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/getPolicyPathByAuthorityId", Description: "Get permission list"},

		{ApiGroup: "Menu", Method: "POST", Path: "/menu/addBaseMenu", Description: "Add menu"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/getMenu", Description: "Get menu tree (required)"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/deleteBaseMenu", Description: "Delete menu"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/updateBaseMenu", Description: "Update menu"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/getBaseMenuById", Description: "Get menu by id"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/getMenuList", Description: "Paginated retrieval of basic menu list"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/getBaseMenuTree", Description: "Get user dynamic routes"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/getMenuAuthority", Description: "Get specified role menu"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/addMenuAuthority", Description: "Add menu and role association"},

		//{ApiGroup: "Chunk Upload", Method: "GET", Path: "/fileUploadAndDownload/findFile", Description: "Find target file (instant upload)"},
		//{ApiGroup: "Chunk Upload", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinue", Description: "Breakpoint resume"},
		//{ApiGroup: "Chunk Upload", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinueFinish", Description: "Breakpoint resume completion"},
		//{ApiGroup: "Chunk Upload", Method: "POST", Path: "/fileUploadAndDownload/removeChunk", Description: "Remove file after upload completion"},

		//{ApiGroup: "File Upload and Download", Method: "POST", Path: "/fileUploadAndDownload/upload", Description: "File upload example"},
		//{ApiGroup: "File Upload and Download", Method: "POST", Path: "/fileUploadAndDownload/deleteFile", Description: "Delete file"},
		//{ApiGroup: "File Upload and Download", Method: "POST", Path: "/fileUploadAndDownload/editFileName", Description: "Edit file name or remarks"},
		//{ApiGroup: "File Upload and Download", Method: "POST", Path: "/fileUploadAndDownload/getFileList", Description: "Get uploaded file list"},

		{ApiGroup: "System Services", Method: "POST", Path: "/system/getServerInfo", Description: "Get server information"},
		{ApiGroup: "System Services", Method: "POST", Path: "/system/getSystemConfig", Description: "Get configuration file content"},
		{ApiGroup: "System Services", Method: "POST", Path: "/system/setSystemConfig", Description: "Set configuration file content"},

		//{ApiGroup: "Customer", Method: "PUT", Path: "/contact/contact", Description: "Update contact"},
		//{ApiGroup: "Customer", Method: "POST", Path: "/contact/contact", Description: "Create contact"},
		//{ApiGroup: "Customer", Method: "DELETE", Path: "/contact/contact", Description: "Delete contact"},
		//{ApiGroup: "Customer", Method: "GET", Path: "/contact/contact", Description: "Get single contact"},
		//{ApiGroup: "Customer", Method: "GET", Path: "/contact/customerList", Description: "Get contact list"},

		//{ApiGroup: "Code Generator", Method: "GET", Path: "/autoCode/getDB", Description: "Get all databases"},
		//{ApiGroup: "Code Generator", Method: "GET", Path: "/autoCode/getTables", Description: "Get database tables"},
		//{ApiGroup: "Code Generator", Method: "POST", Path: "/autoCode/createTemp", Description: "Automate code generation"},
		//{ApiGroup: "Code Generator", Method: "POST", Path: "/autoCode/preview", Description: "Preview automated code"},
		//{ApiGroup: "Code Generator", Method: "GET", Path: "/autoCode/getColumn", Description: "Get all fields of the selected table"},
		//{ApiGroup: "Code Generator", Method: "POST", Path: "/autoCode/createPlug", Description: "Automatically create plugin package"},
		//{ApiGroup: "Code Generator", Method: "POST", Path: "/autoCode/installPlugin", Description: "Install plugin"},
		//{ApiGroup: "Code Generator", Method: "POST", Path: "/autoCode/pubPlug", Description: "Package plugin"},
		//
		//{ApiGroup: "Package (pkg) Generator", Method: "POST", Path: "/autoCode/createPackage", Description: "Generate package"},
		//{ApiGroup: "Package (pkg) Generator", Method: "POST", Path: "/autoCode/getPackage", Description: "Get all packages"},
		//{ApiGroup: "Package (pkg) Generator", Method: "POST", Path: "/autoCode/delPackage", Description: "Delete package"},
		//
		//{ApiGroup: "Code Generator History", Method: "POST", Path: "/autoCode/getMeta", Description: "Get meta information"},
		//{ApiGroup: "Code Generator History", Method: "POST", Path: "/autoCode/rollback", Description: "Rollback automatically generated code"},
		//{ApiGroup: "Code Generator History", Method: "POST", Path: "/autoCode/getSysHistory", Description: "Query rollback records"},
		//{ApiGroup: "Code Generator History", Method: "POST", Path: "/autoCode/delSysHistory", Description: "Delete rollback records"},

		//{ApiGroup: "System Dictionary Details", Method: "PUT", Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: "Update dictionary content"},
		//{ApiGroup: "System Dictionary Details", Method: "POST", Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: "Create dictionary content"},
		//{ApiGroup: "System Dictionary Details", Method: "DELETE", Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: "Delete dictionary content"},
		//{ApiGroup: "System Dictionary Details", Method: "GET", Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: "Get dictionary content by ID"},
		//{ApiGroup: "System Dictionary Details", Method: "GET", Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: "Get dictionary content list"},
		//
		//{ApiGroup: "System Dictionary", Method: "POST", Path: "/sysDictionary/createSysDictionary", Description: "Create dictionary"},
		//{ApiGroup: "System Dictionary", Method: "DELETE", Path: "/sysDictionary/deleteSysDictionary", Description: "Delete dictionary"},
		//{ApiGroup: "System Dictionary", Method: "PUT", Path: "/sysDictionary/updateSysDictionary", Description: "Update dictionary"},
		//{ApiGroup: "System Dictionary", Method: "GET", Path: "/sysDictionary/findSysDictionary", Description: "Get dictionary by ID"},
		//{ApiGroup: "System Dictionary", Method: "GET", Path: "/sysDictionary/getSysDictionaryList", Description: "Get dictionary list"},
		{ApiGroup: "Operation Records", Method: "POST", Path: "/sysOperationRecord/createSysOperationRecord", Description: "Create operation record"},
		{ApiGroup: "Operation Records", Method: "GET", Path: "/sysOperationRecord/findSysOperationRecord", Description: "Get operation record by ID"},
		{ApiGroup: "Operation Records", Method: "GET", Path: "/sysOperationRecord/getSysOperationRecordList", Description: "Get operation record list"},
		{ApiGroup: "Operation Records", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecord", Description: "Delete operation record"},
		{ApiGroup: "Operation Records", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: "Batch delete operation history"},
		//
		//{ApiGroup: "Resumable Upload (Plugin Version)", Method: "POST", Path: "/simpleUploader/upload", Description: "Plugin version chunked upload"},
		//{ApiGroup: "Resumable Upload (Plugin Version)", Method: "GET", Path: "/simpleUploader/checkFileMd5", Description: "File integrity verification"},
		//{ApiGroup: "Resumable Upload (Plugin Version)", Method: "GET", Path: "/simpleUploader/mergeFileMd5", Description: "Merge file after upload completion"},

		{ApiGroup: "Email", Method: "POST", Path: "/email/emailTest", Description: "Send test email"},
		{ApiGroup: "Email", Method: "POST", Path: "/email/emailSend", Description: "Send email example"},

		{ApiGroup: "Button Permissions", Method: "POST", Path: "/authorityBtn/setAuthorityBtn", Description: "Set button permissions"},
		{ApiGroup: "Button Permissions", Method: "POST", Path: "/authorityBtn/getAuthorityBtn", Description: "Get existing button permissions"},
		{ApiGroup: "Button Permissions", Method: "POST", Path: "/authorityBtn/canRemoveAuthorityBtn", Description: "Remove button"},

		//{ApiGroup: "Table Templates", Method: "POST", Path: "/sysExportTemplate/createSysExportTemplate", Description: "Create export template"},
		//{ApiGroup: "Table Templates", Method: "DELETE", Path: "/sysExportTemplate/deleteSysExportTemplate", Description: "Delete export template"},
		//{ApiGroup: "Table Templates", Method: "DELETE", Path: "/sysExportTemplate/deleteSysExportTemplateByIds", Description: "Batch delete export template"},
		//{ApiGroup: "Table Templates", Method: "PUT", Path: "/sysExportTemplate/updateSysExportTemplate", Description: "Update export template"},
		//{ApiGroup: "Table Templates", Method: "GET", Path: "/sysExportTemplate/findSysExportTemplate", Description: "Get export template by ID"},
		//{ApiGroup: "Table Templates", Method: "GET", Path: "/sysExportTemplate/getSysExportTemplateList", Description: "Get export template list"},
		//{ApiGroup: "Table Templates", Method: "GET", Path: "/sysExportTemplate/exportExcel", Description: "Export Excel"},
		//{ApiGroup: "Table Templates", Method: "GET", Path: "/sysExportTemplate/exportTemplate", Description: "Download template"},
		//{ApiGroup: "Table Templates", Method: "POST", Path: "/sysExportTemplate/importExcel", Description: "Import Excel"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysApi{}.TableName()+" table data initialization failed!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initApi) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ? AND method = ?", "/authorityBtn/canRemoveAuthorityBtn", "POST").
		First(&sysModel.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
