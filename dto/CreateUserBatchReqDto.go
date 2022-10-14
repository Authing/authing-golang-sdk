package dto


type CreateUserBatchReqDto struct{
    List  []CreateUserInfoDto `json:"list"`
    Options  CreateUserOptionsDto `json:"options,omitempty"`
}

