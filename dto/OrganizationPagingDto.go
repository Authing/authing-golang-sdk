package dto


type OrganizationPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []OrganizationDto `json:"list"`
}

