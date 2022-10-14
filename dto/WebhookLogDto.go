package dto


type WebhookLogDto struct{
    WebhookId  string `json:"webhookId"`
    EventName  string `json:"eventName"`
    RequestBody  interface{} `json:"requestBody"`
    RequestHeaders  interface{} `json:"requestHeaders"`
    ResponseStatusCode  int `json:"responseStatusCode"`
    ResponseHeaders  interface{} `json:"responseHeaders"`
    ResponseBody  interface{} `json:"responseBody"`
    Timestamp  string `json:"timestamp"`
    Success  bool `json:"success"`
    ErrorMessage  string `json:"errorMessage,omitempty"`
}

