package dto


type TriggerWebhookDto struct{
    WebhookId  string `json:"webhookId"`
    RequestHeaders  interface{} `json:"requestHeaders,omitempty"`
    RequestBody  interface{} `json:"requestBody,omitempty"`
}

