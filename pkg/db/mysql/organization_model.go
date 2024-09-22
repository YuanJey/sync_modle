package mysql

import (
	"github.com/YuanJey/sync_modle/pkg/db/mysql/structs"
	"github.com/YuanJey/sync_modle/pkg/log"
	"github.com/YuanJey/sync_modle/pkg/utils"
)

const DeptTableName = "organizations"

func GetRootDept() (structs.Organization, error) {
	var dept structs.Organization
	err := DbMysql.Db.Table(DeptTableName).Where("is_root = ?", 1).First(&dept).Error
	return dept, err
}
func BatchInsertOrganization(deptList []*structs.Organization) error {
	if deptList == nil {
		return nil
	}
	return utils.Wrap(DbMysql.Db.Table(DeptTableName).Create(deptList).Error, "BatchInsertOrganization failed")
}
func GetSubOrganizationList(superCode string) ([]structs.Organization, error) {
	var deptList []structs.Organization
	err := DbMysql.Db.Table(DeptTableName).Where("parent_id = ?", superCode).Find(&deptList).Error
	return deptList, err
}
func GetOrganizationByCode(code string) (structs.Organization, error) {
	var dept structs.Organization
	err := DbMysql.Db.Table(DeptTableName).Where("id = ?", code).First(&dept).Error
	return dept, err
}
func InsertOrUpdateOrganization(deptList []*structs.Organization) error {
	if len(deptList) == 0 || deptList == nil {
		err := DeleteAllOrganization()
		if err != nil {
			return err
		}
		return nil
	}
	var ids []string
	for i := range deptList {
		ids = append(ids, deptList[i].ID)
	}
	err := DbMysql.Db.Table(DeptTableName).Where("id not in ?", ids).Delete(&structs.Organization{}).Error
	if err != nil {
		log.Error("InsertOrUpdateOrganization Delete err :", err.Error())
	}
	return utils.Wrap(DbMysql.Db.Table(DeptTableName).Save(deptList).Error, "InsertOrUpdateOrganization failed")
}

func DeleteAllOrganization() error {
	return utils.Wrap(DbMysql.Db.Table(DeptTableName).Where("1 = 1").Delete(&structs.Organization{}).Error, "DeleteAllOrganization failed")
}
