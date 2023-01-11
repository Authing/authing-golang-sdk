package dto

type AuthorizeApplicationAccessDto struct {
	AppId string                            `json:"appId"`
	List  []ApplicationPermissionRecordItem `json:"list"`
}
