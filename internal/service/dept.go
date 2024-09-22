package service

import (
	"github.com/YuanJey/sync_modle/pkg/config"
	"github.com/YuanJey/sync_modle/pkg/db/mysql/structs"
	"github.com/YuanJey/sync_modle/pkg/log"
	wpsApi "github.com/YuanJey/wps-api"
	"github.com/YuanJey/wps-api/pkg/api_req"
)

func CreateDept(operationID string, dept *structs.Organization) {
	var infos []api_req.CreateDeptInfo
	infos = append(infos, api_req.CreateDeptInfo{
		Name:            dept.Name,
		Source:          "sync",
		ThirdPlatformId: config.Config.WPS.PlatformId,
		ThirdUnionId:    dept.ID,
		Weight:          dept.Weight,
	})
	parentDept, err := wpsApi.Sdk.Dept.GetDeptInfoByThirdId(operationID, dept.ParentId)
	if err != nil {
		return
	}
	req := api_req.BatchCreateDeptReq{Detps: infos}
	createDept, err := wpsApi.Sdk.Dept.BatchCreateDept(operationID, parentDept.Data.Id, req)
	if err != nil {
		log.Error(operationID, "create dept failed", err)
		return
	}
	log.Info(operationID, "create dept success", createDept)
}
func DeleteDept(operationID string, thirdId string) {
	req := api_req.BatchDeleteThirdDeptReq{
		PlatformId: config.Config.WPS.PlatformId,
		UnionIds:   []string{thirdId},
	}
	dept, err := wpsApi.Sdk.Dept.BatchDeleteThirdDept(operationID, req)
	if err != nil {
		log.Error(operationID, "delete dept failed", err)
		return
	}
	log.Info(operationID, "delete dept success", dept)
}
func UpdateDept(operationID string, dept *structs.Organization) {
	req := api_req.UpdateDeptReq{
		Name:   dept.Name,
		Weight: dept.Weight,
	}
	wpsDept, err := wpsApi.Sdk.Dept.GetDeptInfoByThirdId(operationID, dept.ID)
	if err != nil {
		return
	}
	info, err := wpsApi.Sdk.Dept.UpdateDeptInfo(operationID, wpsDept.Data.Id, req)
	if err != nil {
		log.Error(operationID, "update dept failed", err)
		return
	}
	log.Info(operationID, "update dept success", info)
}
func MoveDept(operationID string, dept *structs.Organization) {
	wpsParentDept, err := wpsApi.Sdk.Dept.GetDeptInfoByThirdId(operationID, dept.ParentId)
	if err != nil {
		return
	}
	wpsDept, err := wpsApi.Sdk.Dept.GetDeptInfoByThirdId(operationID, dept.ID)
	if err != nil {
		return
	}
	req := api_req.MoveDeptReq{
		DeptId:     wpsDept.Data.Id,
		ToParentId: wpsParentDept.Data.Id,
	}
	moveDept, err := wpsApi.Sdk.Dept.MoveDept(operationID, req)
	if err != nil {
		log.Error(operationID, "move dept failed", err)
		return
	}
	log.Info(operationID, "move dept success", moveDept)
}
