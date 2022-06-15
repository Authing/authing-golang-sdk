package dto

type CustomFieldDto struct {
	TargetType  string                    `json:"targetType"`
	CreatedAt   string                    `json:"createdAt"`
	DataType    string                    `json:"dataType"`
	Key         string                    `json:"key"`
	Label       string                    `json:"label"`
	Description string                    `json:"description,omitempty"`
	Encrypted   bool                      `json:"encrypted,omitempty"`
	Options     []CustomFieldSelectOption `json:"options,omitempty"`
}
