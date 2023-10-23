package dto


type UpdatePublicAccountBatchReqDto struct{
    List  []UpdatePublicAccountInfoDto `json:"list"`
    Options  UpdatePublicAccountBatchOptionsDto `json:"options,omitempty"`
}

