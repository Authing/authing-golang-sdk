package dto


type UserFieldDecryptRespItemDto struct{
    UserId  string `json:"userId,omitempty"`
    FieldName  string `json:"fieldName,omitempty"`
    DecryptedContent  string `json:"decryptedContent"`
}

