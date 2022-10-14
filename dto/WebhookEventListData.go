package dto


type WebhookEventListData struct{
    Categories  []WebhookCategoryDto `json:"categories"`
    Events  []WebhookEventDto `json:"events"`
}

