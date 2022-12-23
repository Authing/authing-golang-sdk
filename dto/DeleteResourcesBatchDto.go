package dto


type DeleteResourcesBatchDto struct{
    Namespace  string `json:"namespace,omitempty"`
    CodeList  []string `json:"codeList,omitempty"`
    Ids  []string `json:"ids,omitempty"`
}

