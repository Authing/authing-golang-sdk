package dto


type UpdateUserReqDto struct{
    UserId  string `json:"userId"`
    PhoneCountryCode  string `json:"phoneCountryCode,omitempty"`
    Name  string `json:"name,omitempty"`
    Nickname  string `json:"nickname,omitempty"`
    Photo  string `json:"photo,omitempty"`
    ExternalId  string `json:"externalId,omitempty"`
    Status  string  `json:"status,omitempty"`
    EmailVerified  bool `json:"emailVerified,omitempty"`
    PhoneVerified  bool `json:"phoneVerified,omitempty"`
    Birthdate  string `json:"birthdate,omitempty"`
    Country  string `json:"country,omitempty"`
    Province  string `json:"province,omitempty"`
    City  string `json:"city,omitempty"`
    Address  string `json:"address,omitempty"`
    StreetAddress  string `json:"streetAddress,omitempty"`
    PostalCode  string `json:"postalCode,omitempty"`
    Gender  string  `json:"gender,omitempty"`
    Username  string `json:"username,omitempty"`
    PasswordEncryptType  string  `json:"passwordEncryptType,omitempty"`
    Email  string `json:"email,omitempty"`
    Phone  string `json:"phone,omitempty"`
    Password  string `json:"password,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
}

