package dto


type UpdateApplicationMfaSettingsDto struct{
    AppId  string `json:"appId"`
    EnabledFactors  []string `json:"enabledFactors,omitempty"`
    DisabledFactors  []string `json:"disabledFactors,omitempty"`
}

