package dto


type LoginFailCheckConfigDto struct{
    Enabled  bool `json:"enabled"`
    Limit  int `json:"limit"`
    TimeInterval  int `json:"timeInterval"`
}

