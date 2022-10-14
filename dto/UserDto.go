package dto


type UserDto struct{
    UserId  string `json:"userId"`
    CreatedAt  string `json:"createdAt"`
    UpdatedAt  string `json:"updatedAt"`
    Status  string  `json:"status"`
    ExternalId  string `json:"externalId,omitempty"`
    Email  string `json:"email,omitempty"`
    Phone  string `json:"phone,omitempty"`
    PhoneCountryCode  string `json:"phoneCountryCode,omitempty"`
    Username  string `json:"username,omitempty"`
    Name  string `json:"name,omitempty"`
    Nickname  string `json:"nickname,omitempty"`
    Photo  string `json:"photo,omitempty"`
    LoginsCount  int `json:"loginsCount,omitempty"`
    LastLogin  string `json:"lastLogin,omitempty"`
    LastIp  string `json:"lastIp,omitempty"`
    Gender  string  `json:"gender"`
    EmailVerified  bool `json:"emailVerified"`
    PhoneVerified  bool `json:"phoneVerified"`
    PasswordLastSetAt  string `json:"passwordLastSetAt,omitempty"`
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
    UserSourceType  string  `json:"userSourceType"`
    UserSourceId  string `json:"userSourceId,omitempty"`
    LastLoginApp  string `json:"lastLoginApp,omitempty"`
    MainDepartmentId  string `json:"mainDepartmentId,omitempty"`
    LastMfaTime  string `json:"lastMfaTime,omitempty"`
    PasswordSecurityLevel  int `json:"passwordSecurityLevel,omitempty"`
    ResetPasswordOnNextLogin  bool `json:"resetPasswordOnNextLogin,omitempty"`
    DepartmentIds  []string `json:"departmentIds,omitempty"`
    Identities  []IdentityDto `json:"identities,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
    StatusChangedAt  string `json:"statusChangedAt,omitempty"`
}

