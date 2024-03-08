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

		{ApiGroup: "System Services", Method: "POST", Path: "/system/getServerInfo", Description: "Get server information"},
		{ApiGroup: "System Services", Method: "POST", Path: "/system/getSystemConfig", Description: "Get configuration file content"},
		{ApiGroup: "System Services", Method: "POST", Path: "/system/setSystemConfig", Description: "Set configuration file content"},

		{ApiGroup: "Operation Records", Method: "POST", Path: "/sysOperationRecord/createSysOperationRecord", Description: "Create operation record"},
		{ApiGroup: "Operation Records", Method: "GET", Path: "/sysOperationRecord/findSysOperationRecord", Description: "Get operation record by ID"},
		{ApiGroup: "Operation Records", Method: "GET", Path: "/sysOperationRecord/getSysOperationRecordList", Description: "Get operation record list"},
		{ApiGroup: "Operation Records", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecord", Description: "Delete operation record"},
		{ApiGroup: "Operation Records", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: "Batch delete operation history"},

		{ApiGroup: "Email", Method: "POST", Path: "/email/emailTest", Description: "Send test email"},
		{ApiGroup: "Email", Method: "POST", Path: "/email/emailSend", Description: "Send email example"},

		{ApiGroup: "Button Permissions", Method: "POST", Path: "/authorityBtn/setAuthorityBtn", Description: "Set button permissions"},
		{ApiGroup: "Button Permissions", Method: "POST", Path: "/authorityBtn/getAuthorityBtn", Description: "Get existing button permissions"},
		{ApiGroup: "Button Permissions", Method: "POST", Path: "/authorityBtn/canRemoveAuthorityBtn", Description: "Remove button"},

		{ApiGroup: "Customer", Method: "POST", Path: "/customer/createCustomer", Description: "Crear cliente"},
		{ApiGroup: "Customer", Method: "PUT", Path: "/customer/updateCustomer", Description: "Actualizar cliente"},
		{ApiGroup: "Customer", Method: "DELETE", Path: "/customer/deleteCustomer", Description: "Eliminar cliente"},
		{ApiGroup: "Customer", Method: "GET", Path: "/customer/getCustomer", Description: "Obtener cliente"},
		{ApiGroup: "Customer", Method: "GET", Path: "/customer/customerList", Description: "Obtener lista de clientes"},

		{ApiGroup: "Unit of Measure", Method: "POST", Path: "/unitOfMeasure/createUnitOfMeasure", Description: "Crear unidad de medida"},
		{ApiGroup: "Unit of Measure", Method: "PUT", Path: "/unitOfMeasure/updateUnitOfMeasure", Description: "Actualizar unidad de medida"},
		{ApiGroup: "Unit of Measure", Method: "DELETE", Path: "/unitOfMeasure/deleteUnitOfMeasure", Description: "Eliminar unidad de medida"},
		{ApiGroup: "Unit of Measure", Method: "GET", Path: "/unitOfMeasure/getUnitOfMeasure", Description: "Obtener unidad de medida"},
		{ApiGroup: "Unit of Measure", Method: "GET", Path: "/unitOfMeasure/getUnitOfMeasureList", Description: "Obtener lista de unidades de medida"},

		{ApiGroup: "Supply Category", Method: "POST", Path: "/supplyCategory/createSupplyCategory", Description: "Crear categoría de suministro"},
		{ApiGroup: "Supply Category", Method: "PUT", Path: "/supplyCategory/updateSupplyCategory", Description: "Actualizar categoría de suministro"},
		{ApiGroup: "Supply Category", Method: "DELETE", Path: "/supplyCategory/deleteSupplyCategory", Description: "Eliminar categoría de suministro"},
		{ApiGroup: "Supply Category", Method: "GET", Path: "/supplyCategory/getSupplyCategory", Description: "Obtener categoría de suministro"},
		{ApiGroup: "Supply Category", Method: "GET", Path: "/supplyCategory/getSupplyCategoryList", Description: "Obtener lista de categorías de suministro"},

		{ApiGroup: "Supply", Method: "POST", Path: "/supply/createSupply", Description: "Crear suministro"},
		{ApiGroup: "Supply", Method: "PUT", Path: "/supply/updateSupply", Description: "Actualizar suministro"},
		{ApiGroup: "Supply", Method: "DELETE", Path: "/supply/deleteSupply", Description: "Eliminar suministro"},
		{ApiGroup: "Supply", Method: "GET", Path: "/supply/getSupply", Description: "Obtener suministro"},
		{ApiGroup: "Supply", Method: "GET", Path: "/supply/getSupplyList", Description: "Obtener lista de suministros"},

		{ApiGroup: "Item Category", Method: "POST", Path: "/itemCategory/createItemCategory", Description: "Crear categoría de ítem"},
		{ApiGroup: "Item Category", Method: "PUT", Path: "/itemCategory/updateItemCategory", Description: "Actualizar categoría de ítem"},
		{ApiGroup: "Item Category", Method: "DELETE", Path: "/itemCategory/deleteItemCategory", Description: "Eliminar categoría de ítem"},
		{ApiGroup: "Item Category", Method: "GET", Path: "/itemCategory/getItemCategory", Description: "Obtener categoría de ítem"},
		{ApiGroup: "Item Category", Method: "GET", Path: "/itemCategory/getItemCategoryList", Description: "Obtener lista de categorías de ítem"},

		{ApiGroup: "Item", Method: "POST", Path: "/item/createItem", Description: "Crear ítem"},
		{ApiGroup: "Item", Method: "PUT", Path: "/item/updateItem", Description: "Actualizar ítem"},
		{ApiGroup: "Item", Method: "DELETE", Path: "/item/deleteItem", Description: "Eliminar ítem"},
		{ApiGroup: "Item", Method: "GET", Path: "/item/getItem", Description: "Obtener ítem"},
		{ApiGroup: "Item", Method: "GET", Path: "/item/getItemList", Description: "Obtener lista de ítems"},

		{ApiGroup: "Floor", Method: "POST", Path: "/floor/createFloor", Description: "Crear piso"},
		{ApiGroup: "Floor", Method: "PUT", Path: "/floor/updateFloor", Description: "Actualizar piso"},
		{ApiGroup: "Floor", Method: "DELETE", Path: "/floor/deleteFloor", Description: "Eliminar piso"},
		{ApiGroup: "Floor", Method: "GET", Path: "/floor/getFloor", Description: "Obtener piso"},
		{ApiGroup: "Floor", Method: "GET", Path: "/floor/getFloorList", Description: "Obtener lista de pisos"},
		{ApiGroup: "Floor", Method: "GET", Path: "/floor/getAllList", Description: "Obtener lista completa"},
		{ApiGroup: "Floor", Method: "GET", Path: "/floor/getTreeFloor", Description: "Obtener estructura de árbol de pisos"},

		{ApiGroup: "Environment", Method: "POST", Path: "/environment/createEnvironment", Description: "Crear ambiente de mesa"},
		{ApiGroup: "Environment", Method: "PUT", Path: "/environment/updateEnvironment", Description: "Actualizar ambiente de mesa"},
		{ApiGroup: "Environment", Method: "DELETE", Path: "/environment/deleteEnvironment", Description: "Eliminar ambiente de mesa"},
		{ApiGroup: "Environment", Method: "GET", Path: "/environment/getEnvironment", Description: "Obtener ambiente de mesa"},
		{ApiGroup: "Environment", Method: "GET", Path: "/environment/getEnvironmentList", Description: "Obtener lista de ambientes de mesa"},

		{ApiGroup: "Table", Method: "POST", Path: "/table/createTable", Description: "Crear mesa"},
		{ApiGroup: "Table", Method: "PUT", Path: "/table/updateTable", Description: "Actualizar mesa"},
		{ApiGroup: "Table", Method: "DELETE", Path: "/table/deleteTable", Description: "Eliminar mesa"},
		{ApiGroup: "Table", Method: "GET", Path: "/table/getTable", Description: "Obtener mesa"},
		{ApiGroup: "Table", Method: "GET", Path: "/table/getTableList", Description: "Obtener lista de mesas"},

		{ApiGroup: "Order", Method: "POST", Path: "/order/createOrder", Description: "Crear orden"},
		{ApiGroup: "Order", Method: "PUT", Path: "/order/updateOrder", Description: "Actualizar orden"},
		{ApiGroup: "Order", Method: "DELETE", Path: "/order/deleteOrder", Description: "Eliminar orden"},
		{ApiGroup: "Order", Method: "GET", Path: "/order/getOrder", Description: "Obtener orden"},
		{ApiGroup: "Order", Method: "GET", Path: "/order/getOrderList", Description: "Obtener lista de órdenes"},

		// {ApiGroup: "Payment", Method: "POST", Path: "/payment/createPayment", Description: "Crear pago"},
		// {ApiGroup: "Payment", Method: "PUT", Path: "/payment/updatePayment", Description: "Actualizar pago"},
		// {ApiGroup: "Payment", Method: "DELETE", Path: "/payment/deletePayment", Description: "Eliminar pago"},
		// {ApiGroup: "Payment", Method: "GET", Path: "/payment/getPayment", Description: "Obtener pago"},
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
