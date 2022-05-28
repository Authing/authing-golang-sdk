package dto


type AuthorizeResourcesDto struct{
    List  []AuthorizeResourceItem `json:"list"`
    Namespace  string `json:"namespace,omitempty"`
}

