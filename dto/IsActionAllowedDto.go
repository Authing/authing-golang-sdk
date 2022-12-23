package dto


type IsActionAllowedDto struct{
    Resource  string `json:"resource"`
    Action  string `json:"action"`
    UserId  string `json:"userId"`
    Namespace  string `json:"namespace,omitempty"`
}

