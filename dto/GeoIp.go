package dto


type GeoIp struct{
    Location  GeoIpLocation `json:"location"`
    CountryName  string `json:"country_name"`
    CountryCode2  string `json:"country_code2"`
    CountryCode3  string `json:"country_code3"`
    RegionName  string `json:"region_name"`
    RegionCode  string `json:"region_code"`
    CityName  string `json:"city_name"`
    ContinentCode  string `json:"continent_code"`
    Timezone  string `json:"timezone"`
}

