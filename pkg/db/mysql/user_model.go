package mysql

import (
	"github.com/YuanJey/sync_modle/pkg/db/mysql/structs"
	"github.com/YuanJey/sync_modle/pkg/log"
	"github.com/YuanJey/sync_modle/pkg/utils"
)

const UserTableName = "users"

func BatchInsertUser(userList []*structs.User) error {
	if userList == nil {
		return nil
	}
	return utils.Wrap(DbMysql.Db.Table(UserTableName).Create(userList).Error, "BatchInsertUser failed")
}
func GetDeptUserList(orgCode string) ([]structs.User, error) {
	var userList []structs.User
	err := DbMysql.Db.Table(UserTableName).Where("org_id = ?", orgCode).Find(&userList).Error
	return userList, err
}
func GetUserByIdentityCard(identityCard string) (structs.User, error) {
	var user structs.User
	err := DbMysql.Db.Table(UserTableName).Where("id = ?", identityCard).First(&user).Error
	return user, err
}
func InsertOrUpdateUser(userList []*structs.User) error {
	if len(userList) == 0 || userList == nil {
		err := DeleteAllUser()
		if err != nil {
			return err
		}
		return nil
	}
	var ids []string
	for i := range userList {
		ids = append(ids, userList[i].ID)
	}
	err := DbMysql.Db.Table(UserTableName).Where("id not in ?", ids).Delete(&structs.User{}).Error
	if err != nil {
		log.Error("InsertOrUpdateUser Delete err :", err.Error())
	}
	return utils.Wrap(DbMysql.Db.Table(UserTableName).Save(userList).Error, "InsertOrUpdateUser failed")
}
func DeleteAllUser() error {
	return utils.Wrap(DbMysql.Db.Table(UserTableName).Where("1 = 1").Delete(&structs.User{}).Error, "DeleteAllUser failed")
}
