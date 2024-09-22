package sync

var DeptSync *DeptSyncHandler
var UserSync *UserSyncHandler

func init() {
	DeptSync = NewDeptSyncHandler()
	UserSync = NewUserSyncHandler()
}
