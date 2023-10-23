package dto


type TreeStructs struct{
    Code  string `json:"code"`
    Name  string `json:"name"`
    Value  string `json:"value,omitempty"`
    Actions  []string `json:"actions,omitempty"`
    Children  []TreeStructs `json:"children,omitempty"`
}

