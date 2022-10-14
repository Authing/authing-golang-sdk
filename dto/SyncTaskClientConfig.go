package dto


type SyncTaskClientConfig struct{
    LarkConfig  SyncTaskLarkClientConfig `json:"larkConfig,omitempty"`
    LarkInternationalConfig  SyncTaskLarkClientConfig `json:"larkInternationalConfig,omitempty"`
    WechatworkConfig  SyncTaskWechatworkClientConfig `json:"wechatworkConfig,omitempty"`
    DingtalkConfig  SyncTaskDingtalkClientConfig `json:"dingtalkConfig,omitempty"`
    MokaConfig  SyncTaskMokaClientConfig `json:"mokaConfig,omitempty"`
    ScimConfig  SyncTaskScimClientConfig `json:"scimConfig,omitempty"`
    ActiveDirectoryConfig  SyncTaskActiveDirectoryClientConfig `json:"activeDirectoryConfig,omitempty"`
    LdapConfig  SyncTaskLdapClientConfig `json:"ldapConfig,omitempty"`
    ItalentConfig  SyncTaskItalentClientConfig `json:"italentConfig,omitempty"`
    MaycurConfig  SyncTaskMaycurClientConfig `json:"maycurConfig,omitempty"`
    FxiaokeConfig  SyncTaskFxiaokeClientConfig `json:"fxiaokeConfig,omitempty"`
    XiaoshouyiConfig  SyncTaskXiaoshouyiClientConfig `json:"xiaoshouyiConfig,omitempty"`
    KayangConfig  SyncTaskKayangClientConfig `json:"kayangConfig,omitempty"`
}

