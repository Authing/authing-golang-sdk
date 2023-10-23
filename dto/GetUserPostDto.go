package dto


type GetUserPostDto struct{
    UserId string `json:"userId,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
}

