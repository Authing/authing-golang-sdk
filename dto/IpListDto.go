package dto


type IpListDto struct{
    IpType string `json:"ipType,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

