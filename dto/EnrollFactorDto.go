package dto


type EnrollFactorDto struct{
    EnrollmentData  EnrollFactorEnrollmentDataDto `json:"enrollmentData"`
    EnrollmentToken  string `json:"enrollmentToken"`
    FactorType  string  `json:"factorType"`
}

