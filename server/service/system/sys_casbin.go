package system

import (
	"errors"
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/system/request"
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
	"strconv"
	"sync"

	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

func (casbinSer *CasbinService) UpdateCasbin(AuthorityID uint, casbinInfos []request.CasbinInfo) error {
	authorityId := strconv.Itoa(int(AuthorityID))
	casbinSer.ClearCasbin(0, authorityId)
	var rules [][]string
	deduplicateMap := make(map[string]bool)
	for _, v := range casbinInfos {
		key := authorityId + v.Path + v.Method
		if _, ok := deduplicateMap[key]; !ok {
			deduplicateMap[key] = true
			rules = append(rules, []string{authorityId, v.Path, v.Method})
		}
	}
	e := casbinSer.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("failed to add policies due to duplicate API, please contact the administrator")
	}
	return nil
}

func (casbinSer *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.GVA_DB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	e := casbinSer.Casbin()
	err = e.LoadPolicy()
	if err != nil {
		return err
	}
	return err
}

func (casbinSer *CasbinService) GetPolicyPathByAuthorityId(AuthorityID uint) (pathMaps []request.CasbinInfo) {
	e := casbinSer.Casbin()
	authorityId := strconv.Itoa(int(AuthorityID))
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

func (casbinSer *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinSer.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

func (casbinSer *CasbinService) RemoveFilteredPolicy(db *gorm.DB, authorityId string) error {
	return db.Delete(&gormadapter.CasbinRule{}, "v0 = ?", authorityId).Error
}

func (casbinSer *CasbinService) SyncPolicy(db *gorm.DB, authorityId string, rules [][]string) error {
	err := casbinSer.RemoveFilteredPolicy(db, authorityId)
	if err != nil {
		return err
	}
	return casbinSer.AddPolicies(db, rules)
}

func (casbinSer *CasbinService) AddPolicies(db *gorm.DB, rules [][]string) error {
	var casbinRules []gormadapter.CasbinRule
	for i := range rules {
		casbinRules = append(casbinRules, gormadapter.CasbinRule{
			Ptype: "p",
			V0:    rules[i][0],
			V1:    rules[i][1],
			V2:    rules[i][2],
		})
	}
	return db.Create(&casbinRules).Error
}

func (casbinSer *CasbinService) FreshCasbin() (err error) {
	e := casbinSer.Casbin()
	err = e.LoadPolicy()
	return err
}

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

func (casbinSer *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.GVA_DB)
		if err != nil {
			zap.L().Error("Failed to adapt to the database, please check if the casbin table is using the InnoDB engine!", zap.Error(err))
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("Failed to load model from string!", zap.Error(err))
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		_ = syncedCachedEnforcer.LoadPolicy()
	})
	return syncedCachedEnforcer
}
