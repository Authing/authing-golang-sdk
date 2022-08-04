package dto

type UserDto struct {
	UserId                   string        `json:"userId"`
	CreatedAt                string        `json:"createdAt"`
	UpdatedAt                string        `json:"updatedAt"`
	Status                   string        `json:"status"`
	Email                    string        `json:"email,omitempty"`
	Phone                    string        `json:"phone,omitempty"`
	PhoneCountryCode         string        `json:"phoneCountryCode,omitempty"`
	Username                 string        `json:"username,omitempty"`
	Name                     string        `json:"name,omitempty"`
	Nickname                 string        `json:"nickname,omitempty"`
	Photo                    string        `json:"photo,omitempty"`
	LoginsCount              int           `json:"loginsCount,omitempty"`
	LastLogin                string        `json:"lastLogin,omitempty"`
	LastIp                   string        `json:"lastIp,omitempty"`
	Gender                   string        `json:"gender"`
	EmailVerified            bool          `json:"emailVerified"`
	PhoneVerified            bool          `json:"phoneVerified"`
	PasswordLastSetAt        string        `json:"passwordLastSetAt,omitempty"`
	Birthdate                string        `json:"birthdate,omitempty"`
	Country                  string        `json:"country,omitempty"`
	Province                 string        `json:"province,omitempty"`
	City                     string        `json:"city,omitempty"`
	Address                  string        `json:"address,omitempty"`
	StreetAddress            string        `json:"streetAddress,omitempty"`
	PostalCode               string        `json:"postalCode,omitempty"`
	ExternalId               string        `json:"externalId,omitempty"`
	ResetPasswordOnNextLogin bool          `json:"resetPasswordOnNextLogin,omitempty"`
	DepartmentIds            []string      `json:"departmentIds,omitempty"`
	Identities               []IdentityDto `json:"identities,omitempty"`
	CustomData               interface{}   `json:"customData,omitempty"`
	StatusChangedAt          string        `json:"statusChangedAt,omitempty"`
}
