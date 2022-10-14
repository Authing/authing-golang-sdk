package dto


type CustomFieldDto struct{
    TargetType  string  `json:"targetType"`
    CreatedAt  string `json:"createdAt,omitempty"`
    DataType  string  `json:"dataType"`
    Key  string `json:"key"`
    Label  string `json:"label"`
    Description  string `json:"description,omitempty"`
    Encrypted  bool `json:"encrypted,omitempty"`
    IsUnique  bool `json:"isUnique"`
    UserEditable  bool `json:"userEditable,omitempty"`
    VisibleInAdminConsole  bool `json:"visibleInAdminConsole"`
    VisibleInUserCenter  bool `json:"visibleInUserCenter,omitempty"`
    I18n  CustomFieldI18n `json:"i18n,omitempty"`
    Options  []CustomFieldSelectOption `json:"options,omitempty"`
}

