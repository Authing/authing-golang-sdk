package dto


type PipelineFunctionDto struct{
    FuncId  string `json:"funcId"`
    FuncName  string `json:"funcName"`
    FuncDescription  string `json:"funcDescription,omitempty"`
    Scene  string  `json:"scene"`
    CreatedAt  string `json:"createdAt"`
    UpdatedAt  string `json:"updatedAt"`
    IsAsynchronous  bool `json:"isAsynchronous"`
    Timeout  int `json:"timeout"`
    TerminateOnTimeout  bool `json:"terminateOnTimeout"`
    SourceCode  string `json:"sourceCode"`
    Status  string  `json:"status"`
    UploadErrMsg  string `json:"uploadErrMsg,omitempty"`
    Enabled  bool `json:"enabled"`
}

