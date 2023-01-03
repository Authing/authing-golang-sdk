package dto

type GetUserResourceStructDto struct {
	ResourceCode  string `json:"resourceCode"`
	UserId        string `json:"userId"`
	NamespaceCode string `json:"namespaceCode"`
}
