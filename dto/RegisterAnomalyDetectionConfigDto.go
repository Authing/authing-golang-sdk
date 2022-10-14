package dto


type RegisterAnomalyDetectionConfigDto struct{
    Enabled  bool `json:"enabled"`
    Limit  int `json:"limit"`
    TimeInterval  int `json:"timeInterval"`
}

