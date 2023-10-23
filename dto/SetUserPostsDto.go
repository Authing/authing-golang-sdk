package dto


type SetUserPostsDto struct{
    PostIds  []string `json:"postIds"`
    UserId  string `json:"userId"`
}

