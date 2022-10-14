package dto


type CountryDetailInfo struct{
    Alpha2  string `json:"alpha2"`
    Alpha3  string `json:"alpha3"`
    PhoneCountryCode  string `json:"phoneCountryCode"`
    Flag  string `json:"flag"`
    Name  LangObject `json:"name"`
}

