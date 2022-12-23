package dto

type LoginAnomalyDetectionConfigDto struct {
	LoginFailStrategy                 string                          `json:"loginFailStrategy"`
	RobotVerify                       string                          `json:"robotVerify"`
	AccountLock                       string                          `json:"accountLock"`
	LoginFailCheck                    LoginFailCheckConfigDto         `json:"loginFailCheck"`
	LoginPasswordFailCheck            LoginPassowrdFailCheckConfigDto `json:"loginPasswordFailCheck"`
	AccountLockLoginPasswordFailCheck LoginPassowrdFailCheckConfigDto `json:"accountLockLoginPasswordFailCheck,omitempty"`
	RobotVerifyLoginPasswordFailCheck LoginPassowrdFailCheckConfigDto `json:"robotVerifyLoginPasswordFailCheck"`
	RobotVerifyLoginIpWhitelistCheck  LoginIpWhitelistCheckConfigDto  `json:"robotVerifyLoginIpWhitelistCheck"`
	RobotVerifyLoginTimeCheckEnable   bool                            `json:"robotVerifyLoginTimeCheckEnable"`
	RobotVerifyloginWeekStartEndTime  []string                        `json:"robotVerifyloginWeekStartEndTime"`
}
