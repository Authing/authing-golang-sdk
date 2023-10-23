package dto


type GetUserPostsDto struct{
    UserId string `json:"userId,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
}

