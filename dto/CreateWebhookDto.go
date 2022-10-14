package dto


type CreateWebhookDto struct{
    ContentType  string  `json:"contentType"`
    Events  []string `json:"events"`
    Url  string `json:"url"`
    Name  string `json:"name"`
    Enabled  bool `json:"enabled,omitempty"`
    Secret  string `json:"secret,omitempty"`
}

