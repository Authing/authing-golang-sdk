package dto

type SetCustomFieldDto struct {
	TargetType  string                    `json:"targetType"`
	DataType    string                    `json:"dataType"`
	Key         string                    `json:"key"`
	Label       string                    `json:"label"`
	Description string                    `json:"description,omitempty"`
	Encrypted   bool                      `json:"encrypted,omitempty"`
	Options     []CustomFieldSelectOption `json:"options,omitempty"`
}
