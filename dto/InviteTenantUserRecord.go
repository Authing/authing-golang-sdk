package dto


type InviteTenantUserRecord struct{
    RecordId  int `json:"recordId"`
    InviteAccount  string `json:"inviteAccount"`
    VerifiedStatus  string `json:"verifiedStatus"`
    InviteLink  string `json:"inviteLink"`
    CreatedAt  string `json:"createdAt"`
    ActivatedAt  string `json:"activatedAt"`
}

