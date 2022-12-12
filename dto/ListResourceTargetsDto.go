package dto

type ListResourceTargetsDto struct {
	NamespaceCode string   `json:"namespaceCode"`
	Actions       []string `json:"actions"`
	Resources     []string `json:"resources"`
}
