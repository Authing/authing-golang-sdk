package dto

type GetExternalUserResourceStructDataDto struct {
	NamespaceCode          string                 `json:"namespaceCode"`
	ResourceCode           string                 `json:"resourceCode"`
	ResourceType           string                 `json:"resourceType"`
	StrResourceAuthAction  StrResourceAuthAction  `json:"strResourceAuthAction,omitempty"`
	ArrResourceAuthAction  ArrResourceAuthAction  `json:"arrResourceAuthAction,omitempty"`
	TreeResourceAuthAction TreeResourceAuthAction `json:"treeResourceAuthAction,omitempty"`
}
