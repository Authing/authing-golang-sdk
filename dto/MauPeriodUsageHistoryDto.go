package dto


type MauPeriodUsageHistoryDto struct{
    PeriodStartTime  string `json:"periodStartTime"`
    PeriodEndTime  string `json:"periodEndTime"`
    Amount  string `json:"amount"`
    Current  string `json:"current"`
}

