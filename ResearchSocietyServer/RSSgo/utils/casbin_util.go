package utils

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"hkc.ink/rss/model"
)

type CasbinUtil struct {
}

func (CasbinUtil) Enforce(db *gorm.DB, sub string, obj string, act string) bool {
	a, _ := gormadapter.NewAdapterByDBWithCustomTable(db, &model.CasbinRule{})
	e, _ := casbin.NewEnforcer("config/casbin_model.conf", a)

	if res, _ := e.Enforce(sub, obj, act); res {
		return true
	} else {
		return false
	}
}
