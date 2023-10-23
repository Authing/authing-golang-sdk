package dto


type ResGroupDto struct{
    Id  string `json:"id"`
    Code  string `json:"code"`
    Name  string `json:"name"`
    Description  string `json:"description"`
    Type  string `json:"type"`
    MetadataSource  []string `json:"metadataSource"`
    Members  []UserDto `json:"members"`
}

