package dto


type UpdateResourceDto struct{
    Code  string `json:"code"`
    Description  string `json:"description,omitempty"`
    Name  string `json:"name,omitempty"`
    Actions  []ResourceAction `json:"actions,omitempty"`
    ApiIdentifier  string `json:"apiIdentifier,omitempty"`
    Namespace  string `json:"namespace,omitempty"`
    Type  string  `json:"type,omitempty"`
}

