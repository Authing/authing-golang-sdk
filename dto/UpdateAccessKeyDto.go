package dto


type UpdateAccessKeyDto struct{
    Enable  bool `json:"enable"`
    AccessKeyId  string `json:"accessKeyId"`
}

