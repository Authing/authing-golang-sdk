package dto

type AuthorizeResourceItem struct {
	TargetType        string            `json:"targetType"`
	TargetIdentifiers []string          `json:"targetIdentifiers"`
	Resources         []ResourceItemDto `json:"resources"`
}
