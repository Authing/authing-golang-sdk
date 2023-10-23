package dto


type MultipleTenantAdminDto struct{
    UserId  string `json:"userId"`
    Name  string `json:"name"`
    Phone  string `json:"phone"`
    Email  string `json:"email"`
    ApiAuthorized  bool `json:"apiAuthorized,omitempty"`
    LastLogin  string `json:"lastLogin"`
}

