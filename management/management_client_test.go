package management

import (
	"authing-go-sdk/dto"
	"fmt"
	"testing"
)

var client *Client

func init() {
	options := ClientOptions{
		AccessKeyId:     "60e043f8cd91b87d712b6365",
		AccessKeySecret: "158c7679333bc196b524d78d745813e5",
	}
	var err error
	client, err = NewClient(&options)
	if err != nil {
		panic(err)
	}
}

func TestClient_GetResource(t *testing.T) {
	request := dto.GetResourceDto{
		Code:      "code_3373",
		Namespace: "namespace_6254",
	}
	response := client.GetResource(&request)
	fmt.Println(response)

}

func TestClient_GetNamespace(t *testing.T) {
	request := dto.GetNamespaceDto{
		Code: "code_1335",
	}
	response := client.GetNamespace(&request)
	fmt.Println(response)

}

func TestClient_HasAnyRole(t *testing.T) {
	request := dto.HasAnyRoleReqDto{
		Roles: []dto.HasRoleRolesDto{
			{
				Namespace: "default",
				Code:      "admin",
			},
		},
		UserId: "611a149db64310ca4764ab15",
	}
	response := client.HasAnyRole(&request)
	fmt.Println(response)

}

func TestClient_SetCustomData(t *testing.T) {
	request := dto.SetCustomDataReqDto{
		List:             nil,
		TargetIdentifier: "targetIdentifier_9468",
		TargetType:       "",
		Namespace:        "namespace_4136",
	}
	response := client.SetCustomData(&request)
	fmt.Println(response)

}

func TestClient_GetCustomData(t *testing.T) {
	request := dto.GetCustomDataDto{
		TargetType:       "targetType_4930",
		TargetIdentifier: "targetIdentifier_5664",
		Namespace:        "namespace_6211",
	}
	response := client.GetCustomData(&request)
	fmt.Println(response)

}

func TestClient_DeleteNamespace(t *testing.T) {
	request := dto.DeleteNamespaceDto{
		Code: "code_4611",
	}
	response := client.DeleteNamespace(&request)
	fmt.Println(response)

}

func TestClient_AuthorizeResources(t *testing.T) {
	request := dto.AuthorizeResourcesDto{
		List:      nil,
		Namespace: "namespace_3067",
	}
	response := client.AuthorizeResources(&request)
	fmt.Println(response)

}

func TestClient_GetManagementToken(t *testing.T) {
	request := dto.GetManagementAccessTokenDto{
		AccessKeySecret: "accessKeySecret_7021",
		AccessKeyId:     "accessKeyId_4717",
	}
	response := client.GetManagementToken(&request)
	fmt.Println(response)

}

func TestClient_GetUser(t *testing.T) {
	request := dto.GetUserDto{
		UserId:            "611a149db64310ca4764ab15",
		WithCustomData:    false,
		WithIdentities:    false,
		WithDepartmentIds: false,
	}
	response := client.GetUser(&request)
	fmt.Println(response)

}

func TestClient_ListUsers(t *testing.T) {
	request := dto.ListUsersDto{
		Page:              1,
		Limit:             10,
		WithCustomData:    false,
		WithIdentities:    true,
		WithDepartmentIds: false,
	}
	response := client.ListUsers(&request)
	fmt.Println(response)

}

func TestClient_GetUserBatch(t *testing.T) {
	request := dto.GetUserBatchDto{
		UserIds:           "611a149db64310ca4764ab15,61bc4e94fc5809256b7c0333",
		WithCustomData:    false,
		WithIdentities:    false,
		WithDepartmentIds: false,
	}
	response := client.GetUserBatch(&request)
	fmt.Println(response)

}

func TestClient_GetUserIdentities(t *testing.T) {
	request := dto.GetUserIdentitiesDto{
		UserId: "611a149db64310ca4764ab15",
	}
	response := client.GetUserIdentities(&request)
	fmt.Println(response)

}

func TestClient_CreateOrganization(t *testing.T) {
	request := dto.CreateOrganizationReqDto{
		OrganizationName: "蒸汽记忆",
		OrganizationCode: "steamory",
	}
	response := client.CreateOrganization(&request)
	fmt.Println(response)

}

