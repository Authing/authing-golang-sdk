package dto


type SignupProfileDto struct{
    Nickname  string `json:"nickname,omitempty"`
    Company  string `json:"company,omitempty"`
    Photo  string `json:"photo,omitempty"`
    Device  string `json:"device,omitempty"`
    Browser  string `json:"browser,omitempty"`
    Name  string `json:"name,omitempty"`
    GivenName  string `json:"givenName,omitempty"`
    FamilyName  string `json:"familyName,omitempty"`
    MiddleName  string `json:"middleName,omitempty"`
    Profile  string `json:"profile,omitempty"`
    PreferredUsername  string `json:"preferredUsername,omitempty"`
    Website  string `json:"website,omitempty"`
    Gender  string  `json:"gender,omitempty"`
    Birthdate  string `json:"birthdate,omitempty"`
    Zoneinfo  string `json:"zoneinfo,omitempty"`
    Locale  string `json:"locale,omitempty"`
    Address  string `json:"address,omitempty"`
    Formatted  string `json:"formatted,omitempty"`
    StreetAddress  string `json:"streetAddress,omitempty"`
    Locality  string `json:"locality,omitempty"`
    Region  string `json:"region,omitempty"`
    PostalCode  string `json:"postalCode,omitempty"`
    Country  string `json:"country,omitempty"`
    Email  string `json:"email,omitempty"`
    Phone  string `json:"phone,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
}

