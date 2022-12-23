package dto


type CreateResourceDto struct{
    Type  string  `json:"type"`
    Code  string `json:"code"`
    Description  string `json:"description,omitempty"`
    Name  string `json:"name,omitempty"`
    Actions  []ResourceAction `json:"actions,omitempty"`
    ApiIdentifier  string `json:"apiIdentifier,omitempty"`
    Namespace  string `json:"namespace,omitempty"`
}

