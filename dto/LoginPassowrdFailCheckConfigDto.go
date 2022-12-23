package dto


type LoginPassowrdFailCheckConfigDto struct{
    Enabled  bool `json:"enabled"`
    Limit  int `json:"limit"`
    TimeInterval  int `json:"timeInterval"`
    Unit  string `json:"unit,omitempty"`
}

