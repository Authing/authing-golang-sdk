package dto


type CreateMultipleTenantAdminDto struct{
    TenantIds  []string `json:"tenantIds"`
    UserId  string `json:"userId"`
    ApiAuthorized  bool `json:"apiAuthorized,omitempty"`
    SendPhoneNotification  bool `json:"sendPhoneNotification,omitempty"`
    SendEmailNotification  bool `json:"sendEmailNotification,omitempty"`
}

