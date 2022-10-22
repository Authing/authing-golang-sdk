package dto


type CurrentUsageDto struct{
    Amount  string `json:"amount"`
    Current  string `json:"current"`
    Experience  bool `json:"experience"`
    ModelCode  string `json:"modelCode"`
    ModelName  string `json:"modelName"`
}

