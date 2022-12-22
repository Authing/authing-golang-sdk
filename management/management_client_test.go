package management

import (
	"fmt"
	"github.com/Authing/authing-golang-sdk/v3/dto"
	"testing"
)

var client *ManagementClient

func init() {
	options := ManagementClientOptions{
		//AccessKeyId:     "YOUR_ACCESS_KEY_ID",
		//AccessKeySecret: "YOUR_ACCESS_KEY_SECRET",
		//Host:            "YOUR_HOST",
		AccessKeyId:     "63a3f1c38f7593d460057436",
		AccessKeySecret: "d70365bf03d334b95133fd27cdb1346f",
		Host:            "https://console.authing.cn",
	}
	var err error
	client, err = NewManagementClient(&options)
	print(client)
	if err != nil {
		panic(err)
	}

}

func TestClient_ListUsers(t *testing.T) {
	request := dto.ListUsersRequestDto{
		Options: dto.ListUsersOptionsDto{
			Pagination: dto.PaginationDto{
				Page:  2,
				Limit: 10,
			},
		},
	}
	response := client.ListUsers(&request)
	fmt.Println(response)
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
	request := dto.ChangeExtIdpConnStateDto{
		AppId:    "appId_8921",
		Enabled:  false,
		Id:       "id_2921",
		TenantId: "tenantId_7497",
	}
	response := client.ChangeExtIdpConnState(&request)
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
		Status:           "",
		Email:            "email_5835@wm.com",
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
	response := client.CreateUsersBatch(&request)
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
	response := client.SetUserDepartments(&request)
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
	response := client.GetUserLoggedinApps(&request)
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

//func TestClient_GetAuthorizedResources(t *testing.T) {
//	response := client.GetAuthorizedResources(&dto.GetAuthorizedResourcesDto{
//		TargetType:       "USER",
//		TargetIdentifier: "sdfs",
//		Namespace:        "default",
//		WithDenied:       false,
//	})
//	fmt.Println(response)
//}

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

// 创建权限空间
func TestClient_CreatePermissionNamespace(t *testing.T) {
	request := dto.CreatePermissionNamespaceDto{
		Name:        "示例权限空间",
		Code:        "examplePermissionNamespace",
		Description: "示例权限空间描述",
	}
	response := client.CreatePermissionNamespace(&request)
	fmt.Println(response)
}

// 删除权限空间
func TestClient_DeletePermissionNamespace(t *testing.T) {
	request := dto.DeletePermissionNamespaceDto{
		Code: "examplePermissionNamespace",
	}
	response := client.DeletePermissionNamespace(&request)
	fmt.Println(response)
}

// 批量创建权限空间
func TestClient_CreatePermissionNamespacesBatch(t *testing.T) {
	request := dto.CreatePermissionNamespacesBatchDto{
		List: []dto.CreatePermissionNamespacesBatchItemDto{
			{
				Name:        "示例权限空间1",
				Code:        "examplePermissionNamespace1",
				Description: "示例权限空间1描述",
			},
			{
				Name:        "示例权限空间2",
				Code:        "examplePermissionNamespace2",
				Description: "示例权限空间2描述",
			},
		},
	}
	response := client.CreatePermissionNamespacesBatch(&request)
	fmt.Println(response)
}

// 批量删除权限空间
func TestClient_DeletePermissionNamespacesBatch(t *testing.T) {
	request := dto.DeletePermissionNamespacesBatchDto{
		Codes: []string{
			"examplePermissionNamespace1",
			"examplePermissionNamespace2",
		},
	}
	response := client.DeletePermissionNamespacesBatch(&request)
	fmt.Println(response)
}

// 获取权限空间详情
func TestClient_GetPermissionNamespace(t *testing.T) {
	request := dto.GetPermissionNamespaceDto{
		Code: "examplePermissionNamespace",
	}
	response := client.GetPermissionNamespace(&request)
	fmt.Println(response)
}

// 批量获取权限空间详情
func TestClient_GetPermissionNamespacesBatch(t *testing.T) {
	request := dto.GetPermissionNamespacesBatchDto{
		Codes: "examplePermissionNamespace1,examplePermissionNamespace2",
	}
	response := client.GetPermissionNamespacesBatch(&request)
	fmt.Println(response)
}

// 更新权限空间
func TestClient_UpdatePermissionNamespace(t *testing.T) {
	request := dto.UpdatePermissionNamespaceDto{
		Code:        "examplePermissionNamespace",
		Name:        "示例新权限空间名称",
		Description: "示例新权限空间描述",
		NewCode:     "exampleNewPermissionNamespace",
	}
	response := client.UpdatePermissionNamespace(&request)
	fmt.Println(response)
}

// 获取权限空间列表
func TestClient_ListPermissionNamespaces(t *testing.T) {
	request := dto.ListPermissionNamespacesDto{
		Page:  1,
		Limit: 10,
		Query: "权限",
	}
	response := client.ListPermissionNamespaces(&request)
	fmt.Println(response)
}

// 校验权限空间 Code 是否可用
func TestClient_CheckPermissionNamespaceExists(t *testing.T) {
	request := dto.CheckPermissionNamespaceExistsDto{
		Code: "examplePermissionNamespace",
	}
	response := client.CheckPermissionNamespaceExists(&request)
	fmt.Println(response)

	request1 := dto.CheckPermissionNamespaceExistsDto{
		Name: "示例权限空间名称",
	}
	response1 := client.CheckPermissionNamespaceExists(&request1)
	fmt.Println(response1)

	request2 := dto.CheckPermissionNamespaceExistsDto{
		Name: "示例权限空间1",
	}
	response2 := client.CheckPermissionNamespaceExists(&request2)
	fmt.Println(response2)
}

// 获取权限空间下所有的角色列表
func TestClient_ListPermissionNamespaceRoles(t *testing.T) {
	request := dto.ListPermissionNamespaceRolesDto{
		Page:  1,
		Limit: 10,
		Query: "exampleRoleCodeOrName",
		Code:  "exampleNewPermissionNamespace",
	}
	response := client.ListPermissionNamespaceRoles(&request)
	fmt.Println(response)
}

// 创建字符串数据资源
func TestClient_CreateStringDataResource(t *testing.T) {
	request := dto.CreateDataResourceDto{
		Actions:       []string{"read", "get"},
		Struct:        "test",
		Type:          "STRING",
		ResourceCode:  "stringResourceCode",
		ResourceName:  "示例字符串数据资源",
		NamespaceCode: "exampleNewPermissionNamespace",
		Description:   "示例字符串数据资源描述",
	}
	response := client.CreateDataResource(&request)
	fmt.Println(response)
}

// 创建数组数据资源
func TestClient_CreateArrayDataResource(t *testing.T) {
	request := dto.CreateDataResourceDto{
		Actions:       []string{"read", "get"},
		Struct:        []string{"exampleArrayStruct1", "exampleArrayStruct2"},
		Type:          "ARRAY",
		ResourceCode:  "arrayResourceCode1",
		ResourceName:  "示例数组数据资源1",
		NamespaceCode: "examplePermissionNamespace",
		Description:   "示例数组数据资源描述",
	}
	response := client.CreateDataResource(&request)
	fmt.Println(response)
}

// 创建数数据资源
func TestClient_CreateTreeDataResource(t *testing.T) {
	request := dto.CreateTreeDataResourceDto{
		NamespaceCode: "examplePermissionNamespace",
		ResourceCode:  "treeResourceCode3",
		ResourceName:  "示例树数据资源3",
		Struct: []dto.DataResourceTreeStructs{
			{
				Code:  "tree1",
				Name:  "树节点1",
				Value: "树节点1描述",
				Children: []any{
					dto.DataResourceTreeStructs{
						Code:  "tree11",
						Name:  "树节点11",
						Value: "树节点11描述",
						Children: []any{
							dto.DataResourceTreeStructs{
								Code:  "tree111",
								Name:  "树节点111",
								Value: "树节点111描述",
							},
							dto.DataResourceTreeStructs{
								Code:  "tree112",
								Name:  "树节点112",
								Value: "树节点112描述",
							},
						},
					},
					dto.DataResourceTreeStructs{
						Code:  "tree12",
						Name:  "树节点12",
						Value: "树节点12描述",
					},
				},
			},
			{
				Code:  "tree2",
				Name:  "树节点2",
				Value: "树节点2描述",
			},
		},
		Description: "示例树数据资源描述",
		Actions:     []string{"get", "read"},
	}
	response := client.CreateDataResourceByTree(&request)
	fmt.Println(response)
}

// 创建字符串数据资源
func TestClient_CreateDataResourceByString(t *testing.T) {
	request := dto.CreateStringDataResourceDto{
		NamespaceCode: "examplePermissionNamespace",
		ResourceCode:  "stringResourceCode",
		ResourceName:  "示例字符串数据资源",
		Struct:        "exampleStringStruct",
		Description:   "示例字符串数据资源描述",
		Actions:       []string{"get", "read"},
	}
	response := client.CreateDataResourceByString(&request)
	fmt.Println(response)
}

// 创建数组数据资源
func TestClient_CreateDataResourceByArray(t *testing.T) {
	request := dto.CreateArrayDataResourceDto{
		NamespaceCode: "examplePermissionNamespace",
		ResourceCode:  "arrayResourceCode",
		ResourceName:  "示例数组数据资源",
		Struct:        []string{"exampleArrayStruct1", "exampleArrayStruct2"},
		Description:   "示例数组数据资源描述",
		Actions:       []string{"get", "read"},
	}
	response := client.CreateDataResourceByArray(&request)
	fmt.Println(response)
}

// 创建数数据资源
func TestClient_CreateDataResourceByTree(t *testing.T) {
	request := dto.CreateTreeDataResourceDto{
		NamespaceCode: "examplePermissionNamespace",
		ResourceCode:  "treeResourceCode2",
		ResourceName:  "示例树数据资源2",
		Struct: []dto.DataResourceTreeStructs{
			{
				Code:  "tree1",
				Name:  "树节点1",
				Value: "树节点1描述",
				Children: []any{
					dto.DataResourceTreeStructs{
						Code:  "tree11",
						Name:  "树节点11",
						Value: "树节点11描述",
					},
				},
			},
			{
				Code:  "tree2",
				Name:  "树节点2",
				Value: "树节点2描述",
			},
		},
		Description: "示例树数据资源描述",
		Actions:     []string{"get", "read"},
	}
	response := client.CreateDataResourceByTree(&request)
	fmt.Println(response)
}

// 删除数据资源
func TestClient_DeleteDataResource(t *testing.T) {
	request := dto.DeleteDataResourceDto{
		ResourceCode:  "arrayResourceCode",
		NamespaceCode: "examplePermissionNamespace",
	}
	response := client.DeleteDataResource(&request)
	fmt.Println(response)
}

// 更新数据资源
func TestClient_UpdateDataResource(t *testing.T) {
	request := dto.UpdateDataResourceDto{
		ResourceCode:  "stringResourceCode",
		ResourceName:  "示例新字符串数据资源",
		NamespaceCode: "examplePermissionNamespace",
		Description:   "示例数据资源新描述",
		Actions:       []string{"read", "get", "update"},
		Struct:        "test",
	}
	response := client.UpdateDataResource(&request)
	fmt.Println(response)
}

// 获取数据资源详情
func TestClient_GetDataResource(t *testing.T) {
	request := dto.GetDataResourceDto{
		ResourceCode:  "stringResourceCode",
		NamespaceCode: "examplePermissionNamespace",
	}
	response := client.GetDataResource(&request)
	fmt.Println(response)
}

// 数据资源列表
func TestClient_ListDataResources(t *testing.T) {
	request := dto.ListDataResourcesDto{
		Page:           1,
		Limit:          10,
		Query:          "stringResourceCode",
		NamespaceCodes: "examplePermissionNamespace,exampleNewPermissionNamespace",
	}
	response := client.ListDataResources(&request)
	fmt.Println(response)
}

// 校验数据资源名称或者 code 是否有效
func TestClient_CheckDataResourceExists(t *testing.T) {
	request := dto.CheckDataResourceExistsDto{
		NamespaceCode: "examplePermissionNamespace",
		ResourceName:  "示例字符串数据资源名称",
	}
	response := client.CheckDataResourceExists(&request)
	fmt.Println(response)

	request1 := dto.CheckDataResourceExistsDto{
		NamespaceCode: "examplePermissionNamespace",
		ResourceCode:  "stringResourceCode2",
	}
	response1 := client.CheckDataResourceExists(&request1)
	fmt.Println(response1)

	request2 := dto.CheckDataResourceExistsDto{
		NamespaceCode: "examplePermissionNamespace",
		ResourceCode:  "stringResourceCode",
	}
	response2 := client.CheckDataResourceExists(&request2)
	fmt.Println(response2)
}

// 创建数据策略
func TestClient_CreateDataPolicy(t *testing.T) {
	request := dto.CreateDataPolicyDto{
		PolicyName: "示例数据策略名称1",
		StatementList: []dto.DataStatementPermissionDto{
			{
				Effect:      "ALLOW",
				Permissions: []string{"examplePermissionNamespace/stringResourceCode/*"},
			},
			{
				Effect:      "ALLOW",
				Permissions: []string{"examplePermissionNamespace/treeResourceCode2/tree1/*"},
			},
		},
		Description: "示例数据策略描述",
	}
	response := client.CreateDataPolicy(&request)
	fmt.Println(response)
}

// 数据策略列表
func TestClient_ListDataPolicies(t *testing.T) {
	request := dto.ListDataPoliciesDto{
		Page:  1,
		Limit: 10,
		Query: "示例数据策略名称",
	}
	response := client.ListDataPolices(&request)
	fmt.Println(response)
}

// 获取数据策略名称
func TestClient_ListSimpleDataPolicies(t *testing.T) {
	request := dto.ListSimpleDataPoliciesDto{
		Page:  1,
		Limit: 10,
		//Query: "examplePolicyName",
	}
	response := client.ListSimpleDataPolices(&request)
	fmt.Println(response)
}

// 获取数据策略详情
func TestClient_GetDataPolicy(t *testing.T) {
	request := dto.GetDataPolicyDto{
		PolicyId: "63a421384c2a336a51f1eecb",
	}
	response := client.GetDataPolicy(&request)
	fmt.Println(response)
}

// 修改数据策略
func TestClient_UpdateDataPolicy(t *testing.T) {
	request := dto.UpdateDataPolicyDto{
		PolicyId:    "63a421384c2a336a51f1eecb",
		PolicyName:  "示例数据策略名称1",
		Description: "示例数据策略描述",
		StatementList: []dto.DataStatementPermissionDto{
			{
				Effect:      "ALLOW",
				Permissions: []string{"examplePermissionNamespace/stringResourceCode/*"},
			},
		},
	}
	response := client.UpdateDataPolicy(&request)
	fmt.Println(response)
}

// 删除数据策略
func TestClient_DeleteDataPolicy(t *testing.T) {
	request := dto.DeleteDataPolicyDto{
		PolicyId: "63a421384c2a336a51f1eecb",
	}
	response := client.DeleteDataPolicy(&request)
	fmt.Println(response)
}

// 校验数据策略名称是否可用
func TestClient_CheckDataPolicyExists(t *testing.T) {
	request := dto.CheckDataPolicyExistsDto{
		PolicyName: "示例数据策略名称",
	}
	response := client.CheckDataPolicyExists(&request)
	fmt.Println(response)
}

// 获取数据策略下所有的授权主体的信息
func TestClient_ListDataPolicyTargets(t *testing.T) {
	request := dto.ListDataPolicyTargetsDto{
		PolicyId: "63a4202d66a7d2af872d0536",
		Page:     1,
		Limit:    10,
		Query:    "4",
		//TargetType: "ROLE",
	}
	response := client.ListDataPolicyTargets(&request)
	fmt.Println(response)
}

// 授权
func TestClient_AuthorizeDataPolicies(t *testing.T) {
	request := dto.CreateAuthorizeDataPolicyDto{
		PolicyIds: []string{"63a4202d66a7d2af872d0536"},
		TargetList: []dto.SubjectDto{
			{
				Id:   "63a43e34dbb5beb923b93ea0",
				Type: "USER",
				Name: "4",
			},
			{
				Id:   "63a43e2f0a6392fc1b4ad325",
				Type: "USER",
				Name: "3",
			},
		},
	}
	response := client.AuthorizeDataPolicies(&request)
	fmt.Println(response)
}

// 取消授权
func TestClient_RevokeDataPolicy(t *testing.T) {
	request := dto.DeleteAuthorizeDataPolicyDto{
		PolicyId:         "63a4202d66a7d2af872d0536",
		TargetIdentifier: "63a43e2f0a6392fc1b4ad325",
		TargetType:       "USER",
	}
	response := client.RevokeDataPolicy(&request)
	fmt.Println(response)
}

// 获取用户的权限
func TestClient_GetUserPermissionList(t *testing.T) {
	request := dto.GetUserPermissionListDto{
		UserIds: []string{"63a43e34dbb5beb923b93ea0"},
		//NamespaceCodes: []string{"examplePermissionNamespace1", "examplePermissionNamespace2"},
	}
	response := client.GetUserPermissionList(&request)
	fmt.Println(response)
}

// 鉴权
func TestClient_CheckPermission(t *testing.T) {
	request := dto.CheckPermissionDto{
		NamespaceCode: "examplePermissionNamespace",
		UserId:        "63a43e34dbb5beb923b93ea0",
		Action:        "read",
		Resources:     []string{"stringResourceCode", "arrayResourceCode1", "/treeResourceCode3/tree1111/resourceStructChildrenCode"},
	}
	response := client.CheckPermission(&request)
	fmt.Println(response)
}

func TestClient_GetUserResourcePermissionList(t *testing.T) {
	request := dto.GetUserResourcePermissionListDto{
		NamespaceCode: "examplePermissionNamespace",
		UserId:        "63a43e34dbb5beb923b93ea0",
		//Resources:     []string{"strResourceCode", "arrayResourceCode", "/treeResourceCode/structCode/resourceStructChildrenCode"},
		Resources: []string{"stringResourceCode"},
	}
	response := client.GetUserResourcePermissionList(&request)
	fmt.Println(response)
}

func TestClient_ListResourceTargets(t *testing.T) {
	request := dto.ListResourceTargetsDto{
		NamespaceCode: "examplePermissionNamespace",
		Actions:       []string{"read", "get"},
		Resources:     []string{"stringResourceCode", "arrayResourceCode", "/treeResourceCode/structCode/resourceStructChildrenCode"},
	}
	response := client.ListResourceTargets(&request)
	fmt.Println(response)
}

func TestClient_CheckUserSameLevelPermission(t *testing.T) {
	request := dto.CheckUserSameLevelPermissionDto{
		NamespaceCode: "examplePermissionNamespace",
		UserId:        "63a43e34dbb5beb923b93ea0",
		Action:        "geta",
		Resource:      "stringResourceCode",
	}
	response := client.CheckUserSameLevelPermission(&request)
	fmt.Println(response)
}
