package dto

type GetResourcesBatchDto struct {
	CodeList  string `json:"codeList,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}
