package management

import (
	"encoding/json"
	"fmt"
	"github.com/Authing/authing-golang-sdk/dto"
	"github.com/valyala/fasthttp"
)

/*
 * @summary 获取 Management API Token
 * @description 通过 AccessKey ID 与 AccessKey Secret 获取 Management API Token，此 Token 可以用来操作 Management API。
 * @param requestBody
 * @returns GetManagementTokenRespDto
 */
func (c *Client) GetManagementToken(reqDto *dto.GetManagementAccessTokenDto) *dto.GetManagementTokenRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-management-token", fasthttp.MethodPost, reqDto)
	var response dto.GetManagementTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户信息
 * @description 通过用户 ID，获取用户详情，可以选择获取自定义数据、identities、选择指定用户 ID 类型等。
 * @param userId 用户 ID
 * @param userIdType 用户 ID 类型，可以指定为用户 ID、手机号、邮箱、用户名和 externalId。
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @returns UserSingleRespDto
 */
func (c *Client) GetUser(reqDto *dto.GetUserDto) *dto.UserSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user", fasthttp.MethodGet, reqDto)
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量获取用户信息
 * @description 通过用户 ID 列表，批量获取用户信息，可以选择获取自定义数据、identities、选择指定用户 ID 类型等。
 * @param userIds 用户 ID 数组
 * @param userIdType 用户 ID 类型，可以指定为用户 ID、手机号、邮箱、用户名和 externalId。
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @returns UserListRespDto
 */
func (c *Client) GetUserBatch(reqDto *dto.GetUserBatchDto) *dto.UserListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-batch", fasthttp.MethodGet, reqDto)
	var response dto.UserListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户列表
 * @description 获取用户列表接口，支持分页，可以选择获取自定义数据、identities 等。
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param status 账户当前状态，如 已停用、已离职、正常状态、已归档
 * @param updatedAtStart 用户创建、修改开始时间，为精确到秒的 UNIX 时间戳；支持获取从某一段时间之后的增量数据
 * @param updatedAtEnd 用户创建、修改终止时间，为精确到秒的 UNIX 时间戳；支持获取某一段时间内的增量数据。默认为当前时间
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @returns UserPaginatedRespDto
 */
func (c *Client) ListUsers(reqDto *dto.ListUsersDto) *dto.UserPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-users", fasthttp.MethodGet, reqDto)
	var response dto.UserPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户的外部身份源
 * @description 通过用户 ID，获取用户的外部身份源、选择指定用户 ID 类型。
 * @param userId 用户 ID
 * @param userIdType 用户 ID 类型，可以指定为用户 ID、手机号、邮箱、用户名和 externalId。
 * @returns IdentityListRespDto
 */
func (c *Client) GetUserIdentities(reqDto *dto.GetUserIdentitiesDto) *dto.IdentityListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-identities", fasthttp.MethodGet, reqDto)
	var response dto.IdentityListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户角色列表
 * @description 通过用户 ID，获取用户角色列表，可以选择所属权限分组 code、选择指定用户 ID 类型等。
 * @param userId 用户 ID
 * @param userIdType 用户 ID 类型，可以指定为用户 ID、手机号、邮箱、用户名和 externalId。
 * @param namespace 所属权限分组的 code
 * @returns RolePaginatedRespDto
 */
func (c *Client) GetUserRoles(reqDto *dto.GetUserRolesDto) *dto.RolePaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-roles", fasthttp.MethodGet, reqDto)
	var response dto.RolePaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户实名认证信息
 * @description 通过用户 ID，获取用户实名认证信息，可以选择指定用户 ID 类型。
 * @param userId 用户 ID
 * @param userIdType 用户 ID 类型，可以指定为用户 ID、手机号、邮箱、用户名和 externalId。
 * @returns PrincipalAuthenticationInfoPaginatedRespDto
 */
