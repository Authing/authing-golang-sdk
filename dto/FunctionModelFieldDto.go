package dto


type FunctionModelFieldDto struct{
    Id  string `json:"id"`
    ModelId  string `json:"modelId"`
    Name  string `json:"name"`
    Key  string `json:"key"`
    Type  int `json:"type"`
    Show  bool `json:"show"`
    Editable  bool `json:"editable"`
    Help  string `json:"help"`
    Default  string `json:"default"`
    Require  bool `json:"require"`
    Unique  bool `json:"unique"`
    MaxLength  int `json:"maxLength"`
    Max  int `json:"max"`
    Min  int `json:"min"`
    Regexp  string `json:"regexp"`
    Format  int `json:"format"`
    DropDown  int `json:"dropDown"`
    RelationType  string `json:"relationType"`
    RelationMultiple  bool `json:"relationMultiple"`
    RelationShowKey  string `json:"relationShowKey"`
    RelationOptionalRange  RelationOptionalRange `json:"relationOptionalRange"`
    UserVisible  bool `json:"userVisible"`
}

