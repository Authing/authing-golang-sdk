package dto


type UpdateUserBatchReqDto struct{
    List  []UpdateUserInfoDto `json:"list"`
    Options  UpdateUserBatchOptionsDto `json:"options,omitempty"`
}