func TestClient_UpdateExtIdpConn(t *testing.T) {
	request := dto.UpdateExtIdpConnDto{
		Fields:      nil,
		DisplayName: "displayName_8594",
		Id:          "id_5185",
		Logo:        "logo_8928",
		LoginOnly:   false,
	}
	response := client.UpdateExtIdpConn(&request)
	fmt.Println(response)

}

func TestClient_DeleteOrganization(t *testing.T) {
	request := dto.DeleteOrganizationReqDto{
		OrganizationCode: "steamory2",
	}
	response := client.DeleteOrganization(&request)
	fmt.Println(response)

}

func TestClient_ChangeConnState(t *testing.T) {
	request := dto.EnableExtIdpConnDto{
		AppId:    "appId_8921",
		Enabled:  false,
		Id:       "id_2921",
		TenantId: "tenantId_7497",
	}
	response := client.ChangeConnState(&request)
	fmt.Println(response)

}

func TestClient_UpdateOrganization(t *testing.T) {
	request := dto.UpdateOrganizationReqDto{
		OrganizationName:    "蒸汽记忆",
		OrganizationCode:    "steamory",
		OrganizationNewCode: "steamory2",
	}
	response := client.UpdateOrganization(&request)
	fmt.Println(response)

}

func TestClient_SetCustomFields(t *testing.T) {
	request := dto.SetCustomFieldsReqDto{
		List: nil,
	}
	response := client.SetCustomFields(&request)
	fmt.Println(response)

}

func TestClient_AssignRole(t *testing.T) {
	request := dto.AssignRoleDto{
		Targets: []dto.TargetDto{
			{
				TargetType:       "USER",
				TargetIdentifier: "6291f220ec919ec7237fcd9f",
			},
		},
		Code:      "admin",
		Namespace: "default",
	}
	response := client.AssignRole(&request)
	fmt.Println(response)

}

func TestClient_CreateGroupBatch(t *testing.T) {
	request := dto.CreateGroupBatchReqDto{
		List: []dto.CreateGroupReqDto{
			{
				Description: "描述",
				Name:        "北京用户",
				Code:        "user_bj",
			},
			{
				Description: "描述",
				Name:        "上海用户",
				Code:        "user_sh",
			},
		},
	}
	response := client.CreateGroupsBatch(&request)
	fmt.Println(response)

}

func TestClient_GetUserMfaInfo(t *testing.T) {
	request := dto.GetUserMfaInfoDto{
		UserId: "611a149db64310ca4764ab15",
	}
	response := client.GetUserMfaInfo(&request)
	fmt.Println(response)

}

func TestClient_GetRole(t *testing.T) {
	request := dto.GetRoleDto{
		Code:      "code_70",
		Namespace: "default",
	}
	response := client.GetRole(&request)
	fmt.Println(response)

}

func TestClient_UpdateUser(t *testing.T) {
	request := dto.UpdateUserReqDto{
		UserId:     "6291f0b39d52f02fea96f862",
		Name:       "name_8726",
		Nickname:   "nickname_6206",
		Photo:      "photo_7414",
		ExternalId: "externalId_3102",
	}
	response := client.UpdateUser(&request)
	fmt.Println(response)

}

func TestClient_AddGroupMembers(t *testing.T) {
	request := dto.AddGroupMembersReqDto{
		UserIds: []string{"611a149db64310ca4764ab15"},
		Code:    "74wr2RzVV0",
	}
	response := client.AddGroupMembers(&request)
	fmt.Println(response)

}

/*func TestClient_RevokeRoleBatch(t *testing.T) {
	request := dto.RevokeRoleBatchDto{
		Targets: []dto.TargetDto{
			{
				TargetType:       "USER",
				TargetIdentifier: "6291f220ec919ec7237fcd9f",
			},
		},
		Roles: []dto.RoleCodeDto{
			{
				Code:      "admin",
				Namespace: "default",
			},
		},
	}
	response := management.RevokeRoleBatch(&request)
	fmt.Println(response)

}*/

func TestClient_DeleteDepartment(t *testing.T) {
	request := dto.DeleteDepartmentReqDto{
		OrganizationCode: "organizationCode_7898",
		DepartmentId:     "departmentId_8897",
		DepartmentIdType: "",
	}
	response := client.DeleteDepartment(&request)
	fmt.Println(response)

}

