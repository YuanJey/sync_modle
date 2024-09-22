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
