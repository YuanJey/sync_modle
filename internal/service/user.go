package service

import (
	"github.com/YuanJey/sync_modle/pkg/base_info"
	"github.com/YuanJey/sync_modle/pkg/config"
	"github.com/YuanJey/sync_modle/pkg/log"
	wpsApi "github.com/YuanJey/wps-api"
	"github.com/YuanJey/wps-api/pkg/api_req"
)

func CreateUser(operationID string, user *base_info.ThirdMember) {
	wpsDept, err := wpsApi.Sdk.Dept.GetDeptInfoByThirdId(operationID, user.ThirdDeptId)
	if err != nil {
		return
	}
	req := api_req.CreateCompaniesMembersReq{
		Account:   user.ThirdUnionId,
		DefDeptId: wpsDept.Data.Id,
		DeptIds: []api_req.DeptIds{
			{
				DeptId: wpsDept.Data.Id,
				Weight: 0,
			},
		},
		NickName:        user.NickName,
		Source:          "sync",
		ThirdUnionId:    user.ThirdUnionId,
		ThirdPlatformId: config.Config.WPS.PlatformId,
	}
	members, err := wpsApi.Sdk.User.CreateCompaniesMembers(operationID, req)
	if err != nil {
		log.Error("CreateUser", "CreateCompaniesMembers failed ", err.Error(), req)
		return
	}
	log.Info("CreateUser", "CreateCompaniesMembers success ", members)
}
func DeleteUser(operationID string, user *base_info.ThirdMember) {
	members, err := wpsApi.Sdk.User.BatchDeleteCompanyMembers(operationID, []string{user.ThirdUnionId})
	if err != nil {
		log.Error("DeleteUser", "BatchDeleteCompanyMembers failed ", err.Error(), user)
		return
	}
	log.Info("DeleteUser", "BatchDeleteCompanyMembers success ", members)
}
func DisableUser(operationID string, user *base_info.ThirdMember) {
	req := api_req.BatchDisableThirdMembersReq{
		PlatformId: config.Config.WPS.PlatformId,
		UnionIds:   []string{user.ThirdUnionId},
	}
	members, err := wpsApi.Sdk.User.BatchDisableThirdMembers(operationID, req)
	if err != nil {
		log.Error("DisableUser", "BatchDisableThirdMembers failed ", err.Error(), req)
		return
	}
	log.Info("DisableUser", "BatchDisableThirdMembers success ", members)
}
func EnableUser(operationID string, user *base_info.ThirdMember) {
	req := api_req.BatchEnableThirdMembersReq{
		PlatformId: config.Config.WPS.PlatformId,
		UnionIds:   []string{user.ThirdUnionId},
	}
	members, err := wpsApi.Sdk.User.BatchEnableThirdMembers(operationID, req)
	if err != nil {
		log.Error("EnableUser", "BatchEnableThirdMembers failed ", err.Error(), req)
		return
	}
	log.Info("EnableUser", "BatchEnableThirdMembers success ", members)
}
func MoveUser(operationID string, user *base_info.ThirdMember) {
	wpsDept, err := wpsApi.Sdk.Dept.GetDeptInfoByThirdId(operationID, user.ThirdDeptId)
	if err != nil {
		return
	}
	wpsUser, err := wpsApi.Sdk.User.BatchGetCompanyMembersByThirdId(operationID, api_req.BatchGetCompanyMembersByThirdIdReq{PlatformId: config.Config.WPS.PlatformId, UnionIds: []string{user.ThirdUnionId}})
	if err != nil || len(wpsUser.Data.CompanyMembers) != 1 {
		log.Error("MoveUser", "BatchGetCompanyMembersByThirdId failed ", err.Error(), user, wpsUser)
		return
	}
	var newDeptIds []api_req.NewDept
	newDeptIds = append(newDeptIds, api_req.NewDept{DeptId: wpsDept.Data.Id})
	req := api_req.ChangeCompanyMembersDeptReq{
		AccountId:  wpsUser.Data.CompanyMembers[0].AccountId,
		DefDeptId:  wpsDept.Data.Id,
		NewDeptIds: newDeptIds,
		OldDeptIds: []string{wpsUser.Data.CompanyMembers[0].DefDeptId},
	}
	members, err := wpsApi.Sdk.User.ChangeCompanyMembersDept(operationID, req)
	if err != nil {
		log.Error("MoveUser", "ChangeCompanyMembersDept failed ", err.Error(), req)
		return
	}
	log.Info("MoveUser", "ChangeCompanyMembersDept success ", members)
}
func UpdateUser(operationID string, user *base_info.ThirdMember) {
	wpsUser, err := wpsApi.Sdk.User.BatchGetCompanyMembersByThirdId(operationID, api_req.BatchGetCompanyMembersByThirdIdReq{PlatformId: config.Config.WPS.PlatformId, UnionIds: []string{user.ThirdUnionId}})
	if err != nil || len(wpsUser.Data.CompanyMembers) != 1 {
		log.Error("UpdateUser", "BatchGetCompanyMembersByThirdId failed ", err.Error(), user, wpsUser)
		return
	}
	req := api_req.UpdateMemberInfoReq{
		Email:            "",
		EmployeeId:       "",
		EmploymentStatus: "",
		EmploymentType:   "",
		Gender:           "",
		MobilePhone:      "",
		NickName:         user.NickName,
		Telephone:        "",
		Title:            "",
		CustomFields:     nil,
		Avatar:           "",
	}
	wpsDept, err := wpsApi.Sdk.User.UpdateMemberInfo(operationID, wpsUser.Data.CompanyMembers[0].AccountId, req)
	if err != nil {
		log.Error("UpdateUser", "UpdateMemberInfo failed ", err.Error(), req)
		return
	}
	log.Info("UpdateUser", "UpdateMemberInfo success ", wpsDept)

}