func TestClient_GetResourcesBatch(t *testing.T) {
	request := dto.GetResourcesBatchDto{
		CodeList:  "codeList_1627",
		Namespace: "namespace_2983",
	}
	response := client.GetResourcesBatch(&request)
	fmt.Println(response)

}

func TestClient_CreateRole(t *testing.T) {
	request := dto.CreateRoleDto{
		Code:        "code_70",
		Namespace:   "default",
		Description: "description_3228",
	}
	response := client.CreateRole(&request)
	fmt.Println(response)

}

func TestClient_GetNamespacesBatch(t *testing.T) {
	request := dto.GetNamespacesBatchDto{
		CodeList: "codeList_1740",
	}
	response := client.GetNamespacesBatch(&request)
	fmt.Println(response)

}

func TestClient_ListResources(t *testing.T) {
	request := dto.ListResourcesDto{
		Namespace: "namespace_3029",
		Type:      "type_3240",
		Page:      0,
		Limit:     0,
	}
	response := client.ListResources(&request)
	fmt.Println(response)

}

func TestClient_GetExtIdp(t *testing.T) {
	request := dto.GetExtIdpDto{
		Id:       "id_8918",
		TenantId: "tenantId_7944",
	}
	response := client.GetExtIdp(&request)
	fmt.Println(response)

}

func TestClient_ListArchivedUsers(t *testing.T) {
	request := dto.ListArchivedUsersDto{
		Page:  1,
		Limit: 10,
	}
	response := client.ListArchivedUsers(&request)
	fmt.Println(response)

}

func TestClient_DeleteRolesBatch(t *testing.T) {
	request := dto.DeleteRoleDto{
		CodeList:  []string{"code_71", "code_2"},
		Namespace: "default",
	}
	response := client.DeleteRolesBatch(&request)
	fmt.Println(response)

}

func TestClient_CreateUser(t *testing.T) {
	request := dto.CreateUserReqDto{
		Status:              "",
		Email:               "email_5835@wm.com",
		PasswordEncryptType: "",

		Phone:            "18310641125",
		PhoneCountryCode: "phoneCountryCode_6088",
		Username:         "username_3276",
		Name:             "name_3190",
		Nickname:         "nickname_5535",
		Photo:            "photo_1501",
		Gender:           "",

		EmailVerified: false,
		PhoneVerified: false,
		Birthdate:     "birthdate_2271",
		Country:       "country_405",
		Province:      "province_5583",
		City:          "city_4082",
		Address:       "address_7628",
		StreetAddress: "streetAddress_2665",
		PostalCode:    "postalCode_2513",
		ExternalId:    "externalId_784",
		DepartmentIds: nil,
		CustomData:    nil,
		Password:      "password_3785",
		TenantIds:     nil,
		Identities:    nil,
		Options: dto.CreateUserOptionsDto{
			KeepPassword:              false,
			ResetPasswordOnFirstLogin: false,
			DepartmentIdType:          "",
		},
	}
	response := client.CreateUser(&request)
	fmt.Println(response)

}

func TestClient_SearchDepartments(t *testing.T) {
	request := dto.SearchDepartmentsReqDto{
		Keywords:         "search_5854",
		OrganizationCode: "organizationCode_1752",
	}
	response := client.SearchDepartments(&request)
	fmt.Println(response)

}

func TestClient_GetUserGroups(t *testing.T) {
	request := dto.GetUserGroupsDto{
		UserId: "611a149db64310ca4764ab15",
	}
	response := client.GetUserGroups(&request)
	fmt.Println(response)

}

func TestClient_IsUserExists(t *testing.T) {
	request := dto.IsUserExistsReqDto{
		Username:   "",
		Email:      "zy@wm.com",
		Phone:      "",
		ExternalId: "",
	}
	response := client.IsUserExists(&request)
	fmt.Println(response)

}

func TestClient_KickUsers(t *testing.T) {
	request := dto.KickUsersDto{
		AppIds: []string{"61cbf5e0c7ce9e8282b85475", "61235292f1bbe2440809802d"},
		UserId: "611a149db64310ca4764ab15",
	}
	response := client.KickUsers(&request)
	fmt.Println(response)

}

