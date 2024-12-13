package dto

type VerifyResetPasswordRequestDto struct {
	VerifyMethod         string                           `json:"verifyMethod"`
	PhonePassCodePayload *ResetPasswordByPhonePassCodeDto `json:"phonePassCodePayload,omitempty"`
	EmailPassCodePayload *ResetPasswordByEmailPassCodeDto `json:"emailPassCodePayload,omitempty"`
}
