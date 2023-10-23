package dto


type UpdateTenantCooperatorDto struct{
    ApiAuthorized  bool `json:"apiAuthorized"`
    Policies  []string `json:"policies"`
}

