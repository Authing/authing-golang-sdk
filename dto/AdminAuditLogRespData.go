package dto


type AdminAuditLogRespData struct{
    TotalCount  int `json:"totalCount"`
    List  []AdminAuditLogDto `json:"list"`
}

