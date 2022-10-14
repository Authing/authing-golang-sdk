package dto


type GetWebhooksData struct{
    TotalCount  int `json:"totalCount"`
    List  []WebhookDto `json:"list"`
}

