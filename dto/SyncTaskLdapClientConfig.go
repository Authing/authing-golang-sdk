package dto


type SyncTaskLdapClientConfig struct{
    Url  string `json:"url"`
    BindDn  string `json:"bindDn"`
    BindCredentials  string `json:"bindCredentials"`
    UsersBaseDn  string `json:"usersBaseDn"`
    GroupsBaseDn  string `json:"groupsBaseDn"`
    UserQueryCriteria  string `json:"userQueryCriteria"`
    DepartmentQueryCriteria  string `json:"departmentQueryCriteria"`
}

