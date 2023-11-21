package dto

type VerifyDeleteAccountRequestDto struct {
	VerifyMethod         string                           `json:"verifyMethod"`
	PhonePassCodePayload *DeleteAccountByPhonePassCodeDto `json:"phonePassCodePayload,omitempty"`
	EmailPassCodePayload *DeleteAccountByEmailPassCodeDto `json:"emailPassCodePayload,omitempty"`
	PasswordPayload      *DeleteAccountByPasswordDto      `json:"passwordPayload,omitempty"`
}
