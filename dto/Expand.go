package dto


type Expand struct{
    Field  string `json:"field"`
    Select  []string `json:"select"`
}

