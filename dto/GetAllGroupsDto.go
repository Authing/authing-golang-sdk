package dto


type GetAllGroupsDto struct{
    FetchMembers bool `json:"fetchMembers,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
}

