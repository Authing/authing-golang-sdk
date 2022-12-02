package dto


type NamespaceListPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []NamespacesListRespDto `json:"list"`
}

