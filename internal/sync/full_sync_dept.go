package sync

import (
	"github.com/YuanJey/sync_modle/internal/check"
	"github.com/YuanJey/sync_modle/internal/service"
	"github.com/YuanJey/sync_modle/pkg/db/mysql"
	"github.com/YuanJey/sync_modle/pkg/db/mysql/structs"
	"github.com/YuanJey/sync_modle/pkg/log"
	wpsApi "github.com/YuanJey/wps-api"
	"github.com/YuanJey/wps-api/dept"
)

type DeptSyncHandler struct {
	logic  check.DeptLogic
	SyncCh chan *ChDeptData
}

func (d *DeptSyncHandler) SetLogic(logic check.DeptLogic) {
	d.logic = logic
}
func NewDeptSyncHandler() *DeptSyncHandler {
	return &DeptSyncHandler{SyncCh: make(chan *ChDeptData)}
}

type ChDeptData struct {
	operationID string
	org         *structs.Organization
	wpsAllDept  *dept.AllWpsDept
}

func (d *DeptSyncHandler) FullSyncDept(operationID string) {
	log.Info(operationID, "FullSyncDept start")
	rootDept, err := mysql.GetRootDept()
	if err != nil {
		log.Error(operationID, "GetRootDept err : ", err)
		return
	}
	wpsDept, err := wpsApi.Sdk.Dept.GetAllWpsDept(operationID)
	if err != nil {
		log.Error(operationID, "GetAllWpsDept err : ", err)
		return
	}
	d.SyncSubTempDept(operationID, rootDept.ID, wpsDept)
}
func (d *DeptSyncHandler) SyncSubTempDept(operationID string, parentId string, wpsAllDept *dept.AllWpsDept) {
	list, err := mysql.GetSubOrganizationList(parentId)
	if err != nil {
		log.Error(operationID, "SyncSubTempDept err : ", err)
		return
	}
	if len(list) > 0 {
		for i := range list {
			//d.process(operationID, &list[i], wpsAllDept)
			d.SyncCh <- &ChDeptData{
				operationID: operationID,
				org:         &list[i],
				wpsAllDept:  wpsAllDept,
			}
		}
		for _, syncDept := range list {
			d.SyncSubTempDept(operationID, syncDept.ID, wpsAllDept)
		}
	}
	for thirdId := range wpsAllDept.TidDeptList {
		if d.logic.IsDelete(operationID, wpsAllDept.TidDeptList[thirdId]) {
			service.DeleteDept(operationID, thirdId)
		}
	}
}
func (d *DeptSyncHandler) process(operationID string, org *structs.Organization, wpsAllDept *dept.AllWpsDept) {
	if d.logic.IsCreate(operationID, org) {
		service.CreateDept(operationID, org)
	}
	if d.logic.IsUpdate(operationID, org) {
		if wpsDept, ok := wpsAllDept.TidDeptList[org.ID]; ok {
			wpsDept.Synced = true
			service.UpdateDept(operationID, org)
		}
	}
	if d.logic.IsMove(operationID, org) {
		if wpsDept, ok := wpsAllDept.TidDeptList[org.ID]; ok {
			wpsDept.Synced = true
			service.MoveDept(operationID, org)
		}
	}
}
func (d *DeptSyncHandler) Work() {
	for {
		select {
		case data := <-d.SyncCh:
			d.process(data.operationID, data.org, data.wpsAllDept)
		default:
		}
	}
}
