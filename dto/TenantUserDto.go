package dto


type TenantUserDto struct{
    Email  string `json:"email,omitempty"`
    Phone  string `json:"phone,omitempty"`
    PhoneCountryCode  string `json:"phoneCountryCode,omitempty"`
    Username  string `json:"username,omitempty"`
    Name  string `json:"name,omitempty"`
    Nickname  string `json:"nickname,omitempty"`
    Photo  string `json:"photo,omitempty"`
    LoginsCount  int `json:"loginsCount,omitempty"`
    LastIp  string `json:"lastIp,omitempty"`
    Gender  string  `json:"gender"`
    Birthdate  string `json:"birthdate,omitempty"`
    Country  string `json:"country,omitempty"`
    Province  string `json:"province,omitempty"`
    City  string `json:"city,omitempty"`
    Address  string `json:"address,omitempty"`
    StreetAddress  string `json:"streetAddress,omitempty"`
    PostalCode  string `json:"postalCode,omitempty"`
    GivenName  string `json:"givenName,omitempty"`
    FamilyName  string `json:"familyName,omitempty"`
    MiddleName  string `json:"middleName,omitempty"`
    PreferredUsername  string `json:"preferredUsername,omitempty"`
    LastLoginApp  string `json:"lastLoginApp,omitempty"`
    UserPoolId  string `json:"userPoolId"`
    TenantId  string `json:"tenantId"`
    MemberId  string `json:"memberId"`
    LinkUserId  string `json:"linkUserId"`
    IsTenantAdmin  bool `json:"isTenantAdmin"`
    Password  string `json:"password,omitempty"`
    Salt  string `json:"salt,omitempty"`
}

