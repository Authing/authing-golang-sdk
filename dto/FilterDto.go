package dto


type FilterDto struct{
    ModelId  string `json:"modelId"`
    Keywords  string `json:"keywords,omitempty"`
    Conjunction  string `json:"conjunction,omitempty"`
    Conditions  []Condition `json:"conditions,omitempty"`
    Sort  []interface{} `json:"sort,omitempty"`
    Page  int `json:"page,omitempty"`
    Limit  int `json:"limit,omitempty"`
    FetchAll  bool `json:"fetchAll,omitempty"`
    WithPath  bool `json:"withPath,omitempty"`
    ShowFieldId  bool `json:"showFieldId,omitempty"`
    PreviewRelation  bool `json:"previewRelation,omitempty"`
    GetRelationFieldDetail  bool `json:"getRelationFieldDetail,omitempty"`
    Scope  ScopeDto `json:"scope,omitempty"`
    FilterRelation  ScopeDto `json:"filterRelation,omitempty"`
    Expand  []Expand `json:"expand,omitempty"`
}

