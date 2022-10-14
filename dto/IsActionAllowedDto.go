package dto


type IsActionAllowedDto struct{
    Action  string `json:"action"`
    Resource  string `json:"resource"`
    UserId  string `json:"userId"`
    Namespace  string `json:"namespace,omitempty"`
}

