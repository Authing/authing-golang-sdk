package dto


type DeleteResourcesBatchDto struct{
    CodeList  []string `json:"codeList"`
    Namespace  string `json:"namespace,omitempty"`
}

