package dto


type DataStatementPermissionDto struct{
    Effect  string  `json:"effect"`
    Permissions  []string `json:"permissions"`
}

