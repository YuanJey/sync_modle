package check

import (
	"github.com/YuanJey/sync_modle/pkg/base_info"
	"github.com/YuanJey/sync_modle/pkg/consts"
	"github.com/YuanJey/sync_modle/pkg/log"
	wpsApi "github.com/YuanJey/wps-api"
	"github.com/YuanJey/wps-api/dept"
	"github.com/YuanJey/wps-api/user"
)

type DefaultUserCheck struct {
	wpsUserMap *user.AllWpsUser
	wpsDeptMap *dept.AllWpsDept
}

func NewDefaultUserCheck() *DefaultUserCheck {
	check := DefaultUserCheck{}
	check.SetWpsData("init WpsData")
	return &check
}
func (d *DefaultUserCheck) SetWpsData(operationID string) {
	wpsDept, err := wpsApi.Sdk.Dept.GetAllWpsDept(operationID)
	if err != nil {
		log.Error(operationID, "get wps dept list failed", err.Error())
		return
	}
	d.wpsDeptMap = wpsDept
	log.Info(operationID, "set wps data success")
	log.Info(operationID, "root dept info :", wpsDept.RootDept)
	log.Info(operationID, "wps dept len :", len(wpsDept.WidDeptList))
	log.Info(operationID, "third dept len :", len(wpsDept.TidDeptList))
	allUser, err := wpsApi.Sdk.User.GetAllUser(operationID, wpsDept)
	if err != nil {
		log.Error(operationID, "get wps user list failed", err.Error())
		return
	}
	d.wpsUserMap = allUser
	log.Info(operationID, "wps user len :", len(d.wpsUserMap.WidUserList))
	log.Info(operationID, "third user len :", len(d.wpsUserMap.TidUserList))
}
func (d *DefaultUserCheck) IsCreate(operationID string, thirdUser *base_info.ThirdMember) bool {
	if _, ok := d.wpsUserMap.TidUserList[thirdUser.ThirdUnionId]; ok {
		log.Info(operationID, " user create ", *thirdUser)
		return false
	}
	return true
}

func (d *DefaultUserCheck) IsUpdate(operationID string, thirdUser *base_info.ThirdMember) bool {
	if wpsUser, ok := d.wpsUserMap.TidUserList[thirdUser.ThirdUnionId]; ok {
		if wpsUser.NickName != thirdUser.NickName {
			log.Info(operationID, " user update ", *thirdUser)
			return true
		}
	}
	return false
}

func (d *DefaultUserCheck) IsMove(operationID string, thirdUser *base_info.ThirdMember) bool {
	if wpsUser, ok := d.wpsUserMap.TidUserList[thirdUser.ThirdUnionId]; ok {
		wpsDept, err := wpsApi.Sdk.Dept.GetDeptInfoByThirdId(operationID, thirdUser.ThirdDeptId)
		if err != nil {
			return false
		}
		if wpsUser.DefDeptId != wpsDept.Data.Id {
			log.Info(operationID, " user move ", *thirdUser)
			return true
		}
	}
	return false
}

func (d *DefaultUserCheck) IsDisable(operationID string, thirdUser *base_info.ThirdMember) bool {
	if wpsUser, ok := d.wpsUserMap.TidUserList[thirdUser.ThirdUnionId]; ok {
		if wpsUser.Status != consts.WpsUserStatusDisabled && thirdUser.Status == consts.ThirdUserStatus.Disable() {
			log.Info(operationID, " user disable ", *thirdUser)
			return true
		}
	}
	return false
}

func (d *DefaultUserCheck) IsEnable(operationID string, thirdUser *base_info.ThirdMember) bool {
	if wpsUser, ok := d.wpsUserMap.TidUserList[thirdUser.ThirdUnionId]; ok {
		if wpsUser.Status == consts.WpsUserStatusDisabled && thirdUser.Status == consts.ThirdUserStatus.Enable() {
			log.Info(operationID, " user enable ", *thirdUser)
			return true
		}
	}
	return false
}

func (d *DefaultUserCheck) IsDelete(operationID string, thirdUser *base_info.ThirdMember) bool {
	if _, ok := d.wpsUserMap.TidUserList[thirdUser.ThirdUnionId]; ok {
		if thirdUser.Status == consts.ThirdUserStatus.Delete() {
			log.Info(operationID, " user delete ", *thirdUser)
			return true
		}
	}
	return false
}
