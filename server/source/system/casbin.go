package system

import (
	"context"

	"github.com/WaynerEP/restaurant-app/server/service/system"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderCasbin = initOrderApi + 1

type initCasbin struct{}

// auto run
func init() {
	system.RegisterInit(initOrderCasbin, &initCasbin{})
}

func (i *initCasbin) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&adapter.CasbinRule{})
}

func (i *initCasbin) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&adapter.CasbinRule{})
}

func (i initCasbin) InitializerName() string {
	var entity adapter.CasbinRule
	return entity.TableName()
}

func (i *initCasbin) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []adapter.CasbinRule{
		{Ptype: "p", V0: "1", V1: "/user/admin_register", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/api/createApi", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/api/getApiList", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/api/getApiById", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/api/deleteApi", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/api/updateApi", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/api/getAllApis", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/api/deleteApisByIds", V2: "DELETE"},

		{Ptype: "p", V0: "1", V1: "/authority/copyAuthority", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/authority/updateAuthority", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/authority/createAuthority", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/authority/deleteAuthority", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/authority/getAuthorityList", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/authority/setDataAuthority", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/menu/getMenu", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/menu/getMenuList", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/menu/addBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/menu/getBaseMenuTree", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/menu/addMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/menu/getMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/menu/deleteBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/menu/updateBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/menu/getBaseMenuById", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/user/getUserInfo", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/user/setUserInfo", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/user/setSelfInfo", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/user/getUserList", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/deleteUser", V2: "DELETE"},
		{Ptype: "p", V0: "1", V1: "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/setUserAuthority", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/setUserAuthorities", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/resetPassword", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/casbin/updateCasbin", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/jwt/jsonInBlacklist", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/system/getSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/system/setSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/system/getServerInfo", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/autoCode/getDB", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/autoCode/getMeta", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/autoCode/preview", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/autoCode/getTables", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/autoCode/getColumn", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/autoCode/rollback", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/autoCode/createTemp", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/autoCode/delSysHistory", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/autoCode/getSysHistory", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/autoCode/createPackage", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/autoCode/getPackage", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/autoCode/delPackage", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/autoCode/createPlug", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/autoCode/installPlugin", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/autoCode/pubPlug", V2: "POST"},

		//{Ptype: "p", V0: "1", V1: "/sysDictionaryDetail/findSysDictionaryDetail", V2: "GET"},
		//{Ptype: "p", V0: "1", V1: "/sysDictionaryDetail/updateSysDictionaryDetail", V2: "PUT"},
		//{Ptype: "p", V0: "1", V1: "/sysDictionaryDetail/createSysDictionaryDetail", V2: "POST"},
		//{Ptype: "p", V0: "1", V1: "/sysDictionaryDetail/getSysDictionaryDetailList", V2: "GET"},
		//{Ptype: "p", V0: "1", V1: "/sysDictionaryDetail/deleteSysDictionaryDetail", V2: "DELETE"},
		//
		//{Ptype: "p", V0: "1", V1: "/sysDictionary/findSysDictionary", V2: "GET"},
		//{Ptype: "p", V0: "1", V1: "/sysDictionary/updateSysDictionary", V2: "PUT"},
		//{Ptype: "p", V0: "1", V1: "/sysDictionary/getSysDictionaryList", V2: "GET"},
		//{Ptype: "p", V0: "1", V1: "/sysDictionary/createSysDictionary", V2: "POST"},
		//{Ptype: "p", V0: "1", V1: "/sysDictionary/deleteSysDictionary", V2: "DELETE"},

		{Ptype: "p", V0: "1", V1: "/sysOperationRecord/findSysOperationRecord", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/sysOperationRecord/updateSysOperationRecord", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/sysOperationRecord/createSysOperationRecord", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/sysOperationRecord/getSysOperationRecordList", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/sysOperationRecord/deleteSysOperationRecord", V2: "DELETE"},
		{Ptype: "p", V0: "1", V1: "/sysOperationRecord/deleteSysOperationRecordByIds", V2: "DELETE"},

		{Ptype: "p", V0: "1", V1: "/email/emailTest", V2: "POST"},

		//{Ptype: "p", V0: "1", V1: "/simpleUploader/upload", V2: "POST"},
		//{Ptype: "p", V0: "1", V1: "/simpleUploader/checkFileMd5", V2: "GET"},
		//{Ptype: "p", V0: "1", V1: "/simpleUploader/mergeFileMd5", V2: "GET"},

		{Ptype: "p", V0: "1", V1: "/authorityBtn/setAuthorityBtn", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/authorityBtn/getAuthorityBtn", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/authorityBtn/canRemoveAuthorityBtn", V2: "POST"},

		//{Ptype: "p", V0: "1", V1: "/sysExportTemplate/createSysExportTemplate", V2: "POST"},
		//{Ptype: "p", V0: "1", V1: "/sysExportTemplate/deleteSysExportTemplate", V2: "DELETE"},
		//{Ptype: "p", V0: "1", V1: "/sysExportTemplate/deleteSysExportTemplateByIds", V2: "DELETE"},
		//{Ptype: "p", V0: "1", V1: "/sysExportTemplate/updateSysExportTemplate", V2: "PUT"},
		//{Ptype: "p", V0: "1", V1: "/sysExportTemplate/findSysExportTemplate", V2: "GET"},
		//{Ptype: "p", V0: "1", V1: "/sysExportTemplate/getSysExportTemplateList", V2: "GET"},
		//{Ptype: "p", V0: "1", V1: "/sysExportTemplate/exportExcel", V2: "GET"},
		//{Ptype: "p", V0: "1", V1: "/sysExportTemplate/exportTemplate", V2: "GET"},
		//{Ptype: "p", V0: "1", V1: "/sysExportTemplate/importExcel", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/customer/createCustomer", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/customer/updateCustomer", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/customer/deleteCustomer", V2: "DELETE"},
		{Ptype: "p", V0: "1", V1: "/customer/getCustomer", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/customer/getCustomerList", V2: "GET"},

		{Ptype: "p", V0: "1", V1: "/unitOfMeasure/createUnitOfMeasure", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/unitOfMeasure/updateUnitOfMeasure", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/unitOfMeasure/deleteUnitOfMeasure", V2: "DELETE"},
		{Ptype: "p", V0: "1", V1: "/unitOfMeasure/getUnitOfMeasure", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/unitOfMeasure/getUnitOfMeasureList", V2: "GET"},

		{Ptype: "p", V0: "1", V1: "/supplyCategory/createSupplyCategory", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/supplyCategory/updateSupplyCategory", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/supplyCategory/deleteSupplyCategory", V2: "DELETE"},
		{Ptype: "p", V0: "1", V1: "/supplyCategory/getSupplyCategory", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/supplyCategory/getSupplyCategoryList", V2: "GET"},

		{Ptype: "p", V0: "1", V1: "/supply/createSupply", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/supply/updateSupply", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/supply/deleteSupply", V2: "DELETE"},
		{Ptype: "p", V0: "1", V1: "/supply/getSupply", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/supply/getSupplyList", V2: "GET"},

		{Ptype: "p", V0: "1", V1: "/itemCategory/createItemCategory", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/itemCategory/updateItemCategory", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/itemCategory/deleteItemCategory", V2: "DELETE"},
		{Ptype: "p", V0: "1", V1: "/itemCategory/getItemCategory", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/itemCategory/getItemCategoryList", V2: "GET"},

		{Ptype: "p", V0: "1", V1: "/item/createItem", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/item/updateItem", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/item/deleteItem", V2: "DELETE"},
		{Ptype: "p", V0: "1", V1: "/item/getItem", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/item/getItemList", V2: "GET"},

		{Ptype: "p", V0: "1", V1: "/floor/createFloor", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/floor/updateFloor", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/floor/deleteFloor", V2: "DELETE"},
		{Ptype: "p", V0: "1", V1: "/floor/getFloor", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/floor/getFloorList", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/floor/getAllList", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/floor/getTreeFloor", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/floor/getOptionsForSelect", V2: "GET"},

		{Ptype: "p", V0: "1", V1: "/environment/createEnvironment", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/environment/createFloorEnvironment", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/environment/updateEnvironment", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/environment/deleteEnvironment", V2: "DELETE"},
		{Ptype: "p", V0: "1", V1: "/environment/deleteFloorEnvironment", V2: "DELETE"},
		{Ptype: "p", V0: "1", V1: "/environment/getEnvironment", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/environment/getEnvironmentList", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/environment/getOptionsForSelect", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/environment/getEnvironmentsByFloorId/:id", V2: "GET"},

		{Ptype: "p", V0: "1", V1: "/table/createTable", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/table/createFloorEnvTable", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/table/updateTable", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/table/deleteTable", V2: "DELETE"},
		{Ptype: "p", V0: "1", V1: "/table/deleteFloorEnvironmentTable", V2: "DELETE"},
		{Ptype: "p", V0: "1", V1: "/table/getTable", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/table/getTableList", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/table/getOptionsForSelect", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/table/getTablesByFloorEnvironmentId/:id", V2: "GET"},

		{Ptype: "p", V0: "1", V1: "/menuOrder/readyMenuOrder", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/menuOrder/approveMenuOrder", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/menuOrder/rejectMenuOrder", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/menuOrder/updateStatusMenuOrder", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/menuOrder/createMenuOrder", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/menuOrder/updateMenuOrder", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/menuOrder/deleteMenuOrder", V2: "DELETE"},
		{Ptype: "p", V0: "1", V1: "/menuOrder/getMenuOrder", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/menuOrder/getMenuOrderList", V2: "GET"},

		{Ptype: "p", V0: "2", V1: "/user/admin_register", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/api/createApi", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/api/getApiList", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/api/getApiById", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/api/deleteApi", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/api/updateApi", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/api/getAllApis", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/authority/createAuthority", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/authority/deleteAuthority", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/authority/getAuthorityList", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/authority/setDataAuthority", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/menu/getMenu", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/menu/getMenuList", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/menu/addBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/menu/getBaseMenuTree", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/menu/addMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/menu/getMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/menu/deleteBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/menu/updateBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/menu/getBaseMenuById", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/user/getUserList", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/user/setUserAuthority", V2: "POST"},
		//{Ptype: "p", V0: "2", V1: "/fileUploadAndDownload/upload", V2: "POST"},
		//{Ptype: "p", V0: "2", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
		//{Ptype: "p", V0: "2", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
		//{Ptype: "p", V0: "2", V1: "/fileUploadAndDownload/editFileName", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/casbin/updateCasbin", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/jwt/jsonInBlacklist", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/system/getSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/system/setSystemConfig", V2: "POST"},

		{Ptype: "p", V0: "3", V1: "/user/admin_register", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/api/createApi", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/api/getApiList", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/api/getApiById", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/api/deleteApi", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/api/updateApi", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/api/getAllApis", V2: "POST"},

		{Ptype: "p", V0: "3", V1: "/authority/createAuthority", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/authority/deleteAuthority", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/authority/getAuthorityList", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/authority/setDataAuthority", V2: "POST"},

		{Ptype: "p", V0: "3", V1: "/menu/getMenu", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/menu/getMenuList", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/menu/addBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/menu/getBaseMenuTree", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/menu/addMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/menu/getMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/menu/deleteBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/menu/updateBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/menu/getBaseMenuById", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/user/getUserList", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/user/setUserAuthority", V2: "POST"},

		{Ptype: "p", V0: "3", V1: "/casbin/updateCasbin", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/jwt/jsonInBlacklist", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/system/getSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/system/setSystemConfig", V2: "POST"},

		{Ptype: "p", V0: "3", V1: "/user/getUserInfo", V2: "GET"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, "Casbin 表 ("+i.InitializerName()+") 数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initCasbin) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where(adapter.CasbinRule{Ptype: "p", V0: "3", V1: "/user/getUserInfo", V2: "GET"}).
		First(&adapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
