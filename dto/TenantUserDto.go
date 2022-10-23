package dto


type TenantUserDto struct{
    TenantId  string `json:"tenantId"`
    UserPoolId  string `json:"userPoolId"`
    Username  string `json:"username"`
    Name  string `json:"name"`
    Nickname  string `json:"nickname"`
    Email  string `json:"email"`
    Phone  string `json:"phone"`
    Address  string `json:"address"`
    Birthdate  string `json:"birthdate"`
    Blocked  bool `json:"blocked"`
    IsTenantAdmin  bool `json:"isTenantAdmin"`
    LastIP  string `json:"lastIP"`
    LastLoginApp  string `json:"lastLoginApp"`
    LastLoginAppLogo  string `json:"lastLoginAppLogo"`
    LastLoginAppName  string `json:"lastLoginAppName"`
    LoginsCount  int `json:"loginsCount"`
    MemberId  string `json:"memberId"`
    LinkUserId  string `json:"linkUserId"`
}

