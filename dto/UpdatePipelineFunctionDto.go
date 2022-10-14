package dto


type UpdatePipelineFunctionDto struct{
    FuncId  string `json:"funcId"`
    FuncName  string `json:"funcName,omitempty"`
    FuncDescription  string `json:"funcDescription,omitempty"`
    SourceCode  string `json:"sourceCode,omitempty"`
    IsAsynchronous  bool `json:"isAsynchronous,omitempty"`
    Timeout  int `json:"timeout,omitempty"`
    TerminateOnTimeout  bool `json:"terminateOnTimeout,omitempty"`
    Enabled  bool `json:"enabled,omitempty"`
}

