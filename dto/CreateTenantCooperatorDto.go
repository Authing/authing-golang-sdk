package dto


type CreateTenantCooperatorDto struct{
    Policies  []string `json:"policies"`
    UserId  string `json:"userId"`
    ApiAuthorized  bool `json:"apiAuthorized,omitempty"`
    SendPhoneNotification  bool `json:"sendPhoneNotification,omitempty"`
    SendEmailNotification  bool `json:"sendEmailNotification,omitempty"`
}

