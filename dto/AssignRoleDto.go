package dto


type AssignRoleDto struct{
    Targets  []TargetDto `json:"targets"`
    Code  string `json:"code"`
    EndTime  int `json:"endTime,omitempty"`
    EnableTime  int `json:"enableTime,omitempty"`
    Namespace  string `json:"namespace,omitempty"`
}

