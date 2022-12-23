package dto


type NamespacesListRespDto struct{
    Name  string `json:"name"`
    Code  string `json:"code"`
    Description  string `json:"description,omitempty"`
    Status  int `json:"status,omitempty"`
}

