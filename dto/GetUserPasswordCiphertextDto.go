package dto


type GetUserPasswordCiphertextDto struct{
    UserId  string `json:"userId"`
    UserIdType  string  `json:"userIdType,omitempty"`
}

