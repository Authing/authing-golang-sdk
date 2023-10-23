package dto


type GetFastpassClientAppsDto struct{
    QrcodeId string `json:"qrcodeId,omitempty"`
    AppId string `json:"appId,omitempty"`
}

