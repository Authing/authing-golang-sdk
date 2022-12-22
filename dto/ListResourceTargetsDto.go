package dto

type ListResourceTargetsDto struct {
	Resources     []string `json:"resources"`
	Actions       []string `json:"actions"`
	NamespaceCode string   `json:"namespaceCode"`
}