func (c *Client) GetUserPrincipalAuthenticationInfo(reqDto *dto.GetUserPrincipalAuthenticationInfoDto) *dto.PrincipalAuthenticationInfoPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-principal-authentication-info", fasthttp.MethodGet, reqDto)
	var response dto.PrincipalAuthenticationInfoPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除用户实名认证信息
 * @description 通过用户 ID，删除用户实名认证信息，可以选择指定用户 ID 类型等。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) ResetUserPrincipalAuthenticationInfo(reqDto *dto.ResetUserPrincipalAuthenticationInfoDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/reset-user-principal-authentication-info", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户部门列表
 * @description 通过用户 ID，获取用户部门列表，支持分页，可以选择获取自定义数据、选择指定用户 ID 类型、增序或降序等。
 * @param userId 用户 ID
 * @param userIdType 用户 ID 类型，可以指定为用户 ID、手机号、邮箱、用户名和 externalId。
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param withCustomData 是否获取自定义数据
 * @param sortBy 排序依据，如 部门创建时间、加入部门时间、部门名称、部门标志符
 * @param orderBy 增序或降序
 * @returns UserDepartmentPaginatedRespDto
 */
func (c *Client) GetUserDepartments(reqDto *dto.GetUserDepartmentsDto) *dto.UserDepartmentPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-departments", fasthttp.MethodGet, reqDto)
	var response dto.UserDepartmentPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 设置用户所在部门
 * @description 通过用户 ID，设置用户所在部门，可以选择指定用户 ID 类型等。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) SetUserDepartments(reqDto *dto.SetUserDepartmentsDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/set-user-departments", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户分组列表
 * @description 通过用户 ID，获取用户分组列表，可以选择指定用户 ID 类型等。
 * @param userId 用户 ID
 * @param userIdType 用户 ID 类型，可以指定为用户 ID、手机号、邮箱、用户名和 externalId。
 * @returns GroupPaginatedRespDto
 */
func (c *Client) GetUserGroups(reqDto *dto.GetUserGroupsDto) *dto.GroupPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-groups", fasthttp.MethodGet, reqDto)
	var response dto.GroupPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除用户
 * @description 通过用户 ID 列表，删除用户，支持批量删除，可以选择指定用户 ID 类型等。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) DeleteUsersBatch(reqDto *dto.DeleteUsersBatchDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-users-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户 MFA 绑定信息
 * @description 通过用户 ID，获取用户 MFA 绑定信息，可以选择指定用户 ID 类型等。
 * @param userId 用户 ID
 * @param userIdType 用户 ID 类型，可以指定为用户 ID、手机号、邮箱、用户名和 externalId。
 * @returns UserMfaSingleRespDto
 */
func (c *Client) GetUserMfaInfo(reqDto *dto.GetUserMfaInfoDto) *dto.UserMfaSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-mfa-info", fasthttp.MethodGet, reqDto)
	var response dto.UserMfaSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取已归档的用户列表
 * @description 获取已归档的用户列表，支持分页，可以筛选开始时间等。
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param startAt 开始时间，为精确到秒的 UNIX 时间戳，默认不指定
 * @returns ListArchivedUsersSingleRespDto
 */
func (c *Client) ListArchivedUsers(reqDto *dto.ListArchivedUsersDto) *dto.ListArchivedUsersSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-archived-users", fasthttp.MethodGet, reqDto)
	var response dto.ListArchivedUsersSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 强制下线用户
 * @description 通过用户 ID、App ID 列表，强制让用户下线，可以选择指定用户 ID 类型等。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) KickUsers(reqDto *dto.KickUsersDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/kick-users", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 判断用户是否存在
 * @description 根据条件判断用户是否存在，可以筛选用户名、邮箱、手机号、第三方外部 ID 等。
 * @param requestBody
 * @returns IsUserExistsRespDto
 */
func (c *Client) IsUserExists(reqDto *dto.IsUserExistsReqDto) *dto.IsUserExistsRespDto {
	b, err := c.SendHttpRequest("/api/v3/is-user-exists", fasthttp.MethodPost, reqDto)
	var response dto.IsUserExistsRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建用户
 * @description 创建用户，邮箱、手机号、用户名必须包含其中一个，邮箱、手机号、用户名、externalId 用户池内唯一，此接口将以管理员身份创建用户因此不需要进行手机号验证码检验等安全检测。
 * @param requestBody
 * @returns UserSingleRespDto
 */
func (c *Client) CreateUser(reqDto *dto.CreateUserReqDto) *dto.UserSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-user", fasthttp.MethodPost, reqDto)
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量创建用户
 * @description 批量创建用户，邮箱、手机号、用户名必须包含其中一个，邮箱、手机号、用户名、externalId 用户池内唯一，此接口将以管理员身份批量创建用户因此不需要进行手机号验证码检验等安全检测。
 * @param requestBody
 * @returns UserListRespDto
 */
func (c *Client) CreateUsersBatch(reqDto *dto.CreateUserBatchReqDto) *dto.UserListRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-users-batch", fasthttp.MethodPost, reqDto)
	var response dto.UserListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 修改用户资料
 * @description 通过用户 ID，修改用户资料，邮箱、手机号、用户名、externalId 用户池内唯一，此接口将以管理员身份修改用户资料因此不需要进行手机号验证码检验等安全检测。
 * @param requestBody
 * @returns UserSingleRespDto
 */
func (c *Client) UpdateUser(reqDto *dto.UpdateUserReqDto) *dto.UserSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-user", fasthttp.MethodPost, reqDto)
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户可访问的应用
 * @description 通过用户 ID，获取用户可访问的应用，可以选择指定用户 ID 类型等。
 * @param userId 用户 ID
 * @param userIdType 用户 ID 类型，可以指定为用户 ID、手机号、邮箱、用户名和 externalId。
 * @returns AppListRespDto
 */
func (c *Client) GetUserAccessibleApps(reqDto *dto.GetUserAccessibleAppsDto) *dto.AppListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-accessible-apps", fasthttp.MethodGet, reqDto)
	var response dto.AppListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户授权的应用
 * @description 通过用户 ID，获取用户授权的应用，可以选择指定用户 ID 类型等。
 * @param userId 用户 ID
 * @param userIdType 用户 ID 类型，可以指定为用户 ID、手机号、邮箱、用户名和 externalId。
 * @returns AppListRespDto
 */
func (c *Client) GetUserAuthorizedApps(reqDto *dto.GetUserAuthorizedAppsDto) *dto.AppListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-authorized-apps", fasthttp.MethodGet, reqDto)
	var response dto.AppListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 判断用户是否有某个角色
 * @description 通过用户 ID，判断用户是否有某个角色，支持传入多个角色，可以选择指定用户 ID 类型等。
 * @param requestBody
 * @returns HasAnyRoleRespDto
 */
func (c *Client) HasAnyRole(reqDto *dto.HasAnyRoleReqDto) *dto.HasAnyRoleRespDto {
	b, err := c.SendHttpRequest("/api/v3/has-any-role", fasthttp.MethodPost, reqDto)
	var response dto.HasAnyRoleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户的登录历史记录
 * @description 通过用户 ID，获取用户登录历史记录，支持分页，可以选择指定用户 ID 类型、应用 ID、开始与结束时间戳等。
 * @param userId 用户 ID
 * @param userIdType 用户 ID 类型，可以指定为用户 ID、手机号、邮箱、用户名和 externalId。
 * @param appId 应用 ID
 * @param clientIp 客户端 IP
 * @param start 开始时间戳（毫秒）
 * @param end 结束时间戳（毫秒）
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns UserLoginHistoryPaginatedRespDto
 */
func (c *Client) GetUserLoginHistory(reqDto *dto.GetUserLoginHistoryDto) *dto.UserLoginHistoryPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-login-history", fasthttp.MethodGet, reqDto)
	var response dto.UserLoginHistoryPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 通过用户 ID，获取用户曾经登录过的应用，可以选择指定用户 ID 类型等。
 * @description 获取用户曾经登录过的应用
 * @param userId 用户 ID
 * @param userIdType 用户 ID 类型，可以指定为用户 ID、手机号、邮箱、用户名和 externalId。
 * @returns UserLoggedInAppsListRespDto
 */
func (c *Client) GetUserLoggedinApps(reqDto *dto.GetUserLoggedinAppsDto) *dto.UserLoggedInAppsListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-loggedin-apps", fasthttp.MethodGet, reqDto)
	var response dto.UserLoggedInAppsListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 通过用户 ID，获取用户曾经登录过的身份源，可以选择指定用户 ID 类型等。
 * @description 获取用户曾经登录过的身份源
 * @param userId 用户 ID
 * @param userIdType 用户 ID 类型，可以指定为用户 ID、手机号、邮箱、用户名和 externalId。
 * @returns UserLoggedInIdentitiesRespDto
 */
func (c *Client) GetUserLoggedinIdentities(reqDto *dto.GetUserLoggedInIdentitiesDto) *dto.UserLoggedInIdentitiesRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-logged-in-identities", fasthttp.MethodGet, reqDto)
	var response dto.UserLoggedInIdentitiesRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户被授权的所有资源
 * @description 通过用户 ID，获取用户被授权的所有资源，可以选择指定用户 ID 类型等，用户被授权的资源是用户自身被授予、通过分组继承、通过角色继承、通过组织机构继承的集合。
 * @param userId 用户 ID
 * @param userIdType 用户 ID 类型，可以指定为用户 ID、手机号、邮箱、用户名和 externalId。
 * @param namespace 所属权限分组的 code
 * @param resourceType 资源类型，如 数据、API、菜单、按钮
 * @returns AuthorizedResourcePaginatedRespDto
 */
func (c *Client) GetUserAuthorizedResources(reqDto *dto.GetUserAuthorizedResourcesDto) *dto.AuthorizedResourcePaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-authorized-resources", fasthttp.MethodGet, reqDto)
	var response dto.AuthorizedResourcePaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取分组详情
 * @description 通过分组 code，获取分组详情。
 * @param code 分组 code
 * @returns GroupSingleRespDto
 */
func (c *Client) GetGroup(reqDto *dto.GetGroupDto) *dto.GroupSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-group", fasthttp.MethodGet, reqDto)
	var response dto.GroupSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取分组列表
 * @description 获取分组列表，支持分页。
 * @param keywords 搜索分组 code 或分组名称
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns GroupPaginatedRespDto
 */
func (c *Client) ListGroups(reqDto *dto.ListGroupsDto) *dto.GroupPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-groups", fasthttp.MethodGet, reqDto)
	var response dto.GroupPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建分组
 * @description 创建分组，一个分组必须包含分组名称与唯一标志符 code，且必须为一个合法的英文标志符，如 developers。
 * @param requestBody
 * @returns GroupSingleRespDto
 */
func (c *Client) CreateGroup(reqDto *dto.CreateGroupReqDto) *dto.GroupSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-group", fasthttp.MethodPost, reqDto)
	var response dto.GroupSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量创建分组
 * @description 批量创建分组，一个分组必须包含分组名称与唯一标志符 code，且必须为一个合法的英文标志符，如 developers。
 * @param requestBody
 * @returns GroupListRespDto
 */
func (c *Client) CreateGroupsBatch(reqDto *dto.CreateGroupBatchReqDto) *dto.GroupListRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-groups-batch", fasthttp.MethodPost, reqDto)
	var response dto.GroupListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 修改分组
 * @description 通过分组 code，修改分组，可以修改此分组的 code。
 * @param requestBody
 * @returns GroupSingleRespDto
 */
func (c *Client) UpdateGroup(reqDto *dto.UpdateGroupReqDto) *dto.GroupSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-group", fasthttp.MethodPost, reqDto)
	var response dto.GroupSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量删除分组
 * @description 通过分组 code，批量删除分组。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) DeleteGroupsBatch(reqDto *dto.DeleteGroupsReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-groups-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 添加分组成员
 * @description 添加分组成员，成员以用户 ID 数组形式传递。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) AddGroupMembers(reqDto *dto.AddGroupMembersReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/add-group-members", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量移除分组成员
 * @description 批量移除分组成员，成员以用户 ID 数组形式传递。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) RemoveGroupMembers(reqDto *dto.RemoveGroupMembersReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/remove-group-members", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取分组成员列表
 * @description 通过分组 code，获取分组成员列表，支持分页，可以获取自定义数据、identities、部门 ID 列表。
 * @param code 分组 code
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @returns UserPaginatedRespDto
 */
func (c *Client) ListGroupMembers(reqDto *dto.ListGroupMembersDto) *dto.UserPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-group-members", fasthttp.MethodGet, reqDto)
	var response dto.UserPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取分组被授权的资源列表
 * @description 通过分组 code，获取分组被授权的资源列表，可以通过资源类型、权限分组 code 筛选。
 * @param code 分组 code
 * @param namespace 所属权限分组的 code
 * @param resourceType 资源类型
 * @returns AuthorizedResourceListRespDto
 */
func (c *Client) GetGroupAuthorizedResources(reqDto *dto.GetGroupAuthorizedResourcesDto) *dto.AuthorizedResourceListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-group-authorized-resources", fasthttp.MethodGet, reqDto)
	var response dto.AuthorizedResourceListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取角色详情
 * @description 通过权限分组内角色 code，获取角色详情。
 * @param code 权限分组内角色的唯一标识符
 * @param namespace 所属权限分组的 code
 * @returns RoleSingleRespDto
 */
func (c *Client) GetRole(reqDto *dto.GetRoleDto) *dto.RoleSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-role", fasthttp.MethodGet, reqDto)
	var response dto.RoleSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 分配角色
 * @description 通过权限分组内角色 code，分配角色，被分配者可以是用户或部门。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) AssignRole(reqDto *dto.AssignRoleDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/assign-role", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 移除分配的角色
 * @description 通过权限分组内角色 code，移除分配的角色，被分配者可以是用户或部门。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) RevokeRole(reqDto *dto.RevokeRoleDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/revoke-role", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取角色被授权的资源列表
 * @description 通过权限分组内角色 code，获取角色被授权的资源列表。
 * @param code 权限分组内角色的唯一标识符
 * @param namespace 所属权限分组的 code
 * @param resourceType 资源类型，如 数据、API、按钮、菜单
 * @returns RoleAuthorizedResourcePaginatedRespDto
 */
func (c *Client) GetRoleAuthorizedResources(reqDto *dto.GetRoleAuthorizedResourcesDto) *dto.RoleAuthorizedResourcePaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-role-authorized-resources", fasthttp.MethodGet, reqDto)
	var response dto.RoleAuthorizedResourcePaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取角色成员列表
 * @description 通过权限分组内内角色 code，获取角色成员列表，支持分页，可以选择或获取自定义数据、identities 等。
 * @param code 权限分组内角色的唯一标识符
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @param namespace 所属权限分组的 code
 * @returns UserPaginatedRespDto
 */
func (c *Client) ListRoleMembers(reqDto *dto.ListRoleMembersDto) *dto.UserPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-role-members", fasthttp.MethodGet, reqDto)
	var response dto.UserPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取角色的部门列表
 * @description 通过权限分组内角色 code，获取角色的部门列表，支持分页。
 * @param code 权限分组内角色的唯一标识符
 * @param namespace 所属权限分组的 code
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns RoleDepartmentListPaginatedRespDto
 */
func (c *Client) ListRoleDepartments(reqDto *dto.ListRoleDepartmentsDto) *dto.RoleDepartmentListPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-role-departments", fasthttp.MethodGet, reqDto)
	var response dto.RoleDepartmentListPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建角色
 * @description 通过权限分组内角色 code，创建角色，可以选择权限分组、角色描述等。
 * @param requestBody
 * @returns RoleSingleRespDto
 */
func (c *Client) CreateRole(reqDto *dto.CreateRoleDto) *dto.RoleSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-role", fasthttp.MethodPost, reqDto)
	var response dto.RoleSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取角色列表
 * @description 获取角色列表，支持分页。
 * @param keywords 搜索角色 code
 * @param namespace 所属权限分组的 code
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns RolePaginatedRespDto
 */
func (c *Client) ListRoles(reqDto *dto.ListRolesDto) *dto.RolePaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-roles", fasthttp.MethodGet, reqDto)
	var response dto.RolePaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除角色
 * @description 删除角色，可以批量删除。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) DeleteRolesBatch(reqDto *dto.DeleteRoleDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-roles-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量创建角色
 * @description 批量创建角色，可以选择权限分组、角色描述等。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) CreateRolesBatch(reqDto *dto.CreateRolesBatch) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-roles-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 修改角色
 * @description 通过权限分组内角色新旧 code，修改角色，可以选择角色描述等。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) UpdateRole(reqDto *dto.UpdateRoleDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-role", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取顶层组织机构列表
 * @description 获取顶层组织机构列表，支持分页。
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param fetchAll 拉取所有
 * @returns OrganizationPaginatedRespDto
 */
func (c *Client) ListOrganizations(reqDto *dto.ListOrganizationsDto) *dto.OrganizationPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-organizations", fasthttp.MethodGet, reqDto)
	var response dto.OrganizationPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建顶层组织机构
 * @description 创建组织机构，会创建一个只有一个节点的组织机构，可以选择组织描述信息、根节点自定义 ID、多语言等。
 * @param requestBody
 * @returns OrganizationSingleRespDto
 */
func (c *Client) CreateOrganization(reqDto *dto.CreateOrganizationReqDto) *dto.OrganizationSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-organization", fasthttp.MethodPost, reqDto)
	var response dto.OrganizationSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 修改顶层组织机构
 * @description 通过组织 code，修改顶层组织机构，可以选择部门描述、新组织 code、组织名称等。
 * @param requestBody
 * @returns OrganizationSingleRespDto
 */
func (c *Client) UpdateOrganization(reqDto *dto.UpdateOrganizationReqDto) *dto.OrganizationSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-organization", fasthttp.MethodPost, reqDto)
	var response dto.OrganizationSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除组织机构
 * @description 通过组织 code，删除组织机构树。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) DeleteOrganization(reqDto *dto.DeleteOrganizationReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-organization", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 搜索顶层组织机构列表
 * @description 通过搜索关键词，搜索顶层组织机构列表，支持分页。
 * @param keywords 搜索关键词，如组织机构名称
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns OrganizationPaginatedRespDto
 */
func (c *Client) SearchOrganizations(reqDto *dto.SearchOrganizationsDto) *dto.OrganizationPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/search-organizations", fasthttp.MethodGet, reqDto)
	var response dto.OrganizationPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取部门信息
 * @description 通过组织 code 以及 部门 ID 或 部门 code，获取部门信息，可以获取自定义数据。
 * @param organizationCode 组织 code
 * @param departmentId 部门 ID，根部门传 `root`。departmentId 和 departmentCode 必传其一。
 * @param departmentCode 部门 code。departmentId 和 departmentCode 必传其一。
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @param withCustomData 是否获取自定义数据
 * @returns DepartmentSingleRespDto
 */
func (c *Client) GetDepartment(reqDto *dto.GetDepartmentDto) *dto.DepartmentSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-department", fasthttp.MethodGet, reqDto)
	var response dto.DepartmentSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建部门
 * @description 通过组织 code、部门名称、父部门 ID，创建部门，可以设置多种参数。
 * @param requestBody
 * @returns DepartmentSingleRespDto
 */
func (c *Client) CreateDepartment(reqDto *dto.CreateDepartmentReqDto) *dto.DepartmentSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-department", fasthttp.MethodPost, reqDto)
	var response dto.DepartmentSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 修改部门
 * @description 通过组织 code、部门 ID，修改部门，可以设置多种参数。
 * @param requestBody
 * @returns DepartmentSingleRespDto
 */
func (c *Client) UpdateDepartment(reqDto *dto.UpdateDepartmentReqDto) *dto.DepartmentSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-department", fasthttp.MethodPost, reqDto)
	var response dto.DepartmentSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除部门
 * @description 通过组织 code、部门 ID，删除部门。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) DeleteDepartment(reqDto *dto.DeleteDepartmentReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-department", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 搜索部门
 * @description 通过组织 code、搜索关键词，搜索部门，可以搜索组织名称等。
 * @param requestBody
 * @returns DepartmentListRespDto
 */
func (c *Client) SearchDepartments(reqDto *dto.SearchDepartmentsReqDto) *dto.DepartmentListRespDto {
	b, err := c.SendHttpRequest("/api/v3/search-departments", fasthttp.MethodPost, reqDto)
	var response dto.DepartmentListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取子部门列表
 * @description 通过组织 code、部门 ID，获取子部门列表，可以选择获取自定义数据、虚拟组织等。
 * @param organizationCode 组织 code
 * @param departmentId 需要获取的部门 ID
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @param excludeVirtualNode 是否要排除虚拟组织
 * @param onlyVirtualNode 是否只包含虚拟组织
 * @param withCustomData 是否获取自定义数据
 * @returns DepartmentPaginatedRespDto
 */
func (c *Client) ListChildrenDepartments(reqDto *dto.ListChildrenDepartmentsDto) *dto.DepartmentPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-children-departments", fasthttp.MethodGet, reqDto)
	var response dto.DepartmentPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取部门成员列表
 * @description 通过组织 code、部门 ID、排序，获取部门成员列表，支持分页，可以选择获取自定义数据、identities 等。
 * @param organizationCode 组织 code
 * @param departmentId 部门 ID，根部门传 `root`
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @param includeChildrenDepartments 是否包含子部门的成员
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @param sortBy 排序依据
 * @param orderBy 增序还是倒序
 * @returns UserPaginatedRespDto
 */
func (c *Client) ListDepartmentMembers(reqDto *dto.ListDepartmentMembersDto) *dto.UserPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-department-members", fasthttp.MethodGet, reqDto)
	var response dto.UserPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取部门直属成员 ID 列表
 * @description 通过组织 code、部门 ID，获取部门直属成员 ID 列表。
 * @param organizationCode 组织 code
 * @param departmentId 部门 ID，根部门传 `root`
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @returns UserIdListRespDto
 */
func (c *Client) ListDepartmentMemberIds(reqDto *dto.ListDepartmentMemberIdsDto) *dto.UserIdListRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-department-member-ids", fasthttp.MethodGet, reqDto)
	var response dto.UserIdListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 搜索部门下的成员
 * @description 通过组织 code、部门 ID、搜索关键词，搜索部门下的成员，支持分页，可以选择获取自定义数据、identities 等。
 * @param organizationCode 组织 code
 * @param departmentId 部门 ID，根部门传 `root`
 * @param keywords 搜索关键词，如成员名称
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @param includeChildrenDepartments 是否包含子部门的成员
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @returns UserPaginatedRespDto
 */
func (c *Client) SearchDepartmentMembers(reqDto *dto.SearchDepartmentMembersDto) *dto.UserPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/search-department-members", fasthttp.MethodGet, reqDto)
	var response dto.UserPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 部门下添加成员
 * @description 通过部门 ID、组织 code，添加部门下成员。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) AddDepartmentMembers(reqDto *dto.AddDepartmentMembersReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/add-department-members", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 部门下删除成员
 * @description 通过部门 ID、组织 code，删除部门下成员。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) RemoveDepartmentMembers(reqDto *dto.RemoveDepartmentMembersReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/remove-department-members", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取父部门信息
 * @description 通过组织 code、部门 ID，获取父部门信息，可以选择获取自定义数据等。
 * @param organizationCode 组织 code
 * @param departmentId 部门 ID
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @param withCustomData 是否获取自定义数据
 * @returns DepartmentSingleRespDto
 */
func (c *Client) GetParentDepartment(reqDto *dto.GetParentDepartmentDto) *dto.DepartmentSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-parent-department", fasthttp.MethodGet, reqDto)
	var response dto.DepartmentSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 判断用户是否在某个部门下
 * @description 通过组织 code、部门 ID，判断用户是否在某个部门下，可以选择包含子部门。
 * @param userId 用户 ID
 * @param organizationCode 组织 code
 * @param departmentId 部门 ID，根部门传 `root`。departmentId 和 departmentCode 必传其一。
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @param includeChildrenDepartments 是否包含子部门
 * @returns IsUserInDepartmentRespDto
 */
func (c *Client) IsUserInDepartment(reqDto *dto.IsUserInDepartmentDto) *dto.IsUserInDepartmentRespDto {
	b, err := c.SendHttpRequest("/api/v3/is-user-in-department", fasthttp.MethodGet, reqDto)
	var response dto.IsUserInDepartmentRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取身份源列表
 * @description 获取身份源列表，可以指定 租户 ID 筛选。
 * @param tenantId 租户 ID
 * @returns ExtIdpListPaginatedRespDto
 */
func (c *Client) ListExtIdp(reqDto *dto.ListExtIdpDto) *dto.ExtIdpListPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-ext-idp", fasthttp.MethodGet, reqDto)
	var response dto.ExtIdpListPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取身份源详情
 * @description 通过 身份源 ID，获取身份源详情，可以指定 租户 ID 筛选。
 * @param id 身份源 ID
 * @param tenantId 租户 ID
 * @returns ExtIdpDetailSingleRespDto
 */
func (c *Client) GetExtIdp(reqDto *dto.GetExtIdpDto) *dto.ExtIdpDetailSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-ext-idp", fasthttp.MethodGet, reqDto)
	var response dto.ExtIdpDetailSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建身份源
 * @description 创建身份源，可以设置身份源名称、连接类型、租户 ID 等。
 * @param requestBody
 * @returns ExtIdpSingleRespDto
 */
func (c *Client) CreateExtIdp(reqDto *dto.CreateExtIdpDto) *dto.ExtIdpSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-ext-idp", fasthttp.MethodPost, reqDto)
	var response dto.ExtIdpSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 更新身份源配置
 * @description 更新身份源配置，可以设置身份源 ID 与 名称。
 * @param requestBody
 * @returns ExtIdpSingleRespDto
 */
func (c *Client) UpdateExtIdp(reqDto *dto.UpdateExtIdpDto) *dto.ExtIdpSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-ext-idp", fasthttp.MethodPost, reqDto)
	var response dto.ExtIdpSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除身份源
 * @description 通过身份源 ID，删除身份源。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) DeleteExtIdp(reqDto *dto.DeleteExtIdpDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-ext-idp", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 在某个已有身份源下创建新连接
 * @description 在某个已有身份源下创建新连接，可以设置身份源图标、是否只支持登录等。
 * @param requestBody
 * @returns ExtIdpConnDetailSingleRespDto
 */
func (c *Client) CreateExtIdpConn(reqDto *dto.CreateExtIdpConnDto) *dto.ExtIdpConnDetailSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-ext-idp-conn", fasthttp.MethodPost, reqDto)
	var response dto.ExtIdpConnDetailSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 更新身份源连接
 * @description 更新身份源连接，可以设置身份源图标、是否只支持登录等。
 * @param requestBody
 * @returns ExtIdpConnDetailSingleRespDto
 */
func (c *Client) UpdateExtIdpConn(reqDto *dto.UpdateExtIdpConnDto) *dto.ExtIdpConnDetailSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-ext-idp-conn", fasthttp.MethodPost, reqDto)
	var response dto.ExtIdpConnDetailSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除身份源连接
 * @description 通过身份源连接 ID，删除身份源连接。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) DeleteExtIdpConn(reqDto *dto.DeleteExtIdpConnDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-ext-idp-conn", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 身份源连接开关
 * @description 身份源连接开关，可以打开或关闭身份源连接。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) ChangeConnState(reqDto *dto.EnableExtIdpConnDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/enable-ext-idp-conn", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户池配置的自定义字段列表
 * @description 通过主体类型，获取用户池配置的自定义字段列表。
 * @param targetType 主体类型，目前支持用户、角色、分组、部门
 * @returns CustomFieldListRespDto
 */
func (c *Client) GetCustomFields(reqDto *dto.GetCustomFieldsDto) *dto.CustomFieldListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-custom-fields", fasthttp.MethodGet, reqDto)
	var response dto.CustomFieldListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建/修改自定义字段定义
 * @description 创建/修改自定义字段定义，如果传入的 key 不存在则创建，存在则更新。
 * @param requestBody
 * @returns CustomFieldListRespDto
 */
func (c *Client) SetCustomFields(reqDto *dto.SetCustomFieldsReqDto) *dto.CustomFieldListRespDto {
	b, err := c.SendHttpRequest("/api/v3/set-custom-fields", fasthttp.MethodPost, reqDto)
	var response dto.CustomFieldListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 设置自定义字段的值
 * @description 给用户、角色、部门设置自定义字段的值，如果存在则更新，不存在则创建。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) SetCustomData(reqDto *dto.SetCustomDataReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/set-custom-data", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户、分组、角色、组织机构的自定义字段值
 * @description 通过筛选条件，获取用户、分组、角色、组织机构的自定义字段值。
 * @param targetType 主体类型，目前支持用户、角色、分组、部门
 * @param targetIdentifier 目标对象唯一标志符
 * @param namespace 所属权限分组的 code，当 targetType 为角色的时候需要填写，否则可以忽略
 * @returns GetCustomDataRespDto
 */
func (c *Client) GetCustomData(reqDto *dto.GetCustomDataDto) *dto.GetCustomDataRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-custom-data", fasthttp.MethodGet, reqDto)
	var response dto.GetCustomDataRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建资源
 * @description 创建资源，可以设置资源的描述、定义的操作类型、URL 标识等。
 * @param requestBody
 * @returns ResourceRespDto
 */
func (c *Client) CreateResource(reqDto *dto.CreateResourceDto) *dto.ResourceRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-resource", fasthttp.MethodPost, reqDto)
	var response dto.ResourceRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量创建资源
 * @description 批量创建资源，可以设置资源的描述、定义的操作类型、URL 标识等。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) CreateResourcesBatch(reqDto *dto.CreateResourcesBatchDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-resources-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取资源详情
 * @description 根据筛选条件，获取资源详情。
 * @param code 资源唯一标志符
 * @param namespace 所属权限分组的 code
 * @returns ResourceRespDto
 */
func (c *Client) GetResource(reqDto *dto.GetResourceDto) *dto.ResourceRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-resource", fasthttp.MethodGet, reqDto)
	var response dto.ResourceRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量获取资源详情
 * @description 根据筛选条件，批量获取资源详情。
 * @param codeList 资源 code 列表，批量可以使用逗号分隔
 * @param namespace 所属权限分组的 code
 * @returns ResourceListRespDto
 */
func (c *Client) GetResourcesBatch(reqDto *dto.GetResourcesBatchDto) *dto.ResourceListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-resources-batch", fasthttp.MethodGet, reqDto)
	var response dto.ResourceListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 分页获取资源列表
 * @description 根据筛选条件，分页获取资源详情列表。
 * @param namespace 所属权限分组的 code
 * @param type 资源类型
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns ResourcePaginatedRespDto
 */
func (c *Client) ListResources(reqDto *dto.ListResourcesDto) *dto.ResourcePaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-resources", fasthttp.MethodGet, reqDto)
	var response dto.ResourcePaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 修改资源
 * @description 修改资源，可以设置资源的描述、定义的操作类型、URL 标识等。
 * @param requestBody
 * @returns ResourceRespDto
 */
func (c *Client) UpdateResource(reqDto *dto.UpdateResourceDto) *dto.ResourceRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-resource", fasthttp.MethodPost, reqDto)
	var response dto.ResourceRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除资源
 * @description 通过资源唯一标志符以及所属权限分组，删除资源。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) DeleteResource(reqDto *dto.DeleteResourceDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-resource", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量删除资源
 * @description 通过资源唯一标志符以及所属权限分组，批量删除资源
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) DeleteResourcesBatch(reqDto *dto.DeleteResourcesBatchDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-resources-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建权限分组
 * @description 创建权限分组，可以设置分组名称与描述信息。
 * @param requestBody
 * @returns NamespaceRespDto
 */
func (c *Client) CreateNamespace(reqDto *dto.CreateNamespaceDto) *dto.NamespaceRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-namespace", fasthttp.MethodPost, reqDto)
	var response dto.NamespaceRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量创建权限分组
 * @description 批量创建权限分组，可以分别设置分组名称与描述信息。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) CreateNamespacesBatch(reqDto *dto.CreateNamespacesBatchDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-namespaces-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取权限分组详情
 * @description 通过权限分组唯一标志符，获取权限分组详情。
 * @param code 权限分组唯一标志符
 * @returns NamespaceRespDto
 */
func (c *Client) GetNamespace(reqDto *dto.GetNamespaceDto) *dto.NamespaceRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-namespace", fasthttp.MethodGet, reqDto)
	var response dto.NamespaceRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量获取权限分组详情
 * @description 分别通过权限分组唯一标志符，批量获取权限分组详情。
 * @param codeList 资源 code 列表，批量可以使用逗号分隔
 * @returns NamespaceListRespDto
 */
func (c *Client) GetNamespacesBatch(reqDto *dto.GetNamespacesBatchDto) *dto.NamespaceListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-namespaces-batch", fasthttp.MethodGet, reqDto)
	var response dto.NamespaceListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 修改权限分组信息
 * @description 修改权限分组信息，可以修改名称、描述信息以及新的唯一标志符。
 * @param requestBody
 * @returns UpdateNamespaceRespDto
 */
func (c *Client) UpdateNamespace(reqDto *dto.UpdateNamespaceDto) *dto.UpdateNamespaceRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-namespace", fasthttp.MethodPost, reqDto)
	var response dto.UpdateNamespaceRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除权限分组信息
 * @description 通过权限分组唯一标志符，删除权限分组信息。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) DeleteNamespace(reqDto *dto.DeleteNamespaceDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-namespace", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量删除权限分组
 * @description 分别通过权限分组唯一标志符，批量删除权限分组。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) DeleteNamespacesBatch(reqDto *dto.DeleteNamespacesBatchDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-namespaces-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 授权资源
 * @description 将一个/多个资源授权给用户、角色、分组、组织机构等主体，且可以分别指定不同的操作权限。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *Client) AuthorizeResources(reqDto *dto.AuthorizeResourcesDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/authorize-resources", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取某个主体被授权的资源列表
 * @description 根据筛选条件，获取某个主体被授权的资源列表。
 * @param targetType 目标对象类型
 * @param targetIdentifier 目标对象唯一标志符
 * @param namespace 所属权限分组的 code
 * @param resourceType 限定资源类型，如数据、API、按钮、菜单
 * @param resourceList 限定查询的资源列表，如果指定，只会返回所指定的资源列表。
 * @param withDenied 是否获取被拒绝的资源
 * @returns AuthorizedResourcePaginatedRespDto
 */
func (c *Client) GetAuthorizedResources(reqDto *dto.GetAuthorizedResourcesDto) *dto.AuthorizedResourcePaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-authorized-resources", fasthttp.MethodGet, reqDto)
	var response dto.AuthorizedResourcePaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 判断用户是否对某个资源的某个操作有权限
 * @description 判断用户是否对某个资源的某个操作有权限。
 * @param requestBody
 * @returns IsActionAllowedRespDtp
 */
func (c *Client) IsActionAllowed(reqDto *dto.IsActionAllowedDto) *dto.IsActionAllowedRespDtp {
	b, err := c.SendHttpRequest("/api/v3/is-action-allowed", fasthttp.MethodPost, reqDto)
	var response dto.IsActionAllowedRespDtp
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