func TestClient_CreateUserBatch(t *testing.T) {

	request := dto.CreateUserBatchReqDto{
		List: []dto.CreateUserInfoDto{
			{
				Email:      "emai-112@wm.com",
				Phone:      "18310641126",
				Username:   "Username",
				ExternalId: "ExternalId",
			},
		},
		Options: dto.CreateUserOptionsDto{},
	}
	response := client.CreateUserBatch(&request)
	fmt.Println(response)

}

func TestClient_RemoveGroupMembers(t *testing.T) {
	request := dto.RemoveGroupMembersReqDto{
		UserIds: []string{"611a149db64310ca4764ab15"},
		Code:    "74wr2RzVV0",
	}
	response := client.RemoveGroupMembers(&request)
	fmt.Println(response)

}

func TestClient_ListRoleMembers(t *testing.T) {
	request := dto.ListRoleMembersDto{
		Code:              "admin",
		Page:              1,
		Limit:             10,
		WithCustomData:    false,
		WithIdentities:    false,
		WithDepartmentIds: false,
		Namespace:         "default",
	}
	response := client.ListRoleMembers(&request)
	fmt.Println(response)

}

func TestClient_CreateExtIdpConn(t *testing.T) {
	request := dto.CreateExtIdpConnDto{
		Fields:      nil,
		DisplayName: "displayName_7799",
		Identifier:  "identifier_6069",
		Type:        "",

		ExtIdpId:  "extIdpId_9458",
		LoginOnly: false,
		Logo:      "logo_7996",
	}
	response := client.CreateExtIdpConn(&request)
	fmt.Println(response)

}

/*func TestClient_AssignRoleBatch(t *testing.T) {
	request := dto.AssignRoleBatchDto{
		Targets: []dto.TargetDto{
			{
				TargetType:       "USER",
				TargetIdentifier: "6291f220ec919ec7237fcd9f",
			},
		},
		Roles: []dto.RoleCodeDto{
			{
				Code:      "admin",
				Namespace: "default",
			},
		},
	}
	response := management.AssignRoleBatch(&request)
	fmt.Println(response)

}*/

func TestClient_DeleteExtIdpConn(t *testing.T) {
	request := dto.DeleteExtIdpConnDto{
		Id: "id_3553",
	}
	response := client.DeleteExtIdpConn(&request)
	fmt.Println(response)

}

func TestClient_GetGroup(t *testing.T) {
	request := dto.GetGroupDto{
		Code: "pm",
	}
	response := client.GetGroup(&request)
	fmt.Println(response)

}

func TestClient_GetCustomFields(t *testing.T) {
	request := dto.GetCustomFieldsDto{
		TargetType: "targetType_6791",
	}
	response := client.GetCustomFields(&request)
	fmt.Println(response)

}

func TestClient_CreateResource(t *testing.T) {
	request := dto.CreateResourceDto{
		Type: "",

		Code:          "code_6877",
		Description:   "description_4762",
		Actions:       nil,
		ApiIdentifier: "apiIdentifier_521",
		Namespace:     "namespace_519",
	}
	response := client.CreateResource(&request)
	fmt.Println(response)

}

func TestClient_UpdateResource(t *testing.T) {
	request := dto.UpdateResourceDto{
		Code:          "code_3665",
		Description:   "description_3254",
		Actions:       nil,
		ApiIdentifier: "apiIdentifier_5164",
		Namespace:     "namespace_2270",
		Type:          "",
	}
	response := client.UpdateResource(&request)
	fmt.Println(response)

}

func TestClient_ListRoles(t *testing.T) {
	request := dto.ListRolesDto{
		Namespace: "default",
		Page:      1,
		Limit:     10,
	}
	response := client.ListRoles(&request)
	fmt.Println(response)

}

func TestClient_CreateRolesBatch(t *testing.T) {
	request := dto.CreateRolesBatch{
		List: []dto.RoleListItem{
			{
				Code:        "code_1",
				Description: "description_1",
				Namespace:   "default",
			},
			{
				Code:        "code_2",
				Description: "description_2",
				Namespace:   "default",
			},
		},
	}
	response := client.CreateRolesBatch(&request)
	fmt.Println(response)

}

