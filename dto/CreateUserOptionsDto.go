package dto


type CreateUserOptionsDto struct{
    KeepPassword  bool `json:"keepPassword,omitempty"`
    ResetPasswordOnFirstLogin  bool `json:"resetPasswordOnFirstLogin,omitempty"`
    DepartmentIdType  string  `json:"departmentIdType,omitempty"`
}

