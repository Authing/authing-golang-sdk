package dto


type CreateUserInfoDto struct{
    Status  string  `json:"status,omitempty"`
    Email  string `json:"email,omitempty"`
    PasswordEncryptType  string  `json:"passwordEncryptType,omitempty"`
    Phone  string `json:"phone,omitempty"`
    PhoneCountryCode  string `json:"phoneCountryCode,omitempty"`
    Username  string `json:"username,omitempty"`
    Name  string `json:"name,omitempty"`
    Nickname  string `json:"nickname,omitempty"`
    Photo  string `json:"photo,omitempty"`
    Gender  string  `json:"gender,omitempty"`
    EmailVerified  bool `json:"emailVerified,omitempty"`
    PhoneVerified  bool `json:"phoneVerified,omitempty"`
    Birthdate  string `json:"birthdate,omitempty"`
    Country  string `json:"country,omitempty"`
    Province  string `json:"province,omitempty"`
    City  string `json:"city,omitempty"`
    Address  string `json:"address,omitempty"`
    StreetAddress  string `json:"streetAddress,omitempty"`
    PostalCode  string `json:"postalCode,omitempty"`
    ExternalId  string `json:"externalId,omitempty"`
    DepartmentIds  []string `json:"departmentIds,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
    Password  string `json:"password,omitempty"`
    TenantIds  []string `json:"tenantIds,omitempty"`
    Identities  []CreateIdentityDto `json:"identities,omitempty"`
}

