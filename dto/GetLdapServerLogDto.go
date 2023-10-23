package dto


type GetLdapServerLogDto struct{
    Type int `json:"type,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    Connection int `json:"connection,omitempty"`
    OperationNumber int `json:"operationNumber,omitempty"`
    ErrorCode int `json:"errorCode,omitempty"`
    Message string `json:"message,omitempty"`
    StartTime int `json:"startTime,omitempty"`
    EndTime int `json:"endTime,omitempty"`
}

