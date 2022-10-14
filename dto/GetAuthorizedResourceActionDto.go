package dto


type GetAuthorizedResourceActionDto struct{
    Op  string  `json:"op"`
    List  []string `json:"list"`
}

