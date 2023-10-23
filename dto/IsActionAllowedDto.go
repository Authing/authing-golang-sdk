package dto


type IsActionAllowedDto struct{
    UserId  string `json:"userId"`
    Action  string `json:"action"`
    Resource  string `json:"resource"`
    Namespace  string `json:"namespace,omitempty"`
}

