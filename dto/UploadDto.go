package dto


type UploadDto struct{
    Folder  string `json:"folder,omitempty"`
    IsPrivate  bool `json:"isPrivate,omitempty"`
}

