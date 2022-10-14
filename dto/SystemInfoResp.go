package dto


type SystemInfoResp struct{
    Rsa  SystmeInfoRSAConfig `json:"rsa"`
    Sm2  SystmeInfoSM2Config `json:"sm2"`
    Version  SystmeInfoVersion `json:"version"`
    PublicIps  []string `json:"publicIps"`
}