func TestClient_DeleteExtIdp(t *testing.T) {
	request := dto.DeleteExtIdpDto{
		Id: "id_5139",
	}
	response := client.DeleteExtIdp(&request)
	fmt.Println(response)

}

func TestClient_CreateNamespace(t *testing.T) {
	request := dto.CreateNamespaceDto{
		Code:        "code_7508",
		Name:        "name_9713",
		Description: "description_8998",
	}
	response := client.CreateNamespace(&request)
	fmt.Println(response)

}

func TestClient_DeleteResource(t *testing.T) {
	request := dto.DeleteResourceDto{
		Code:      "code_2243",
		Namespace: "namespace_1961",
	}
	response := client.DeleteResource(&request)
	fmt.Println(response)

}

func TestClient_GetDepartment(t *testing.T) {
	request := dto.GetDepartmentDto{
		OrganizationCode: "steamory",
		DepartmentId:     "6267ba6f09caf94dbe2ad291",
	}
	response := client.GetDepartment(&request)
	fmt.Println(response)

}

func TestClient_RevokeRole(t *testing.T) {
	request := dto.RevokeRoleDto{
		Targets: []dto.TargetDto{
			{
				TargetType:       "USER",
				TargetIdentifier: "6291f220ec919ec7237fcd9f",
			},
		},
		Code:      "admin",
		Namespace: "default",
	}
	response := client.RevokeRole(&request)
	fmt.Println(response)

}

func TestClient_GetUserRoles(t *testing.T) {
	request := dto.GetUserRolesDto{
		UserId:    "611a149db64310ca4764ab15",
		Namespace: "default",
	}
	response := client.GetUserRoles(&request)
	fmt.Println(response)

}

func TestClient_CreateGroup(t *testing.T) {
	request := dto.CreateGroupReqDto{
		Description: "description_715",
		Name:        "name_1878",
		Code:        "code_6657",
	}
	response := client.CreateGroup(&request)
	fmt.Println(response)

}

func TestClient_SetUserDepartment(t *testing.T) {
	var departments []dto.SetUserDepartmentDto
	departments = append(departments, dto.SetUserDepartmentDto{
		DepartmentId:     "60eea7b8f96188800df2e0bd",
		IsLeader:         false,
		IsMainDepartment: false,
	})
	request := dto.SetUserDepartmentsDto{
		Departments: departments,
		UserId:      "611a149db64310ca4764ab15",
	}
	response := client.SetUserDepartment(&request)
	fmt.Println(response)

}

func TestClient_ListOrganizations(t *testing.T) {
	request := dto.ListOrganizationsDto{
		Page:  1,
		Limit: 10,
	}
	response := client.ListOrganizations(&request)
	fmt.Println(response)

}

func TestClient_CreateExtIdp(t *testing.T) {
	request := dto.CreateExtIdpDto{
		Type: "",

		Name:     "name_1550",
		TenantId: "tenantId_3102",
	}
	response := client.CreateExtIdp(&request)
	fmt.Println(response)

}

func TestClient_UpdateNamespace(t *testing.T) {
	request := dto.UpdateNamespaceDto{
		Code:        "code_8527",
		Description: "description_6479",
		Name:        "name_4334",
		NewCode:     "newCode_8628",
	}
	response := client.UpdateNamespace(&request)
	fmt.Println(response)

}

func TestClient_DeleteUserBatch(t *testing.T) {
	request := dto.DeleteUsersBatchDto{
		UserIds: []string{"61075957e56172777e9f06e9", "610758491a92ce8b3a5ec41c"},
	}
	response := client.DeleteUsersBatch(&request)
	fmt.Println(response)

}

func TestClient_GetGroupList(t *testing.T) {
	request := dto.ListGroupsDto{
		Page:  1,
		Limit: 10,
	}
	response := client.ListGroups(&request)
	fmt.Println(response)

}

func TestClient_ListGroupMembers(t *testing.T) {
	request := dto.ListGroupMembersDto{
		Code:              "pm",
		Page:              1,
		Limit:             10,
		WithCustomData:    false,
		WithIdentities:    true,
		WithDepartmentIds: false,
	}
	response := client.ListGroupMembers(&request)
	fmt.Println(response)

}

func TestClient_ListExtIdp(t *testing.T) {
	request := dto.ListExtIdpDto{
		TenantId: "tenantId_1328",
	}
	response := client.ListExtIdp(&request)
	fmt.Println(response)

}

