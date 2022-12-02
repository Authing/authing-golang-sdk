package dto


type CheckUserSameLevelPermissionRespDto struct{
    Action  string `json:"action"`
    ResourceNodeCode  string `json:"resourceNodeCode,omitempty"`
    Enabled  bool `json:"enabled"`
}

