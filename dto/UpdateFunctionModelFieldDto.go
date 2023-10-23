package dto


type UpdateFunctionModelFieldDto struct{
    UserVisible  bool `json:"userVisible"`
    RelationOptionalRange  RelationOptionalRange `json:"relationOptionalRange"`
    RelationShowKey  string `json:"relationShowKey"`
    ForLogin  bool `json:"forLogin"`
    FuzzySearch  bool `json:"fuzzySearch"`
    DropDown  []string `json:"dropDown"`
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
    Name  string `json:"name"`
    ModelId  string `json:"modelId"`
    Id  string `json:"id"`
}

