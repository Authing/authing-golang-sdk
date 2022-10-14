package dto


type SendEnrollFactorRequestDataDto struct{
    EnrollmentToken  string `json:"enrollmentToken"`
    OtpData  SendEnrollFactorRequestOtpDataDto `json:"otpData,omitempty"`
}

