package dto


type RowDto struct{
    RowId  string `json:"rowId"`
    CellList  []CellDto `json:"cellList"`
}

