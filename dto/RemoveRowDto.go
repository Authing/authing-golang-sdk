package dto


type RemoveRowDto struct{
    RowIdList  []string `json:"rowIdList"`
    ModelId  string `json:"modelId"`
    Recursive  bool `json:"recursive,omitempty"`
}

