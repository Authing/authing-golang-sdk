package dto


type GetUserAuthorizedAppsDto struct{
    UserId string `json:"userId,omitempty"`
    UserIdType string `json:"userIdType,omitempty"`
}

