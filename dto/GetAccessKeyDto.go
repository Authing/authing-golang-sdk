package dto


type GetAccessKeyDto struct{
    UserId string `json:"userId,omitempty"`
    AccessKeyId string `json:"accessKeyId,omitempty"`
}

