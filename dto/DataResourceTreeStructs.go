package dto


type DataResourceTreeStructs struct{
    Code  string `json:"code"`
    Name  string `json:"name"`
    Value  string `json:"value,omitempty"`
    Children  []interface{} `json:"children,omitempty"`
}

