package dto


type DataResourcePolicyTreeStructs struct{
    Code  string `json:"code"`
    Value  string `json:"value,omitempty"`
    Name  string `json:"name"`
    Action  string `json:"action"`
    Enabled  bool `json:"enabled"`
    Children  []string `json:"children,omitempty"`
}

