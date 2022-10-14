package dto


type SyncTaskDingtalkClientConfig struct{
    CorpId  string `json:"corpId"`
    AppKey  string `json:"appKey"`
    AppSecret  string `json:"appSecret"`
    AesKey  string `json:"aes_key,omitempty"`
    Token  string `json:"token,omitempty"`
}

