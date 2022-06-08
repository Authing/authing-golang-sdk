package dto

type GetUserAuthorizedResourcesDto struct {
	UserId       string `json:"userId,omitempty"`
	Namespace    string `json:"namespace,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
}
