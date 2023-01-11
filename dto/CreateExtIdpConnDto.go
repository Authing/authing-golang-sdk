package dto

type CreateExtIdpConnDto struct {
	ExtIdpId    string      `json:"extIdpId"`
	Type        string      `json:"type"`
	Identifier  string      `json:"identifier"`
	DisplayName string      `json:"displayName"`
	Fields      interface{} `json:"fields"`
	LoginOnly   bool        `json:"loginOnly,omitempty"`
	Logo        string      `json:"logo,omitempty"`
	TenantId    string      `json:"tenantId,omitempty"`
}
