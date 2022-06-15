package dto

type CreateResourcesBatchDto struct {
	List      []CreateResourceBatchItemDto `json:"list"`
	Namespace string                       `json:"namespace,omitempty"`
}
