package dto


type GetUserLoggedInIdentitiesDto struct{
    UserId string `json:"userId,omitempty"`
    UserIdType string `json:"userIdType,omitempty"`
}

