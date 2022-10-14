package dto


type GetEmailTemplatesDataDto struct{
    Templates  []EmailTemplateDto `json:"templates"`
    Categories  []EmailTemplateCategoryDto `json:"categories"`
}

