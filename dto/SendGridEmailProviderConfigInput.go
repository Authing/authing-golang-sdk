package dto


type SendGridEmailProviderConfigInput struct{
    Sender  string `json:"sender"`
    Apikey  string `json:"apikey"`
}

