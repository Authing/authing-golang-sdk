package dto


type SubjectDto struct{
    Id  string `json:"id"`
    Type  string  `json:"type"`
    Name  string `json:"name,omitempty"`
}

