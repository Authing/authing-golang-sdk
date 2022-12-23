package dto

type UserInfoDto struct {
	Username         string `json:"username,omitempty"`
	Email            string `json:"email,omitempty"`
	Phone            string `json:"phone,omitempty"`
	PhoneCountryCode string `json:"phoneCountryCode,omitempty"`
	Name             string `json:"name,omitempty"`
	Gender           string `json:"gender,omitempty"`
	Country          string `json:"country,omitempty"`
	Province         string `json:"province,omitempty"`
	City             string `json:"city,omitempty"`
}
