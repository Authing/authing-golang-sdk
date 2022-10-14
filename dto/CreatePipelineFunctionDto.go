package dto


type CreatePipelineFunctionDto struct{
    SourceCode  string `json:"sourceCode"`
    Scene  string  `json:"scene"`
    FuncName  string `json:"funcName"`
    FuncDescription  string `json:"funcDescription,omitempty"`
    IsAsynchronous  bool `json:"isAsynchronous,omitempty"`
    Timeout  int `json:"timeout,omitempty"`
    TerminateOnTimeout  bool `json:"terminateOnTimeout,omitempty"`
    Enabled  bool `json:"enabled,omitempty"`
}

