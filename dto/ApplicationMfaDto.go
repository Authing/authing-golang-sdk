package dto


type ApplicationMfaDto struct{
    MfaPolicy  string `json:"mfaPolicy"`
    Status  int `json:"status"`
    Sort  int `json:"sort"`
}

