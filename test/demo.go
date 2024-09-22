package main

import (
	"github.com/YuanJey/sync_modle/internal/check"
	"github.com/YuanJey/sync_modle/internal/sync"
)

func main() {
	deptCheck := check.NewDefaultDeptCheck()
	sync.DeptSync.SetLogic(deptCheck)
	userCheck := check.NewDefaultUserCheck()
	sync.UserSync.SetLogic(userCheck)
	sync.DeptSync.FullSyncDept("test demo")
	go sync.DeptSync.Work()
	sync.UserSync.FullSyncUser("test demo", nil)
	go sync.UserSync.Work()
}
