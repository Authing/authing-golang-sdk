package dto


type GetGroupDto struct{
    Code string `json:"code,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
}

