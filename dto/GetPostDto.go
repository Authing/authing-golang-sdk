package dto


type GetPostDto struct{
    Code string `json:"code,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
}

