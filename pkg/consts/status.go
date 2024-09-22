package consts

const (
	ThirdUserStatusEnable  = "0"
	ThirdUserStatusDisable = "1"
	ThirdUserStatusDelete  = "2"

	WpsUserStatusDisabled  = "disabled"
	WpsUserStatusActive    = "active"
	WpsUserStatusNotActive = "notactive"
	WpsUserStatusDimission = "dimission"
)

type thirdUserStatus interface {
	Enable() string
	Disable() string
	Delete() string
}
type SyncUserStatus struct{}

func (s *SyncUserStatus) Enable() string {
	return ThirdUserStatusEnable
}

func (s *SyncUserStatus) Disable() string {
	return ThirdUserStatusDisable
}

func (s *SyncUserStatus) Delete() string {
	return ThirdUserStatusDelete
}

var ThirdUserStatus = &SyncUserStatus{}
