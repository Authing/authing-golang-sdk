package dto


type ListWebhookLogsData struct{
    TotalCount  int `json:"totalCount"`
    List  []WebhookLogDto `json:"list"`
}