func TestClient_UpdateDepartment(t *testing.T) {
	request := dto.UpdateDepartmentReqDto{
		OrganizationCode:   "steamory",
		DepartmentId:       "6267c25597f8fd5757943b65",
		ParentDepartmentId: "",
		Code:               "development2",
		Name:               "研发部2",
	}
	response := client.UpdateDepartment(&request)
	fmt.Println(response)

}

func TestClient_GetUserDepartments(t *testing.T) {
	request := dto.GetUserDepartmentsDto{
		UserId: "611a149db64310ca4764ab15",
	}
	response := client.GetUserDepartments(&request)
	fmt.Println(response)

}

func TestClient_UpdateExtIdp(t *testing.T) {
	request := dto.UpdateExtIdpDto{
		Id:   "id_9024",
		Name: "name_1707",
	}
	response := client.UpdateExtIdp(&request)
	fmt.Println(response)

}

func TestClient_CreateDepartment(t *testing.T) {
	request := dto.CreateDepartmentReqDto{
		OrganizationCode: "steamory",
		Name:             "开发部2",
		Code:             "development2",
		OpenDepartmentId: "dpt_development2",
	}
	response := client.CreateDepartment(&request)
	fmt.Println(response)

}

func TestClient_UpdateGroup(t *testing.T) {
	request := dto.UpdateGroupReqDto{
		Description: "description_715",
		Name:        "广州用户",
		Code:        "code_6657",
		NewCode:     "user_gz",
	}
	response := client.UpdateGroup(&request)
	fmt.Println(response)

}

func TestClient_DeleteGroups(t *testing.T) {
	request := dto.DeleteGroupsReqDto{
		CodeList: []string{"v0l17C93S1", "5O24e7OW5b"},
	}
	response := client.DeleteGroupsBatch(&request)
	fmt.Println(response)

}

func TestClient_UpdateRole(t *testing.T) {
	request := dto.UpdateRoleDto{
		NewCode:     "code_71",
		Code:        "code_70",
		Namespace:   "default",
		Description: "description_377",
	}
	response := client.UpdateRole(&request)
	fmt.Println(response)

}

func TestClient_GetUserLoginHistory(t *testing.T) {
	request := dto.GetUserLoginHistoryDto{
		UserId: "611a149db64310ca4764ab15",
	}
	response := client.GetUserLoginHistory(&request)
	fmt.Println(response)

}

func TestClient_GetUserLoggedInApps(t *testing.T) {
	request := dto.GetUserLoggedinAppsDto{
		UserId: "611a149db64310ca4764ab15",
	}
	response := client.GetUserLoggedInApps(&request)
	fmt.Println(response)

}

func TestClient_GetUserAuthorizedApps(t *testing.T) {
	request := dto.GetUserAuthorizedAppsDto{
		UserId: "611a149db64310ca4764ab15",
	}
	response := client.GetUserAuthorizedApps(&request)
	fmt.Println(response)

}

func TestClient_GetGroupAuthorizedResources(t *testing.T) {
	request := dto.GetGroupAuthorizedResourcesDto{
		Code:         "pm",
		Namespace:    "default",
		ResourceType: "API",
	}
	response := client.GetGroupAuthorizedResources(&request)
	fmt.Println(response)

}

func TestClient_GetUserAccessibleApps(t *testing.T) {
	request := dto.GetUserAccessibleAppsDto{
		UserId: "611a149db64310ca4764ab15",
	}
	response := client.GetUserAccessibleApps(&request)
	fmt.Println(response)

}

func TestClient_GetUserAuthorizedResources(t *testing.T) {
	request := dto.GetUserAuthorizedResourcesDto{
		UserId: "611a149db64310ca4764ab15",
	}
	response := client.GetUserAuthorizedResources(&request)
	fmt.Println(response)

}

func TestClient_GetPrincipalAuthenticationInfo(t *testing.T) {
	request := dto.GetUserPrincipalAuthenticationInfoDto{
		UserId: "611a149db64310ca4764ab15",
	}
	response := client.GetUserPrincipalAuthenticationInfo(&request)
	fmt.Println(response)

}

