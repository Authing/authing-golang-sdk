package dto


type GetUserPrincipalAuthenticationInfoDto struct{
    UserId string `json:"userId,omitempty"`
    UserIdType string `json:"userIdType,omitempty"`
}

