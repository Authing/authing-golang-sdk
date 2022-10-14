package dto


type ApplicationRegisterConfigInputDto struct{
    EnabledBasicRegisterMethods  []string `json:"enabledBasicRegisterMethods"`
    DefaultRegisterMethod  string  `json:"defaultRegisterMethod"`
}

