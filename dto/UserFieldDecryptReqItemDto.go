package dto


type UserFieldDecryptReqItemDto struct{
    UserId  string `json:"userId,omitempty"`
    FieldName  string `json:"fieldName,omitempty"`
    EncryptedContent  string `json:"encryptedContent"`
}

