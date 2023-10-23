package dto


type GetMfaStatusDto struct{
    UserId string `json:"userId,omitempty"`
    UserIdType string `json:"userIdType,omitempty"`
}

