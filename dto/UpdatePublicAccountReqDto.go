package dto


type UpdatePublicAccountReqDto struct{
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
    Email  string `json:"email,omitempty"`
    Phone  string `json:"phone,omitempty"`
    Password  string `json:"password,omitempty"`
    Company  string `json:"company,omitempty"`
    Browser  string `json:"browser,omitempty"`
    Device  string `json:"device,omitempty"`
    GivenName  string `json:"givenName,omitempty"`
    FamilyName  string `json:"familyName,omitempty"`
    MiddleName  string `json:"middleName,omitempty"`
    Profile  string `json:"profile,omitempty"`
    PreferredUsername  string `json:"preferredUsername,omitempty"`
    Website  string `json:"website,omitempty"`
    Zoneinfo  string `json:"zoneinfo,omitempty"`
    Locale  string `json:"locale,omitempty"`
    Formatted  string `json:"formatted,omitempty"`
    Region  string `json:"region,omitempty"`
    IdentityNumber  string `json:"identityNumber,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
    Options  UpdatePublicAccountOptionsDto `json:"options,omitempty"`
}