func TestClient_ListChildrenDepartments(t *testing.T) {
	request := dto.ListChildrenDepartmentsDto{
		DepartmentId:     "departmentId_2132",
		OrganizationCode: "organizationCode_5389",
		DepartmentIdType: "departmentIdType_4221",
	}
	response := client.ListChildrenDepartments(&request)
	fmt.Println(response)

}

func TestClient_GetRoleAuthorizedResources(t *testing.T) {
	request := dto.GetRoleAuthorizedResourcesDto{
		Code:         "admin",
		Namespace:    "default",
		ResourceType: "API",
	}
	response := client.GetRoleAuthorizedResources(&request)
	fmt.Println(response)

}

func TestClient_ListDepartmentMembers(t *testing.T) {
	request := dto.ListDepartmentMembersDto{
		OrganizationCode:  "organizationCode_5347",
		DepartmentId:      "departmentId_2179",
		DepartmentIdType:  "departmentIdType_6377",
		Page:              0,
		Limit:             0,
		WithCustomData:    false,
		WithIdentities:    false,
		WithDepartmentIds: false,
	}
	response := client.ListDepartmentMembers(&request)
	fmt.Println(response)

}

func TestClient_AddDepartmentMembers(t *testing.T) {
	request := dto.AddDepartmentMembersReqDto{
		UserIds:          nil,
		OrganizationCode: "organizationCode_8148",
		DepartmentId:     "departmentId_6842",
		DepartmentIdType: "",
	}
	response := client.AddDepartmentMembers(&request)
	fmt.Println(response)

}

func TestClient_RemoveDepartmentMembers(t *testing.T) {
	request := dto.RemoveDepartmentMembersReqDto{
		UserIds:          nil,
		OrganizationCode: "organizationCode_7362",
		DepartmentId:     "departmentId_6293",
		DepartmentIdType: "",
	}
	response := client.RemoveDepartmentMembers(&request)
	fmt.Println(response)

}

func TestClient_ListRoleDepartments(t *testing.T) {
	request := dto.ListRoleDepartmentsDto{
		Code:      "admin",
		Namespace: "default",
		Page:      1,
		Limit:     10,
	}
	response := client.ListRoleDepartments(&request)
	fmt.Println(response)

}

func TestClient_GetParentDepartment(t *testing.T) {
	request := dto.GetParentDepartmentDto{
		OrganizationCode: "organizationCode_2019",
		DepartmentId:     "departmentId_3174",
		DepartmentIdType: "departmentIdType_7674",
	}
	response := client.GetParentDepartment(&request)
	fmt.Println(response)

}

func TestClient_CreateResourcesBatch(t *testing.T) {
	request := dto.CreateResourcesBatchDto{
		List:      nil,
		Namespace: "namespace_996",
	}
	response := client.CreateResourcesBatch(&request)
	fmt.Println(response)

}

func TestClient_DeleteResourcesBatch(t *testing.T) {
	request := dto.DeleteResourcesBatchDto{
		CodeList:  nil,
		Namespace: "namespace_9700",
	}
	response := client.DeleteResourcesBatch(&request)
	fmt.Println(response)

}

func TestClient_ListDepartmentMemberIds(t *testing.T) {
	request := dto.ListDepartmentMemberIdsDto{
		OrganizationCode: "organizationCode_4878",
		DepartmentId:     "departmentId_5224",
		DepartmentIdType: "departmentIdType_8269",
	}
	response := client.ListDepartmentMemberIds(&request)
	fmt.Println(response)

}

func TestClient_CreateNamespacesBatch(t *testing.T) {
	request := dto.CreateNamespacesBatchDto{
		List: nil,
	}
	response := client.CreateNamespacesBatch(&request)
	fmt.Println(response)

}

func TestClient_ResetUserPrincipalAuthenticationInfo(t *testing.T) {
	request := dto.ResetUserPrincipalAuthenticationInfoDto{
		UserId: "611a149db64310ca4764ab15",
	}
	response := client.ResetUserPrincipalAuthenticationInfo(&request)
	fmt.Println(response)

}

func TestClient_DeleteNamespacesBatch(t *testing.T) {
	request := dto.DeleteNamespacesBatchDto{
		CodeList: nil,
	}
	response := client.DeleteNamespacesBatch(&request)
	fmt.Println(response)

}
