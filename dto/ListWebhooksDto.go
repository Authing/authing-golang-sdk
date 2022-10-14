package dto


type ListWebhooksDto struct{
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

