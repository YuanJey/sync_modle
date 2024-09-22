package base_info

type ThirdDept struct {
	ThirdUnionId string `json:"third_union_id"`
	Name         string `json:"name"`
	ParentId     string `json:"parent_id"`
	Weight       int    `json:"weight"`
	Synced       bool   `json:"synced"`
	Children     []*ThirdDept
}
type ThirdMember struct {
	ThirdUnionId string `json:"third_union_id"`
	NickName     string `json:"nick_name"`
	ThirdDeptId  string `json:"third_dept_id"`
	Status       string `json:"status"`
	Gender       string `json:"gender"`
	MobilePhone  string `json:"mobile_phone"`
}
