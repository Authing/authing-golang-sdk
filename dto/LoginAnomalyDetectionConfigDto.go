package dto


type LoginAnomalyDetectionConfigDto struct{
    LoginFailStrategy  string  `json:"loginFailStrategy"`
    LoginFailCheck  LoginFailCheckConfigDto `json:"loginFailCheck"`
    LoginPasswordFailCheck  LoginPassowrdFailCheckConfigDto `json:"loginPasswordFailCheck"`
}

