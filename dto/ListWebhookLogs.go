package dto


type ListWebhookLogs struct{
    WebhookId  string `json:"webhookId"`
    Page  int `json:"page,omitempty"`
    Limit  int `json:"limit,omitempty"`
}

