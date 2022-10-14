package dto


type UpdatePhoneByPhonePassCodeDto struct{
    NewPhoneNumber  string `json:"newPhoneNumber"`
    NewPhonePassCode  string `json:"newPhonePassCode"`
    NewPhoneCountryCode  string `json:"newPhoneCountryCode,omitempty"`
    OldPhoneNumber  string `json:"oldPhoneNumber,omitempty"`
    OldPhonePassCode  string `json:"oldPhonePassCode,omitempty"`
    OldPhoneCountryCode  string `json:"oldPhoneCountryCode,omitempty"`
}

