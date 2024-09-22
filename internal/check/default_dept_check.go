package check

import (
	"github.com/YuanJey/sync_modle/pkg/db/mysql/structs"
	"github.com/YuanJey/sync_modle/pkg/log"
	wpsApi "github.com/YuanJey/wps-api"
	"github.com/YuanJey/wps-api/dept"
	"github.com/YuanJey/wps-api/pkg/api_resp"
)

type DefaultDeptCheck struct {
	WpsData *dept.AllWpsDept
}

func NewDefaultDeptCheck() *DefaultDeptCheck {
	check := DefaultDeptCheck{}
	check.SetWpsData("init WpsData")
	return &check
}

func (d *DefaultDeptCheck) SetWpsData(operationID string) {
	wpsDept, err := wpsApi.Sdk.Dept.GetAllWpsDept(operationID)
	if err != nil {
		log.Error(operationID, "get wps dept list failed", err.Error())
		return
	}
	d.WpsData = wpsDept
	log.Info(operationID, "set wps data success")
	log.Info(operationID, "root dept info :", wpsDept.RootDept)
	log.Info(operationID, "wps dept len :", len(wpsDept.WidDeptList))
	log.Info(operationID, "wps third len :", len(wpsDept.TidDeptList))
}
func (d *DefaultDeptCheck) IsCreate(operationID string, thirdDept *structs.Organization) bool {
	if _, ok := d.WpsData.TidDeptList[thirdDept.ID]; ok {
		return false
	}
	log.Info(operationID, "create dept", *thirdDept)
	return true
}

func (d *DefaultDeptCheck) IsUpdate(operationID string, thirdDept *structs.Organization) bool {
	if wd, ok := d.WpsData.TidDeptList[thirdDept.ID]; ok {
		if wd.Name != thirdDept.Name || wd.Weight != thirdDept.Weight {
			log.Info(operationID, "update dept", *thirdDept)
			return true
		}
	}
	return false
}

func (d *DefaultDeptCheck) IsMove(operationID string, thirdDept *structs.Organization) bool {
	if wpsDept, ok := d.WpsData.TidDeptList[thirdDept.ID]; ok {
		parentWpsDept, err := wpsApi.Sdk.Dept.GetDeptInfoByThirdId(operationID, thirdDept.ParentId)
		if err != nil {
			return false
		}
		if parentWpsDept.Data.Id != wpsDept.ParentId {
			log.Info(operationID, "move dept", *thirdDept)
			return true
		}
	}
	return false
}

func (d *DefaultDeptCheck) IsDelete(operationID string, wpsDept *api_resp.Dept) bool {
	if wpsDept.Synced {
		return false
	}
	log.Info(operationID, "delete dept", *wpsDept)
	return true
}
