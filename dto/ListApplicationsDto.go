package dto


type ListApplicationsDto struct{
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    IsIntegrateApp bool `json:"isIntegrateApp,omitempty"`
    IsSelfBuiltApp bool `json:"isSelfBuiltApp,omitempty"`
    SsoEnabled bool `json:"ssoEnabled,omitempty"`
    Keywords string `json:"keywords,omitempty"`
}

