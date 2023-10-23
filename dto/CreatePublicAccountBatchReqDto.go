package dto


type CreatePublicAccountBatchReqDto struct{
    List  []CreatePublicAccountReqDto `json:"list"`
    Options  CreatePublicAccountOptionsDto `json:"options,omitempty"`
}

