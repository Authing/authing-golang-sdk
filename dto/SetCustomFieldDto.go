package dto


type SetCustomFieldDto struct{
    TargetType  string  `json:"targetType"`
    Key  string `json:"key"`
    DataType  string  `json:"dataType,omitempty"`
    Label  string `json:"label,omitempty"`
    Description  string `json:"description,omitempty"`
    Encrypted  bool `json:"encrypted,omitempty"`
    IsUnique  bool `json:"isUnique,omitempty"`
    UserEditable  bool `json:"userEditable,omitempty"`
    VisibleInAdminConsole  bool `json:"visibleInAdminConsole,omitempty"`
    VisibleInUserCenter  bool `json:"visibleInUserCenter,omitempty"`
    ValidateRules  interface{} `json:"validateRules,omitempty"`
    AppIds  []string `json:"appIds,omitempty"`
    Desensitization  bool `json:"desensitization,omitempty"`
    Options  []CustomFieldSelectOption `json:"options,omitempty"`
    I18n  CustomFieldI18n `json:"i18n,omitempty"`
}

