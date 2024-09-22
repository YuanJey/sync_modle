package sync

import (
	"github.com/YuanJey/sync_modle/internal/check"
	"github.com/YuanJey/sync_modle/internal/service"
	"github.com/YuanJey/sync_modle/pkg/base_info"
)

type UserSyncHandler struct {
	logic  check.UserLogic
	SyncCh chan *ChUserData
}

func (u *UserSyncHandler) SetLogic(logic check.UserLogic) {
	u.logic = logic
}
func NewUserSyncHandler() *UserSyncHandler {
	return &UserSyncHandler{SyncCh: make(chan *ChUserData)}
}

type ChUserData struct {
	operationID string
	user        *base_info.ThirdMember
}

func (u *UserSyncHandler) FullSyncUser(operationID string, users []*base_info.ThirdMember) {
	for i := range users {
		//u.process(operationID, users[i])
		u.SyncCh <- &ChUserData{operationID: operationID, user: users[i]}
	}
}
func (u *UserSyncHandler) process(operationID string, user *base_info.ThirdMember) {
	if u.logic.IsCreate(operationID, user) {
		service.CreateUser(operationID, user)
		return
	}
	if u.logic.IsDisable(operationID, user) {
		service.DisableUser(operationID, user)
	}
	if u.logic.IsEnable(operationID, user) {
		service.EnableUser(operationID, user)
	}
	if u.logic.IsMove(operationID, user) {
		service.MoveUser(operationID, user)
	}
	if u.logic.IsUpdate(operationID, user) {
		service.UpdateUser(operationID, user)
	}
	if u.logic.IsDelete(operationID, user) {
		service.DeleteUser(operationID, user)
	}
}
func (u *UserSyncHandler) Work() {
	for {
		select {
		case data := <-u.SyncCh:
			u.process(data.operationID, data.user)
		}
	}
}
