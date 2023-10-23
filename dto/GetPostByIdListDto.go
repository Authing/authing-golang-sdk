package dto


type GetPostByIdListDto struct{
    IdList  string `json:"idList,omitempty"`
    WithCustomData  bool `json:"withCustomData,omitempty"`
}

