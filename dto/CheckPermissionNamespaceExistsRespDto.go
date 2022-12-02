package dto


type CheckPermissionNamespaceExistsRespDto struct{
    IsValid  bool `json:"isValid"`
    Message  string `json:"message,omitempty"`
}

