package dto


type SyncTaskWechatworkClientConfig struct{
    CorpID  string `json:"corpID"`
    Secret  string `json:"secret"`
    Token  string `json:"token,omitempty"`
    EncodingAESKey  string `json:"encodingAESKey,omitempty"`
    AgentUrl  string `json:"agentUrl,omitempty"`
}

