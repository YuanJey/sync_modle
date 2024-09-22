package check

import (
	"github.com/YuanJey/sync_modle/pkg/base_info"
	"github.com/YuanJey/sync_modle/pkg/db/mysql/structs"
)
import "github.com/YuanJey/wps-api/pkg/api_resp"

type DeptLogic interface {
	IsCreate(operationID string, thirdDept *structs.Organization) bool
	IsUpdate(operationID string, thirdDept *structs.Organization) bool
	IsMove(operationID string, thirdDept *structs.Organization) bool
	IsDelete(operationID string, wpsDept *api_resp.Dept) bool
}
type UserLogic interface {
	IsCreate(operationID string, thirdUser *base_info.ThirdMember) bool
	IsUpdate(operationID string, thirdUser *base_info.ThirdMember) bool
	IsMove(operationID string, thirdUser *base_info.ThirdMember) bool
	IsDisable(operationID string, thirdUser *base_info.ThirdMember) bool
	IsEnable(operationID string, thirdUser *base_info.ThirdMember) bool
	IsDelete(operationID string, thirdUser *base_info.ThirdMember) bool
}
