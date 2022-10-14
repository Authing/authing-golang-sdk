package dto


type SelfUnlockAccountConfigDto struct{
    Enabled  bool `json:"enabled"`
    Strategy  string  `json:"strategy"`
}

