package dto


type PipelineFunctionPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []PipelineFunctionDto `json:"list"`
}

