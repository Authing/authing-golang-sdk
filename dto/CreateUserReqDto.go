package dto


type CreateUserReqDto struct{
    Status  string  `json:"status,omitempty"`
    Email  string `json:"email,omitempty"`
    Phone  string `json:"phone,omitempty"`
    PhoneCountryCode  string `json:"phoneCountryCode,omitempty"`
    Username  string `json:"username,omitempty"`
    ExternalId  string `json:"externalId,omitempty"`
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
    Password  string `json:"password,omitempty"`
    Salt  string `json:"salt,omitempty"`
    TenantIds  []string `json:"tenantIds,omitempty"`
    Otp  CreateUserOtpDto `json:"otp,omitempty"`
    DepartmentIds  []string `json:"departmentIds,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
    MetadataSource  interface{} `json:"metadataSource,omitempty"`
    Identities  []CreateIdentityDto `json:"identities,omitempty"`
    IdentityNumber  string `json:"identityNumber,omitempty"`
    Options  CreateUserOptionsDto `json:"options,omitempty"`
}

