package dto

type ResourceItemDto struct {
	Code         string   `json:"code"`
	Actions      []string `json:"actions"`
	ResourceType string   `json:"resourceType"`
}
