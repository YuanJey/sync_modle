package model

import (
	"github.com/YuanJey/sync_modle/pkg/db/mysql"
	"github.com/YuanJey/sync_modle/pkg/db/mysql/structs"
	"github.com/YuanJey/sync_modle/pkg/utils"
)

const membersTableName = "members"

func BatchInsertMember(userList []*structs.Member) error {
	if userList == nil {
		return nil
	}
	return utils.Wrap(mysql.DbMysql.Db.Table(membersTableName).Create(userList).Error, "BatchInsertUser failed")
}
