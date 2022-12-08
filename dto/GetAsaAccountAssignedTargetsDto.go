package dto


type GetAsaAccountAssignedTargetsDto struct{
    AppId string `json:"appId,omitempty"`
    AccountId string `json:"accountId,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

