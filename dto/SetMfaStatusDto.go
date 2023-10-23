package dto


type SetMfaStatusDto struct{
    MfaTriggerData  GetMfaInfoDataDto `json:"mfaTriggerData"`
    UserId  string `json:"userId"`
    UserIdType  string  `json:"userIdType,omitempty"`
}

