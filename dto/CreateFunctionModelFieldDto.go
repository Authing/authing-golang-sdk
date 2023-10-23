package dto


type CreateFunctionModelFieldDto struct{
    UserVisible  bool `json:"userVisible"`
    RelationOptionalRange  Condition `json:"relationOptionalRange"`
    RelationShowKey  string `json:"relationShowKey"`
    RelationMultiple  bool `json:"relationMultiple"`
    RelationType  string `json:"relationType"`
    ForLogin  bool `json:"forLogin"`
    FuzzySearch  bool `json:"fuzzySearch"`
    DropDown  DropDownItemDto `json:"dropDown"`
    Format  int `json:"format"`
    Regexp  string `json:"regexp"`
    Min  int `json:"min"`
    Max  int `json:"max"`
    MaxLength  int `json:"maxLength"`
    Unique  bool `json:"unique"`
    Require  bool `json:"require"`
    Default  interface{} `json:"default"`
    Help  string `json:"help"`
    Editable  bool `json:"editable"`
    Show  bool `json:"show"`
    Type  string  `json:"type"`
    Key  string `json:"key"`
    Name  string `json:"name"`
    ModelId  string `json:"modelId"`
}

