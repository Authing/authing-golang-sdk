package dto


type GetOtpSecretByUserDto struct{
    UserId string `json:"userId,omitempty"`
    UserIdType string `json:"userIdType,omitempty"`
}

