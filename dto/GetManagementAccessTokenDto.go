package dto

type GetManagementAccessTokenDto struct {
	AccessKeySecret string `json:"accessKeySecret"`
	AccessKeyId     string `json:"accessKeyId"`
}
