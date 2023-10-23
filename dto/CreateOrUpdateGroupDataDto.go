package dto


type CreateOrUpdateGroupDataDto struct{
    Created  bool `json:"created"`
    Data  GroupDto `json:"data"`
}

