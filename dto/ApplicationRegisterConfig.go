package dto


type ApplicationRegisterConfig struct{
    EnabledBasicRegisterMethods  []string `json:"enabledBasicRegisterMethods"`
    DefaultRegisterMethod  string  `json:"defaultRegisterMethod"`
}

