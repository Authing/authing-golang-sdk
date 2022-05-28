package dto


type GetUserRolesDto struct{
    UserId string `json:"user_id,omitempty"`
    Namespace string `json:"namespace,omitempty"`
}

