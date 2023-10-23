package dto


type PublicAccountSwitchLoginDto struct{
    TargetUserId  string `json:"targetUserId"`
    Options  SignInBySwitchOptionsDto `json:"options,omitempty"`
}

