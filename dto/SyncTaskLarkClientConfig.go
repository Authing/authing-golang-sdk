package dto


type SyncTaskLarkClientConfig struct{
    AppId  string `json:"app_id"`
    AppSecret  string `json:"app_secret"`
    EncryptKey  string `json:"encrypt_key,omitempty"`
    VerificationToken  string `json:"verification_token,omitempty"`
}

