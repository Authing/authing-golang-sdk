package dto


type WebhookDto struct{
    WebhookId  string `json:"webhookId"`
    CreatedAt  string `json:"createdAt"`
    UpdatedAt  string `json:"updatedAt"`
    Name  string `json:"name"`
    Url  string `json:"url"`
    ContentType  string  `json:"contentType"`
    Enabled  bool `json:"enabled"`
    Events  []string `json:"events,omitempty"`
    Secret  string `json:"secret,omitempty"`
}

