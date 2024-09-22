package structs

type Organization struct {
	ID       string `json:"id" gorm:"column:id;primary_key"`
	Name     string `json:"name" gorm:"column:name"`
	ParentId string `json:"parent_id" gorm:"column:parent_id"`
	Weight   int    `json:"weight" gorm:"column:weight"`
	IsRoot   int    `json:"isRoot" gorm:"column:is_root"`
}
type User struct {
	ID    string `json:"id" gorm:"column:id;primary_key"`
	Name  string `json:"name" gorm:"column:name"`
	OrgID string `json:"orgId" gorm:"column:org_id"`
}
type Dept struct {
	Id              string `json:"id" gorm:"id"`
	CompanyId       string `json:"company_id" gorm:"company_id"`
	Weight          int    `json:"weight" gorm:"weight"`
	ThirdUnionId    string `json:"third_union_id" gorm:"third_union_id"`
	ThirdPlatformId string `json:"third_platform_id" gorm:"third_platform_id"`
	Ctime           int    `json:"ctime" gorm:"ctime"`
	Mtime           int    `json:"mtime" gorm:"mtime"`
	Type            string `json:"type" gorm:"type"`
	ParentId        string `json:"parent_id" gorm:"parent_id"`
	Name            string `json:"name" gorm:"name"`
	Alias           string `json:"alias" gorm:"alias"`
	LeaderId        string `json:"leader_id" gorm:"leader_id"`
	CreatorId       string `json:"creator_id" gorm:"creator_id"`
	AbsPath         string `json:"abs_path" gorm:"abs_path"`
	IdPath          string `json:"id_path" gorm:"id_path"`
	Source          string `json:"source" gorm:"source"`
	Synced          bool   `json:"synced" gorm:"synced"`
}
type Member struct {
	AccountId         string `json:"account_id" gorm:"account_id"`
	Address           string `json:"address" gorm:"address"`
	CompanyId         string `json:"company_id" gorm:"company_id"`
	CompanyUid        string `json:"company_uid" gorm:"company_uid"`
	Ctime             int    `json:"ctime" gorm:"ctime"`
	DefDeptId         string `json:"def_dept_id" gorm:"def_dept_id"`
	DeptNum           int    `json:"dept_num" gorm:"dept_num"`
	Email             string `json:"email" gorm:"email"`
	EmployeeId        string `json:"employee_id" gorm:"employee_id"`
	EmployeeType      string `json:"employee_type" gorm:"employee_type"`
	EmploymentStatus  string `json:"employment_status" gorm:"employment_status"`
	Gender            string `json:"gender" gorm:"gender"`
	IsInnerSuperAdmin bool   `json:"is_inner_super_admin" gorm:"is_inner_super_admin"`
	Leader            string `json:"leader" gorm:"leader"`
	LoginName         string `json:"login_name" gorm:"login_name"`
	MobilePhone       string `json:"mobile_phone" gorm:"mobile_phone"`
	Mtime             int    `json:"mtime" gorm:"mtime"`
	NickName          string `json:"nick_name" gorm:"nick_name"`
	PreStatus         int    `json:"pre_status" gorm:"pre_status"`
	Role              string `json:"role" gorm:"role"`
	Source            string `json:"source" gorm:"source"`
	Status            string `json:"status" gorm:"status"`
	Telephone         string `json:"telephone" gorm:"telephone"`
	ThirdPlatformId   string `json:"third_platform_id" gorm:"third_platform_id"`
	ThirdUnionId      string `json:"third_union_id" gorm:"third_union_id"`
	Title             string `json:"title" gorm:"title"`
	WorkPlace         string `json:"work_place" gorm:"work_place"`
	//CustomFields      []CustomField `json:"custom_fields" gorm:"custom_fields"`
	//DeptList          []Dept        `json:"depts" gorm:"depts"`
	Synced             bool   `json:"synced"`
	Operation          string `json:"operation" gorm:"operation"`
	MemberUpdateFields []string
}
type CustomField struct {
	FieldId string `json:"field_id" gorm:"field_id"`
	Text    string `json:"text" gorm:"text"`
}
