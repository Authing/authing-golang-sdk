package dto


type UpdateWebhookDto struct{
    WebhookId  string `json:"webhookId"`
    Name  string `json:"name,omitempty"`
    Url  string `json:"url,omitempty"`
    Events  []string `json:"events,omitempty"`
    ContentType  string  `json:"contentType,omitempty"`
    Enabled  bool `json:"enabled,omitempty"`
    Secret  string `json:"secret,omitempty"`
}

