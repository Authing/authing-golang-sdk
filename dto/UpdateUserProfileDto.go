package dto


type UpdateUserProfileDto struct{
    Name  string `json:"name,omitempty"`
    Nickname  string `json:"nickname,omitempty"`
    Photo  string `json:"photo,omitempty"`
    ExternalId  string `json:"externalId,omitempty"`
    Birthdate  string `json:"birthdate,omitempty"`
    Country  string `json:"country,omitempty"`
    Province  string `json:"province,omitempty"`
    City  string `json:"city,omitempty"`
    Address  string `json:"address,omitempty"`
    StreetAddress  string `json:"streetAddress,omitempty"`
    PostalCode  string `json:"postalCode,omitempty"`
    Gender  string  `json:"gender,omitempty"`
    Username  string `json:"username,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
}

