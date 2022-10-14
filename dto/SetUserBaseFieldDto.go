package dto


type SetUserBaseFieldDto struct{
    Key  string `json:"key"`
    Label  string `json:"label,omitempty"`
    Description  string `json:"description,omitempty"`
    UserEditable  bool `json:"userEditable,omitempty"`
    VisibleInAdminConsole  bool `json:"visibleInAdminConsole,omitempty"`
    VisibleInUserCenter  bool `json:"visibleInUserCenter,omitempty"`
    I18n  CustomFieldI18n `json:"i18n,omitempty"`
}

