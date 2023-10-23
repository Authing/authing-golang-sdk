package management

import (
	"encoding/json"
	"fmt"

	"github.com/Authing/authing-golang-sdk/v3/dto"
	"github.com/Authing/authing-golang-sdk/v3/util"
	"github.com/valyala/fasthttp"
)

/*
* @summary 数据对象高级搜索
* @description 数据对象高级搜索
* @param requestBody
* @returns FunctionModelValueListResDto
 */
func (client *ManagementClient) ListRow(reqDto *dto.FilterDto) *dto.FunctionModelValueListResDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/filter", fasthttp.MethodPost, reqDto)
	var response dto.FunctionModelValueListResDto
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
* @summary 获取数据对象行信息
* @description 获取数据对象行信息
* @param modelId 功能 id
* @param rowId 行 id
* @param showFieldId 返回结果中是否使用字段 id 作为 key
* @returns FunctionModelValueResDto
 */
func (client *ManagementClient) GetRow(reqDto *dto.GetRowDto) *dto.FunctionModelValueResDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/get-row", fasthttp.MethodGet, reqDto)
	var response dto.FunctionModelValueResDto
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
* @summary 根据属性值获取数据对象行信息
* @description 根据属性值获取数据对象行信息，只允许通过唯一性字段进行精确查询。
* @param modelId 功能 id
* @param key 字段 key
* @param value 字段值
* @param showFieldId 返回结果中是否使用字段 id 作为 key
* @returns FunctionModelValueResDto
 */
func (client *ManagementClient) GetRowByValue(reqDto *dto.GetRowByValueDto) *dto.FunctionModelValueResDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/get-row-by-value", fasthttp.MethodGet, reqDto)
	var response dto.FunctionModelValueResDto
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
* @summary 批量获取行信息
* @description 批量获取行信息
* @param requestBody
* @returns MetadataListResDto
 */
func (client *ManagementClient) GetRowBatch(reqDto *dto.GetRowBatchDto) *dto.MetadataListResDto {
	b, err := client.SendHttpRequest("/get-row-batch", fasthttp.MethodPost, reqDto)
	var response dto.MetadataListResDto
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
* @summary 添加行
* @description 添加行
* @param requestBody
* @returns FunctionModelValueResDto
 */
func (client *ManagementClient) CreateRow(reqDto *dto.CreateRowDto) *dto.FunctionModelValueResDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/create-row", fasthttp.MethodPost, reqDto)
	var response dto.FunctionModelValueResDto
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
* @summary 更新行
* @description 更新行
* @param requestBody
* @returns FunctionModelValueResDto
 */
func (client *ManagementClient) UpdateRow(reqDto *dto.UpdateRowDto) *dto.FunctionModelValueResDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/update-row", fasthttp.MethodPost, reqDto)
	var response dto.FunctionModelValueResDto
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
* @summary 删除行
* @description 删除行
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) RemoveRow(reqDto *dto.RemoveRowDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/remove-row", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 创建数据对象
* @description 利用此接口可以创建一个自定义的数据对象，定义数据对象的基本信息
* @param requestBody
* @returns FunctionModelResDto
 */
func (client *ManagementClient) CreateModel(reqDto *dto.CreateFunctionModelDto) *dto.FunctionModelResDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/create-model", fasthttp.MethodPost, reqDto)
	var response dto.FunctionModelResDto
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
* @summary 获取数据对象详情
* @description 利用功能 id ，获取数据对象的详细信息
* @param id 功能 id 可以从控制台页面获取
* @returns FunctionModelResDto
 */
func (client *ManagementClient) GetModel(reqDto *dto.GetModelDto) *dto.FunctionModelResDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/get-model", fasthttp.MethodGet, reqDto)
	var response dto.FunctionModelResDto
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
* @summary 获取数据对象列表
* @description 获取数据对象列表
* @returns FunctionModelListDto
 */
func (client *ManagementClient) ListModel() *dto.FunctionModelListDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/list-model", fasthttp.MethodGet, nil)
	var response dto.FunctionModelListDto
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
* @summary 删除数据对象
* @description 根据请求的功能 id ，删除对应的数据对象
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) RemoveModel(reqDto *dto.FunctionModelIdDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/remove-model", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 更新数据对象
* @description 更新对应功能 id 的数据对象信息
* @param requestBody
* @returns FunctionModelResDto
 */
func (client *ManagementClient) UpdateModel(reqDto *dto.UpdateFunctionModelDto) *dto.FunctionModelResDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/update-model", fasthttp.MethodPost, reqDto)
	var response dto.FunctionModelResDto
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
* @summary 创建数据对象的字段
* @description 创建相关数据对象的字段，配置字段信息及基本校验规则
* @param requestBody
* @returns FunctionModelFieldResDto
 */
func (client *ManagementClient) CreateField(reqDto *dto.CreateFunctionModelFieldDto) *dto.FunctionModelFieldResDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/create-field", fasthttp.MethodPost, reqDto)
	var response dto.FunctionModelFieldResDto
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
* @summary 更新数据对象的字段
* @description 更新相关数据对象的字段信息及基本校验规则
* @param requestBody
* @returns FunctionModelFieldResDto
 */
func (client *ManagementClient) UpdateField(reqDto *dto.UpdateFunctionModelFieldDto) *dto.FunctionModelFieldResDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/update-field", fasthttp.MethodPost, reqDto)
	var response dto.FunctionModelFieldResDto
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
* @summary 删除数据对象的字段
* @description 根据功能字段 id 、功能 id 、字段属性名删除对应的字段
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) RemoteField(reqDto *dto.FunctionModelFieldIdDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/remove-field", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 获取数据对象字段列表
* @description 获取数据对象字段列表
* @param modelId 功能 id
* @returns FunctionFieldListResDto
 */
func (client *ManagementClient) ListField(reqDto *dto.ListFieldDto) *dto.FunctionFieldListResDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/list-field", fasthttp.MethodGet, reqDto)
	var response dto.FunctionFieldListResDto
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
* @summary 导出全部数据
* @description 导出全部数据
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) Export(reqDto *dto.ExportModelDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/export", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 导入数据
* @description 导入数据
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) Import(reqDto *dto.ImportModelDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/import", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 获取导入模板
* @description 获取导入模板
* @param modelId 功能 id
* @returns GetImportTemplateResDto
 */
func (client *ManagementClient) GetImportTemplate(reqDto *dto.GetImportTemplateDto) *dto.GetImportTemplateResDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/get-import-template", fasthttp.MethodGet, reqDto)
	var response dto.GetImportTemplateResDto
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
* @summary 创建自定义操作
* @description 创建自定义操作
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) CreateOperate(reqDto *dto.CreateOperateModelDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/create-operate", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 移除自定义操作
* @description 移除自定义操作
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) RemoveOperate(reqDto *dto.FunctionModelOperateIdDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/remove-operate", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 执行自定义操作
* @description 执行自定义操作
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) ExecuteOperate(reqDto *dto.FunctionModelOperateIdDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/execute-operate", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 复制自定义操作
* @description 复制自定义操作
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) CopyOperate(reqDto *dto.FunctionModelOperateIdDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/copy-operate", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 操作管理列表(分页)
* @description 操作管理列表(分页)
* @param modelId model Id
* @param keywords 搜索功能名称
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @returns OperateModelDto
 */
func (client *ManagementClient) ListOperate(reqDto *dto.ListOperateDto) *dto.OperateModelDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/list-operate", fasthttp.MethodGet, reqDto)
	var response dto.OperateModelDto
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
* @summary 全部操作管理列表
* @description 全部操作管理列表
* @param modelId model Id
* @returns OperateModelDto
 */
func (client *ManagementClient) ListOperateAll(reqDto *dto.AllOperateDto) *dto.OperateModelDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/all-operate", fasthttp.MethodGet, reqDto)
	var response dto.OperateModelDto
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
* @summary 更新操作管理
* @description 更新操作管理
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) UpdateOperate(reqDto *dto.UpdateOperateModelDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/update-operate", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 获取关联数据详情
* @description 获取关联数据详情
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) GetRelationInfo(reqDto *dto.GetRelationInfoDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/get-relation-info", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 创建行关联数据
* @description 创建行关联数据
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) CreateRowRelation(reqDto *dto.CreateRelationValueDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/create-row-relation", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 获取行关联数据
* @description 获取行关联数据
* @param modelId 功能 id
* @param fieldId 字段 id
* @param rowId 行 id
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @returns RelationValueListResDto
 */
func (client *ManagementClient) GetRelationValue(reqDto *dto.GetRowRelationDto) *dto.RelationValueListResDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/get-row-relation", fasthttp.MethodGet, reqDto)
	var response dto.RelationValueListResDto
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
* @summary 删除行关联数据
* @description 删除行关联数据
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) RemoveRelationValue(reqDto *dto.RemoveRelationValueDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/remove-row-relation", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 导出数据对象
* @description 导出数据对象
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) ExportModel(reqDto *dto.ExportMetadataDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/export/model", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 导入数据对象
* @description 导入数据对象
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) ImportModel(reqDto *dto.ImportMetadataDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/import/model", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary UEBA 上传
* @description UEBA 上传
* @param requestBody
* @returns CreateUEBARespDto
 */
func (client *ManagementClient) Capture(reqDto *dto.CreateUEBADto) *dto.CreateUEBARespDto {
	b, err := client.SendHttpRequest("/api/v3/metadata/ueba/capture", fasthttp.MethodPost, reqDto)
	var response dto.CreateUEBARespDto
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
* @summary 岗位列表
* @description 岗位列表
* @param keywords 搜索岗位 code 或名称
* @param skipCount 是否统计岗位关联的部门数和用户数
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param withMetadata 是否展示元数据内容
* @param withCustomData 是否获取自定义数据
* @param flatCustomData 是否拍平扩展字段
* @returns PostPaginatedRespDto
 */
func (client *ManagementClient) PostList(reqDto *dto.ListPostDto) *dto.PostPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-post", fasthttp.MethodGet, reqDto)
	var response dto.PostPaginatedRespDto
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
* @summary 获取岗位
* @description 获取岗位
* @param code 岗位 code
* @param withCustomData 是否获取自定义数据
* @returns CreatePostDto
 */
func (client *ManagementClient) GetPost(reqDto *dto.GetPostDto) *dto.CreatePostDto {
	b, err := client.SendHttpRequest("/api/v3/get-post", fasthttp.MethodGet, reqDto)
	var response dto.CreatePostDto
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
* @deprecated
* @summary 获取用户关联岗位
* @description 获取用户关联的所有岗位
* @param userId 用户 id
* @param withCustomData 是否获取自定义数据
* @returns PostListRespDto
 */
func (client *ManagementClient) GetUserPosts(reqDto *dto.GetUserPostsDto) *dto.PostListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-posts", fasthttp.MethodGet, reqDto)
	var response dto.PostListRespDto
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
* @deprecated
* @summary 获取用户关联岗位
* @description 此接口只会返回一个岗位，已废弃，请使用 /api/v3/get-user-posts 接口
* @param userId 用户 id
* @param withCustomData 是否获取自定义数据
* @returns CreatePostDto
 */
func (client *ManagementClient) GetUserPost(reqDto *dto.GetUserPostDto) *dto.CreatePostDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-post", fasthttp.MethodGet, reqDto)
	var response dto.CreatePostDto
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
* @summary 获取岗位信息
* @description 根据岗位 id 获取岗位详情
* @param requestBody
* @returns PostRespDto
 */
func (client *ManagementClient) GetPostById(reqDto *dto.GetPostByIdListDto) *dto.PostRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-post-by-id", fasthttp.MethodPost, reqDto)
	var response dto.PostRespDto
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
* @summary 创建岗位
* @description 创建岗位
* @param requestBody
* @returns CreatePostRespDto
 */
func (client *ManagementClient) CreatePost(reqDto *dto.CreatePostDto) *dto.CreatePostRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-post", fasthttp.MethodPost, reqDto)
	var response dto.CreatePostRespDto
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
* @summary 更新岗位信息
* @description 更新岗位信息
* @param requestBody
* @returns CreatePostRespDto
 */
func (client *ManagementClient) UpdatePost(reqDto *dto.CreatePostDto) *dto.CreatePostRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-post", fasthttp.MethodPost, reqDto)
	var response dto.CreatePostRespDto
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
* @summary 删除岗位
* @description 删除岗位
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) RemovePost(reqDto *dto.RemovePostDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/remove-post", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 用户设置岗位
* @description 一次性给用户设置岗位：如果之前的岗位不在传入的列表中，会进行移除；如果有新增的岗位，会加入到新的岗位；如果不变，则不进行任何操作。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) SetUserPosts(reqDto *dto.SetUserPostsDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/set-user-posts", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 用户关联岗位
* @description 用户关联岗位
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) UserConnectionPost(reqDto *dto.UserConnectionPostDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/user-connection-post", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 移除绑定(用户详情页)
* @description 移除绑定(用户详情页)。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) DeleteDevice(reqDto *dto.DeleteTerminalUserDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/delete-device-by-user", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 挂起设备(用户详情页)
* @description 挂起设备(用户详情页)。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) SuspendDevice(reqDto *dto.SuspendTerminalUserDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/suspend-device-by-user", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 停用设备(用户详情页)
* @description 停用设备(用户详情页)。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) DisableDevice(reqDto *dto.UpdateTerminalUserDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/disable-device-by-user", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 启用设备(用户详情页)
* @description 启用设备(用户详情页)。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) EnableDevice(reqDto *dto.UpdateTerminalUserDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/enable-device-by-user", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 获取设备状态
* @description 获取设备状态。
* @param requestBody
* @returns DeviceStatusRespDto
 */
func (client *ManagementClient) GetDeviceStatus(reqDto *dto.TerminalBaseDto) *dto.DeviceStatusRespDto {
	b, err := client.SendHttpRequest("/api/v3/device-status", fasthttp.MethodPost, reqDto)
	var response dto.DeviceStatusRespDto
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
	* @summary 获取/搜索公共账号列表
	* @description
 * 此接口用于获取用户列表，支持模糊搜索，以及通过用户基础字段、用户自定义字段、用户所在部门、用户历史登录应用等维度筛选用户。
 *
 * ### 模糊搜素示例
 *
 * 模糊搜索默认会从 `phone`, `email`, `name`, `username`, `nickname` 五个字段对用户进行模糊搜索，你也可以通过设置 `options.fuzzySearchOn`
 * 决定模糊匹配的字段范围：
 *
 * ```json
 * {
 * "keywords": "北京",
 * "options": {
 * "fuzzySearchOn": [
 * "address"
 * ]
 * }
 * }
 * ```
 *
 * ### 高级搜索示例
 *
 * 你可以通过 `advancedFilter` 进行高级搜索，高级搜索支持通过用户的基础信息、自定义数据、所在部门、用户来源、登录应用、外部身份源信息等维度对用户进行筛选。
 * **且这些筛选条件可以任意组合。**
 *
 * #### 筛选状态为禁用的用户
 *
 * 用户状态（`status`）为字符串类型，可选值为 `Activated` 和 `Suspended`：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "status",
 * "operator": "EQUAL",
 * "value": "Suspended"
 * }
 * ]
 * }
 * ```
 *
 * #### 筛选邮箱中包含 `@example.com` 的用户
 *
 * 用户邮箱（`email`）为字符串类型，可以进行模糊搜索：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "email",
 * "operator": "CONTAINS",
 * "value": "@example.com"
 * }
 * ]
 * }
 * ```
 *
 * #### 根据用户的任意扩展字段进行搜索
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "some-custom-key",
 * "operator": "EQUAL",
 * "value": "some-value"
 * }
 * ]
 * }
 * ```
 *
 * #### 根据用户登录次数筛选
 *
 * 筛选登录次数大于 10 的用户：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "loginsCount",
 * "operator": "GREATER",
 * "value": 10
 * }
 * ]
 * }
 * ```
 *
 * 筛选登录次数在 10 - 100 次的用户：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "loginsCount",
 * "operator": "BETWEEN",
 * "value": [10, 100]
 * }
 * ]
 * }
 * ```
 *
 * #### 根据用户上次登录时间进行筛选
 *
 * 筛选最近 7 天内登录过的用户：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "lastLoginTime",
 * "operator": "GREATER",
 * "value": new Date(Date.now() - 7 * 24 * 60 * 60 * 1000)
 * }
 * ]
 * }
 * ```
 *
 * 筛选在某一段时间内登录过的用户：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "lastLogin",
 * "operator": "BETWEEN",
 * "value": [
 * Date.now() - 14 * 24 * 60 * 60 * 1000,
 * Date.now() - 7 * 24 * 60 * 60 * 1000
 * ]
 * }
 * ]
 * }
 * ```
 *
 * #### 根据用户曾经登录过的应用筛选
 *
 * 筛选出曾经登录过应用 `appId1` 或者 `appId2` 的用户：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "loggedInApps",
 * "operator": "IN",
 * "value": [
 * "appId1",
 * "appId2"
 * ]
 * }
 * ]
 * }
 * ```
 *
 * #### 根据用户所在部门进行筛选
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "department",
 * "operator": "IN",
 * "value": [
 * {
 * "organizationCode": "steamory",
 * "departmentId": "root",
 * "departmentIdType": "department_id",
 * "includeChildrenDepartments": true
 * }
 * ]
 * }
 * ]
 * }
 * ```
 *
 *
	* @param requestBody
	* @returns PublicAccountPaginatedRespDto
*/
func (client *ManagementClient) ListPublicAccounts(reqDto *dto.ListPublicAccountsRequestDto) *dto.PublicAccountPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-public-accounts", fasthttp.MethodPost, reqDto)
	var response dto.PublicAccountPaginatedRespDto
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
	* @summary 获取公共账号信息
	* @description 通过公共账号用户 ID，获取公共账号详情，可以选择获取自定义数据、选择指定用户 ID 类型等。
	* @param userId 公共账号用户 ID
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 *
	* @param withCustomData 是否获取自定义数据
	* @param withDepartmentIds 是否获取部门 ID 列表
	* @returns PublicAccountSingleRespDto
*/
func (client *ManagementClient) GetPublicAccount(reqDto *dto.GetPublicAccountDto) *dto.PublicAccountSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-public-account", fasthttp.MethodGet, reqDto)
	var response dto.PublicAccountSingleRespDto
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
	* @summary 批量获取公共账号信息
	* @description 通过公共账号用户 ID 列表，批量获取公共账号信息，可以选择获取自定义数据、选择指定用户 ID 类型等。
	* @param userIds 公共账号用户 ID 数组
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 *
	* @param withCustomData 是否获取自定义数据
	* @param withDepartmentIds 是否获取部门 ID 列表
	* @returns PublicAccountListRespDto
*/
func (client *ManagementClient) GetPublicAccountBatch(reqDto *dto.GetPublicAccountBatchDto) *dto.PublicAccountListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-public-account-batch", fasthttp.MethodGet, reqDto)
	var response dto.PublicAccountListRespDto
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
* @summary 创建公共账号
* @description 创建公共账号，邮箱、手机号、用户名必须包含其中一个，邮箱、手机号、用户名、externalId 用户池内唯一，此接口将以管理员身份创建公共账号用户因此不需要进行手机号验证码检验等安全检测。
* @param requestBody
* @returns PublicAccountSingleRespDto
 */
func (client *ManagementClient) CreatePublicAccount(reqDto *dto.CreatePublicAccountReqDto) *dto.PublicAccountSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-public-account", fasthttp.MethodPost, reqDto)
	var response dto.PublicAccountSingleRespDto
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
* @summary 批量创建公共账号
* @description 批量创建公共账号，邮箱、手机号、用户名必须包含其中一个，邮箱、手机号、用户名、externalId 用户池内唯一，此接口将以管理员身份创建公共账号用户因此不需要进行手机号验证码检验等安全检测。
* @param requestBody
* @returns PublicAccountListRespDto
 */
func (client *ManagementClient) CreatePublicAccountsBatch(reqDto *dto.CreatePublicAccountBatchReqDto) *dto.PublicAccountListRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-public-accounts-batch", fasthttp.MethodPost, reqDto)
	var response dto.PublicAccountListRespDto
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
* @summary 修改公共账号资料
* @description 通过公共账号用户 ID，修改公共账号资料，邮箱、手机号、用户名、externalId 用户池内唯一，此接口将以管理员身份修改公共账号资料因此不需要进行手机号验证码检验等安全检测。
* @param requestBody
* @returns PublicAccountSingleRespDto
 */
func (client *ManagementClient) UpdatePublicAccount(reqDto *dto.UpdatePublicAccountReqDto) *dto.PublicAccountSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-public-account", fasthttp.MethodPost, reqDto)
	var response dto.PublicAccountSingleRespDto
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
* @summary 批量修改公共账号资料
* @description 批量修改公共账号资料，邮箱、手机号、用户名、externalId 用户池内唯一，此接口将以管理员身份修改公共账号资料因此不需要进行手机号验证码检验等安全检测。
* @param requestBody
* @returns PublicAccountListRespDto
 */
func (client *ManagementClient) UpdatePublicAccountBatch(reqDto *dto.UpdatePublicAccountBatchReqDto) *dto.PublicAccountListRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-public-account-batch", fasthttp.MethodPost, reqDto)
	var response dto.PublicAccountListRespDto
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
* @summary 批量删除公共账号
* @description 通过公共账号 ID 列表，删除公共账号，支持批量删除，可以选择指定用户 ID 类型等。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeletePublicAccountsBatch(reqDto *dto.DeletePublicAccountsBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-public-accounts-batch", fasthttp.MethodPost, reqDto)
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
* @summary 强制下线公共账号
* @description 通过公共账号 ID、App ID 列表，强制让公共账号下线，可以选择指定公共账号 ID 类型等。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) KickPublicAccounts(reqDto *dto.KickPublicAccountsDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/kick-public-accounts", fasthttp.MethodPost, reqDto)
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
* @summary 个人账号转换为公共账号
* @description 通过用户 ID，把个人账号转换为公共账号。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) ChangeIntoPublicAccount(reqDto *dto.CreatePublicAccountFromUserDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/transfer-into-public-account", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 获取用户的公共账号列表
* @description 通过用户 ID，获取用户的公共账号列表。
* @param userId 用户 ID
* @returns PublicAccountPaginatedRespDto
 */
func (client *ManagementClient) GetPublicAccountsOfUser(reqDto *dto.GetPublicAccountsOfUserDto) *dto.PublicAccountPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-public-accounts-of-user", fasthttp.MethodGet, reqDto)
	var response dto.PublicAccountPaginatedRespDto
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
* @summary 公共账号的用户列表
* @description 通过公共账号 ID，获取用户列表。
* @param publicAccountId 公共账号 ID
* @returns PublicAccountPaginatedRespDto
 */
func (client *ManagementClient) GetUsersOfPublicAccount(reqDto *dto.GetUsersOfPublicAccountDto) *dto.PublicAccountPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-users-of-public-account", fasthttp.MethodGet, reqDto)
	var response dto.PublicAccountPaginatedRespDto
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
* @summary 公共账号绑定批量用户
* @description 使用公共账号绑定批量用户
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) BindUsersPublicAccount(reqDto *dto.SetPublicAccountBatchReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/set-public-account-of-users", fasthttp.MethodPost, reqDto)
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
* @summary 用户绑定批量公共账号
* @description 用户绑定批量公共账号
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) SetuserOfPublicAccount(reqDto *dto.SetUserOfPublicAccountBatchReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/set-user-of-public-accounts", fasthttp.MethodPost, reqDto)
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
* @summary 公共账号解绑用户
* @description 公共账号解绑用户
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) UnbindUsersPublicAccount(reqDto *dto.UnbindPublicAccountBatchReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/unbind-public-account-of-user", fasthttp.MethodPost, reqDto)
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
* @summary 获取组织机构详情
* @description 获取组织机构详情
* @param organizationCode 组织 Code（organizationCode）
* @param withCustomData 是否获取自定义数据
* @param withPost 是否获取 部门信息
* @param tenantId 租户 ID
* @returns OrganizationSingleRespDto
 */
func (client *ManagementClient) GetOrganization(reqDto *dto.GetOrganizationDto) *dto.OrganizationSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-organization", fasthttp.MethodGet, reqDto)
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
* @summary 批量获取组织机构详情
* @description 批量获取组织机构详情
* @param organizationCodeList 组织 Code（organizationCode）列表
* @param withCustomData 是否获取自定义数据
* @param withPost 是否获取 部门信息
* @param tenantId 租户 ID
* @returns OrganizationListRespDto
 */
func (client *ManagementClient) GetOrganizationsBatch(reqDto *dto.GetOrganizationBatchDto) *dto.OrganizationListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-organization-batch", fasthttp.MethodGet, reqDto)
	var response dto.OrganizationListRespDto
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
* @summary 获取组织机构列表
* @description 获取组织机构列表，支持分页。
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param fetchAll 拉取所有
* @param withCustomData 是否获取自定义数据
* @param withPost 是否获取 部门信息
* @param tenantId 租户 ID
* @param status 组织的状态
* @returns OrganizationPaginatedRespDto
 */
func (client *ManagementClient) ListOrganizations(reqDto *dto.ListOrganizationsDto) *dto.OrganizationPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-organizations", fasthttp.MethodGet, reqDto)
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
* @summary 创建组织机构
* @description 创建组织机构，会创建一个只有一个节点的组织机构，可以选择组织描述信息、根节点自定义 ID、多语言等。
* @param requestBody
* @returns OrganizationSingleRespDto
 */
func (client *ManagementClient) CreateOrganization(reqDto *dto.CreateOrganizationReqDto) *dto.OrganizationSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-organization", fasthttp.MethodPost, reqDto)
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
* @summary 修改组织机构
* @description 通过组织 code，修改组织机构，可以选择部门描述、新组织 code、组织名称等。
* @param requestBody
* @returns OrganizationSingleRespDto
 */
func (client *ManagementClient) UpdateOrganization(reqDto *dto.UpdateOrganizationReqDto) *dto.OrganizationSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-organization", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) DeleteOrganization(reqDto *dto.DeleteOrganizationReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-organization", fasthttp.MethodPost, reqDto)
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
* @summary 搜索组织机构列表
* @description 通过搜索关键词，搜索组织机构列表，支持分页。
* @param keywords 搜索关键词，如组织机构名称
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param withCustomData 是否获取自定义数据
* @param tenantId 租户 ID
* @returns OrganizationPaginatedRespDto
 */
func (client *ManagementClient) SearchOrganizations(reqDto *dto.SearchOrganizationsDto) *dto.OrganizationPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/search-organizations", fasthttp.MethodGet, reqDto)
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
* @summary 更新组织机构状态
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) UpdateOrganizationStatus(reqDto *dto.UpdateOrganizationStatusReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-organization-status", fasthttp.MethodPost, reqDto)
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
* @summary 获取部门信息
* @description 通过组织 code 以及 部门 ID 或 部门 code，获取部门信息，可以获取自定义数据。
* @param organizationCode 组织 code
* @param departmentId 部门 ID，根部门传 `root`。departmentId 和 departmentCode 必传其一。
* @param departmentCode 部门 code。departmentId 和 departmentCode 必传其一。
* @param departmentIdType 此次调用中使用的部门 ID 的类型
* @param withCustomData 是否获取自定义数据
* @param flatCustomData 是否拍平扩展字段
* @param tenantId 租户 ID
* @returns DepartmentSingleRespDto
 */
func (client *ManagementClient) GetDepartment(reqDto *dto.GetDepartmentDto) *dto.DepartmentSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-department", fasthttp.MethodGet, reqDto)
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
func (client *ManagementClient) CreateDepartment(reqDto *dto.CreateDepartmentReqDto) *dto.DepartmentSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-department", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) UpdateDepartment(reqDto *dto.UpdateDepartmentReqDto) *dto.DepartmentSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-department", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) DeleteDepartment(reqDto *dto.DeleteDepartmentReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-department", fasthttp.MethodPost, reqDto)
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
* @deprecated
* @summary 搜索部门
* @description 通过组织 code、搜索关键词，搜索部门，可以搜索组织名称等。
* @param requestBody
* @returns DepartmentListRespDto
 */
func (client *ManagementClient) SearchDepartments(reqDto *dto.SearchDepartmentsReqDto) *dto.DepartmentListRespDto {
	b, err := client.SendHttpRequest("/api/v3/search-departments", fasthttp.MethodPost, reqDto)
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
* @summary 搜索部门
* @description 通过组织 code、搜索关键词，搜索部门，可以搜索组织名称等。
* @param requestBody
* @returns DepartmentListRespDto
 */
func (client *ManagementClient) SearchDepartmentsList(reqDto *dto.SearchDepartmentsListReqDto) *dto.DepartmentListRespDto {
	b, err := client.SendHttpRequest("/api/v3/search-departments-list", fasthttp.MethodPost, reqDto)
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
* @param status 部门的状态
* @param departmentIdType 此次调用中使用的部门 ID 的类型
* @param excludeVirtualNode 是否要排除虚拟组织
* @param onlyVirtualNode 是否只包含虚拟组织
* @param withCustomData 是否获取自定义数据
* @param tenantId 租户 ID
* @returns DepartmentPaginatedRespDto
 */
func (client *ManagementClient) ListChildrenDepartments(reqDto *dto.ListChildrenDepartmentsDto) *dto.DepartmentPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-children-departments", fasthttp.MethodGet, reqDto)
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
* @summary 获取所有部门列表
* @description 获取所有部门列表，可以用于获取某个组织下的所有部门列表。
* @param organizationCode 组织 code
* @param departmentId 部门 ID，不填写默认为 `root` 根部门 ID
* @param departmentIdType 此次调用中使用的部门 ID 的类型
* @param withCustomData 是否获取自定义数据
* @returns DepartmentPaginatedRespDto
 */
func (client *ManagementClient) GetAllDepartments(reqDto *dto.GetAllDepartmentsDto) *dto.DepartmentPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-all-departments", fasthttp.MethodGet, reqDto)
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
* @param sortBy 排序依据
* @param orderBy 增序还是倒序
* @param departmentIdType 此次调用中使用的部门 ID 的类型
* @param includeChildrenDepartments 是否包含子部门的成员
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param withCustomData 是否获取自定义数据
* @param withIdentities 是否获取 identities
* @param withDepartmentIds 是否获取部门 ID 列表
* @param tenantId 租户 ID
* @returns UserPaginatedRespDto
 */
func (client *ManagementClient) ListDepartmentMembers(reqDto *dto.ListDepartmentMembersDto) *dto.UserPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-department-members", fasthttp.MethodGet, reqDto)
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
* @param tenantId 租户 ID
* @returns UserIdListRespDto
 */
func (client *ManagementClient) ListDepartmentMemberIds(reqDto *dto.ListDepartmentMemberIdsDto) *dto.UserIdListRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-department-member-ids", fasthttp.MethodGet, reqDto)
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
* @param tenantId 租户 ID
* @returns UserPaginatedRespDto
 */
func (client *ManagementClient) SearchDepartmentMembers(reqDto *dto.SearchDepartmentMembersDto) *dto.UserPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/search-department-members", fasthttp.MethodGet, reqDto)
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
func (client *ManagementClient) AddDepartmentMembers(reqDto *dto.AddDepartmentMembersReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/add-department-members", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) RemoveDepartmentMembers(reqDto *dto.RemoveDepartmentMembersReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/remove-department-members", fasthttp.MethodPost, reqDto)
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
* @param tenantId 租户 ID
* @returns DepartmentSingleRespDto
 */
func (client *ManagementClient) GetParentDepartment(reqDto *dto.GetParentDepartmentDto) *dto.DepartmentSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-parent-department", fasthttp.MethodGet, reqDto)
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
* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
* @param organizationCode 组织 code
* @param departmentId 部门 ID，根部门传 `root`。departmentId 和 departmentCode 必传其一。
* @param departmentIdType 此次调用中使用的部门 ID 的类型
* @param includeChildrenDepartments 是否包含子部门
* @param tenantId 租户 ID
* @returns IsUserInDepartmentRespDto
 */
func (client *ManagementClient) IsUserInDepartment(reqDto *dto.IsUserInDepartmentDto) *dto.IsUserInDepartmentRespDto {
	b, err := client.SendHttpRequest("/api/v3/is-user-in-department", fasthttp.MethodGet, reqDto)
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
* @summary 根据部门id查询部门
* @description 根据部门id查询部门
* @param departmentId 部门 ID
* @param tenantId 租户 ID
* @param withCustomData 是否获取自定义数据
* @returns DepartmentSingleRespDto
 */
func (client *ManagementClient) GetDepartmentById(reqDto *dto.GetDepartmentByIdDto) *dto.DepartmentSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-department-by-id", fasthttp.MethodGet, reqDto)
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
* @summary 根据组织树批量创建部门
* @description 根据组织树批量创建部门，部门名称不存在时会自动创建
* @param requestBody
* @returns CreateDepartmentTreeRespDto
 */
func (client *ManagementClient) CreateDepartmentTree(reqDto *dto.CreateDepartmentTreeReqDto) *dto.CreateDepartmentTreeRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-department-tree", fasthttp.MethodPost, reqDto)
	var response dto.CreateDepartmentTreeRespDto
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
* @summary 获取部门绑定的第三方同步关系
* @description 如果在 Authing 中的部门进行了上下游同步，此接口可以用于查询出在第三方的关联用户信息
* @param organizationCode 组织 code
* @param departmentId 部门 ID，根部门传 `root`。departmentId 和 departmentCode 必传其一。
* @param departmentIdType 此次调用中使用的部门 ID 的类型
* @param withCustomData 是否获取自定义数据
* @param tenantId 租户 ID
* @returns SyncRelationListRespDto
 */
func (client *ManagementClient) GetDepartmentSyncRelations(reqDto *dto.GetDepartmentSyncRelationsDto) *dto.SyncRelationListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-department-sync-relations", fasthttp.MethodGet, reqDto)
	var response dto.SyncRelationListRespDto
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
* @summary 删除部门同步关联关系
* @description 如果在 Authing 中的部门进行了上下游同步，此接口可以用于删除某个部门在指定身份源下的关联关系。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteDepartmentSyncRelations(reqDto *dto.DeleteDepartmentSyncRelationReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-department-sync-relations", fasthttp.MethodPost, reqDto)
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
* @summary 更新部门状态
* @description 启用和禁用部门
* @param requestBody
* @returns Node
 */
func (client *ManagementClient) UpdateNodeStatus(reqDto *dto.UpdateDepartmentStatusReqDto) *dto.Node {
	b, err := client.SendHttpRequest("/api/v3/update-department-status", fasthttp.MethodPost, reqDto)
	var response dto.Node
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
	* @summary 获取/搜索用户列表
	* @description
 * 此接口用于获取用户列表，支持模糊搜索，以及通过用户基础字段、用户自定义字段、用户所在部门、用户历史登录应用等维度筛选用户。
 *
 * ### 模糊搜素示例
 *
 * 模糊搜索默认会从 `phone`, `email`, `name`, `username`, `nickname` 五个字段对用户进行模糊搜索，你也可以通过设置 `options.fuzzySearchOn`
 * 决定模糊匹配的字段范围：
 *
 * ```json
 * {
 * "keywords": "北京",
 * "options": {
 * "fuzzySearchOn": [
 * "address"
 * ]
 * }
 * }
 * ```
 *
 * ### 高级搜索示例
 *
 * 你可以通过 `advancedFilter` 进行高级搜索，高级搜索支持通过用户的基础信息、自定义数据、所在部门、用户来源、登录应用、外部身份源信息等维度对用户进行筛选。
 * **且这些筛选条件可以任意组合。**
 *
 * #### 筛选状态为禁用的用户
 *
 * 用户状态（`status`）为字符串类型，可选值为 `Activated` 和 `Suspended`：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "status",
 * "operator": "EQUAL",
 * "value": "Suspended"
 * }
 * ]
 * }
 * ```
 *
 * #### 筛选邮箱中包含 `@example.com` 的用户
 *
 * 用户邮箱（`email`）为字符串类型，可以进行模糊搜索：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "email",
 * "operator": "CONTAINS",
 * "value": "@example.com"
 * }
 * ]
 * }
 * ```
 *
 * #### 根据用户的任意扩展字段进行搜索
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "some-custom-key",
 * "operator": "EQUAL",
 * "value": "some-value"
 * }
 * ]
 * }
 * ```
 *
 * #### 根据用户登录次数筛选
 *
 * 筛选登录次数大于 10 的用户：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "loginsCount",
 * "operator": "GREATER",
 * "value": 10
 * }
 * ]
 * }
 * ```
 *
 * 筛选登录次数在 10 - 100 次的用户：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "loginsCount",
 * "operator": "BETWEEN",
 * "value": [10, 100]
 * }
 * ]
 * }
 * ```
 *
 * #### 根据用户上次登录时间进行筛选
 *
 * 筛选最近 7 天内登录过的用户：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "lastLoginTime",
 * "operator": "GREATER",
 * "value": new Date(Date.now() - 7 * 24 * 60 * 60 * 1000)
 * }
 * ]
 * }
 * ```
 *
 * 筛选在某一段时间内登录过的用户：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "lastLogin",
 * "operator": "BETWEEN",
 * "value": [
 * Date.now() - 14 * 24 * 60 * 60 * 1000,
 * Date.now() - 7 * 24 * 60 * 60 * 1000
 * ]
 * }
 * ]
 * }
 * ```
 *
 * #### 根据用户曾经登录过的应用筛选
 *
 * 筛选出曾经登录过应用 `appId1` 或者 `appId2` 的用户：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "loggedInApps",
 * "operator": "IN",
 * "value": [
 * "appId1",
 * "appId2"
 * ]
 * }
 * ]
 * }
 * ```
 *
 * #### 根据用户所在部门进行筛选
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "department",
 * "operator": "IN",
 * "value": [
 * {
 * "organizationCode": "steamory",
 * "departmentId": "root",
 * "departmentIdType": "department_id",
 * "includeChildrenDepartments": true
 * }
 * ]
 * }
 * ]
 * }
 * ```
 *
 *
	* @param requestBody
	* @returns UserPaginatedRespDto
*/
func (client *ManagementClient) ListUsers(reqDto *dto.ListUsersRequestDto) *dto.UserPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-users", fasthttp.MethodPost, reqDto)
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
* @deprecated
* @summary 获取用户列表
* @description 获取用户列表接口，支持分页，可以选择获取自定义数据、identities 等。
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param status 账户当前状态，如 已停用、已离职、正常状态、已归档
* @param updatedAtStart 用户创建、修改开始时间，为精确到秒的 UNIX 时间戳；支持获取从某一段时间之后的增量数据
* @param updatedAtEnd 用户创建、修改终止时间，为精确到秒的 UNIX 时间戳；支持获取某一段时间内的增量数据。默认为当前时间
* @param withCustomData 是否获取自定义数据
* @param withPost 是否获取 部门信息
* @param withIdentities 是否获取 identities
* @param withDepartmentIds 是否获取部门 ID 列表
* @returns UserPaginatedRespDto
 */
func (client *ManagementClient) ListUsersLegacy(reqDto *dto.ListUsersDto) *dto.UserPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-users", fasthttp.MethodGet, reqDto)
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
	* @summary 获取用户信息
	* @description 通过用户 ID，获取用户详情，可以选择获取自定义数据、identities、选择指定用户 ID 类型等。
	* @param userId 用户 ID
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @param flatCustomData 是否拍平扩展字段
	* @param withCustomData 是否获取自定义数据
	* @param withPost 是否获取 部门信息
	* @param withIdentities 是否获取 identities
	* @param withDepartmentIds 是否获取部门 ID 列表
	* @returns UserSingleRespDto
*/
func (client *ManagementClient) GetUser(reqDto *dto.GetUserDto) *dto.UserSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user", fasthttp.MethodGet, reqDto)
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
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @param withCustomData 是否获取自定义数据
	* @param flatCustomData 是否拍平扩展字段
	* @param withIdentities 是否获取 identities
	* @param withDepartmentIds 是否获取部门 ID 列表
	* @returns UserListRespDto
*/
func (client *ManagementClient) GetUserBatch(reqDto *dto.GetUserBatchDto) *dto.UserListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-batch", fasthttp.MethodGet, reqDto)
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
* @summary 用户属性解密
* @description 接口接收加密信息，返回解密信息
* @param requestBody
* @returns UserFieldDecryptRespDto
 */
func (client *ManagementClient) UserFieldDecrypt(reqDto *dto.UserFieldDecryptReqDto) *dto.UserFieldDecryptRespDto {
	b, err := client.SendHttpRequest("/api/v3/users/field/decrypt", fasthttp.MethodPost, reqDto)
	var response dto.UserFieldDecryptRespDto
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
func (client *ManagementClient) CreateUser(reqDto *dto.CreateUserReqDto) *dto.UserSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-user", fasthttp.MethodPost, reqDto)
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
* @description 批量创建用户，邮箱、手机号、用户名必须包含其中一个，邮箱、手机号、用户名、externalId 用户池内唯一，此接口将以管理员身份创建用户因此不需要进行手机号验证码检验等安全检测。
* @param requestBody
* @returns UserListRespDto
 */
func (client *ManagementClient) CreateUsersBatch(reqDto *dto.CreateUserBatchReqDto) *dto.UserListRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-users-batch", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) UpdateUser(reqDto *dto.UpdateUserReqDto) *dto.UserSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-user", fasthttp.MethodPost, reqDto)
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
* @summary 批量修改用户资料
* @description 批量修改用户资料，邮箱、手机号、用户名、externalId 用户池内唯一，此接口将以管理员身份修改用户资料因此不需要进行手机号验证码检验等安全检测。
* @param requestBody
* @returns UserListRespDto
 */
func (client *ManagementClient) UpdateUserBatch(reqDto *dto.UpdateUserBatchReqDto) *dto.UserListRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-user-batch", fasthttp.MethodPost, reqDto)
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
* @summary 批量删除用户
* @description 通过用户 ID 列表，删除用户，支持批量删除，可以选择指定用户 ID 类型等。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteUsersBatch(reqDto *dto.DeleteUsersBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-users-batch", fasthttp.MethodPost, reqDto)
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
	* @summary 获取用户的外部身份源
	* @description 通过用户 ID，获取用户的外部身份源、选择指定用户 ID 类型。
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @returns IdentityListRespDto
*/
func (client *ManagementClient) GetUserIdentities(reqDto *dto.GetUserIdentitiesDto) *dto.IdentityListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-identities", fasthttp.MethodGet, reqDto)
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
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @param namespace 所属权限分组(权限空间)的 Code
	* @returns RolePaginatedRespDto
*/
func (client *ManagementClient) GetUserRoles(reqDto *dto.GetUserRolesDto) *dto.RolePaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-roles", fasthttp.MethodGet, reqDto)
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
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @returns PrincipalAuthenticationInfoPaginatedRespDto
*/
func (client *ManagementClient) GetUserPrincipalAuthenticationInfo(reqDto *dto.GetUserPrincipalAuthenticationInfoDto) *dto.PrincipalAuthenticationInfoPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-principal-authentication-info", fasthttp.MethodGet, reqDto)
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
func (client *ManagementClient) ResetUserPrincipalAuthenticationInfo(reqDto *dto.ResetUserPrincipalAuthenticationInfoDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/reset-user-principal-authentication-info", fasthttp.MethodPost, reqDto)
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
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @param page 当前页数，从 1 开始
	* @param limit 每页数目，最大不能超过 50，默认为 10
	* @param withCustomData 是否获取自定义数据
	* @param withDepartmentPaths 是否获取部门路径
	* @param sortBy 排序依据，如 部门创建时间、加入部门时间、部门名称、部门标志符
	* @param orderBy 增序或降序
	* @returns UserDepartmentPaginatedRespDto
*/
func (client *ManagementClient) GetUserDepartments(reqDto *dto.GetUserDepartmentsDto) *dto.UserDepartmentPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-departments", fasthttp.MethodGet, reqDto)
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
func (client *ManagementClient) SetUserDepartments(reqDto *dto.SetUserDepartmentsDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/set-user-departments", fasthttp.MethodPost, reqDto)
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
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @returns GroupPaginatedRespDto
*/
func (client *ManagementClient) GetUserGroups(reqDto *dto.GetUserGroupsDto) *dto.GroupPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-groups", fasthttp.MethodGet, reqDto)
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
	* @summary 获取用户 MFA 绑定信息
	* @description 通过用户 ID，获取用户 MFA 绑定信息，可以选择指定用户 ID 类型等。
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @returns UserMfaSingleRespDto
*/
func (client *ManagementClient) GetUserMfaInfo(reqDto *dto.GetUserMfaInfoDto) *dto.UserMfaSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-mfa-info", fasthttp.MethodGet, reqDto)
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
func (client *ManagementClient) ListArchivedUsers(reqDto *dto.ListArchivedUsersDto) *dto.ListArchivedUsersSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-archived-users", fasthttp.MethodGet, reqDto)
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
func (client *ManagementClient) KickUsers(reqDto *dto.KickUsersDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/kick-users", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) IsUserExists(reqDto *dto.IsUserExistsReqDto) *dto.IsUserExistsRespDto {
	b, err := client.SendHttpRequest("/api/v3/is-user-exists", fasthttp.MethodPost, reqDto)
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
	* @summary 获取用户可访问的应用
	* @description 通过用户 ID，获取用户可访问的应用，可以选择指定用户 ID 类型等。
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @returns AppListRespDto
*/
func (client *ManagementClient) GetUserAccessibleApps(reqDto *dto.GetUserAccessibleAppsDto) *dto.AppListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-accessible-apps", fasthttp.MethodGet, reqDto)
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
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @returns AppListRespDto
*/
func (client *ManagementClient) GetUserAuthorizedApps(reqDto *dto.GetUserAuthorizedAppsDto) *dto.AppListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-authorized-apps", fasthttp.MethodGet, reqDto)
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
func (client *ManagementClient) HasAnyRole(reqDto *dto.HasAnyRoleReqDto) *dto.HasAnyRoleRespDto {
	b, err := client.SendHttpRequest("/api/v3/has-any-role", fasthttp.MethodPost, reqDto)
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
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @param appId 应用 ID
	* @param clientIp 客户端 IP
	* @param start 开始时间戳（毫秒）
	* @param end 结束时间戳（毫秒）
	* @param page 当前页数，从 1 开始
	* @param limit 每页数目，最大不能超过 50，默认为 10
	* @returns UserLoginHistoryPaginatedRespDto
*/
func (client *ManagementClient) GetUserLoginHistory(reqDto *dto.GetUserLoginHistoryDto) *dto.UserLoginHistoryPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-login-history", fasthttp.MethodGet, reqDto)
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
	* @summary 获取用户曾经登录过的应用
	* @description 通过用户 ID，获取用户曾经登录过的应用，可以选择指定用户 ID 类型等。
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @returns UserLoggedInAppsListRespDto
*/
func (client *ManagementClient) GetUserLoggedinApps(reqDto *dto.GetUserLoggedinAppsDto) *dto.UserLoggedInAppsListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-loggedin-apps", fasthttp.MethodGet, reqDto)
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
	* @summary 获取用户曾经登录过的身份源
	* @description 通过用户 ID，获取用户曾经登录过的身份源，可以选择指定用户 ID 类型等。
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @returns UserLoggedInIdentitiesRespDto
*/
func (client *ManagementClient) GetUserLoggedinIdentities(reqDto *dto.GetUserLoggedInIdentitiesDto) *dto.UserLoggedInIdentitiesRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-logged-in-identities", fasthttp.MethodGet, reqDto)
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
	* @summary 离职用户
	* @description 离职用户。离职操作会进行以下操作：
 *
 * - 离职后该成员授权、部门、角色、分组、岗位关系将被删除；
 * - 离职后将保留用户基本信息，同时账号将被禁用，无法登录应用；如果需要彻底删除账号，请调用删除接口。
 *
 * 该操作不可恢复，请谨慎操作！
 *
	* @param requestBody
	* @returns ResignUserRespDto
*/
func (client *ManagementClient) ResignUser(reqDto *dto.ResignUserReqDto) *dto.ResignUserRespDto {
	b, err := client.SendHttpRequest("/api/v3/resign-user", fasthttp.MethodPost, reqDto)
	var response dto.ResignUserRespDto
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
	* @summary 批量离职用户
	* @description 批量离职用户。离职操作会进行以下操作：
 *
 * - 离职后该成员授权、部门、角色、分组、岗位关系将被删除；
 * - 离职后将保留用户基本信息，同时账号将被禁用，无法登录应用；如果需要彻底删除账号，请调用删除接口。
 *
 * 该操作不可恢复，请谨慎操作！
 *
	* @param requestBody
	* @returns ResignUserRespDto
*/
func (client *ManagementClient) ResignUserBatch(reqDto *dto.ResignUserBatchReqDto) *dto.ResignUserRespDto {
	b, err := client.SendHttpRequest("/api/v3/resign-user-batch", fasthttp.MethodPost, reqDto)
	var response dto.ResignUserRespDto
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
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @param namespace 所属权限分组(权限空间)的 Code
	* @param resourceType 资源类型，如 数据、API、菜单、按钮
	* @returns AuthorizedResourcePaginatedRespDto
*/
func (client *ManagementClient) GetUserAuthorizedResources(reqDto *dto.GetUserAuthorizedResourcesDto) *dto.AuthorizedResourcePaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-authorized-resources", fasthttp.MethodGet, reqDto)
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
* @summary 检查某个用户在应用下是否具备 Session 登录态
* @description 检查某个用户在应用下是否具备 Session 登录态
* @param requestBody
* @returns CheckSessionStatusRespDto
 */
func (client *ManagementClient) CheckSessionStatus(reqDto *dto.CheckSessionStatusDto) *dto.CheckSessionStatusRespDto {
	b, err := client.SendHttpRequest("/api/v3/check-session-status", fasthttp.MethodPost, reqDto)
	var response dto.CheckSessionStatusRespDto
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
* @summary 导入用户的 OTP
* @description 导入用户的 OTP
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) ImportOtp(reqDto *dto.ImportOtpReqDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/import-otp", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
	* @summary 获取用户绑定 OTP 的秘钥
	* @description 通过用户 ID，获取用户绑定 OTP 的秘钥。可以选择指定用户 ID 类型等。
	* @param userId 用户 ID
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @returns GetOtpSecretRespDto
*/
func (client *ManagementClient) GetOtpSecretByUser(reqDto *dto.GetOtpSecretByUserDto) *dto.GetOtpSecretRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-otp-secret-by-user", fasthttp.MethodGet, reqDto)
	var response dto.GetOtpSecretRespDto
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
* @summary 获取用户自定义加密的密码
* @description 此功能主要是用户在控制台配置加基于 RSA、SM2 等加密的密钥后，加密用户的密码。
* @param requestBody
* @returns GetUserPasswordCiphertextRespDto
 */
func (client *ManagementClient) GetUserPasswordCiphertext(reqDto *dto.GetUserPasswordCiphertextDto) *dto.GetUserPasswordCiphertextRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-password-ciphertext", fasthttp.MethodPost, reqDto)
	var response dto.GetUserPasswordCiphertextRespDto
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
* @summary 给用户绑定一个身份信息
* @description 用户池管理员手动将来自外部身份源的身份信息绑定到用户上。绑定完成后，可以用执行过绑定操作的身份源登录到对应的 Authing 用户。
* @param requestBody
* @returns LinkIdentityResDto
 */
func (client *ManagementClient) LinkIdentity(reqDto *dto.LinkIdentityDto) *dto.LinkIdentityResDto {
	b, err := client.SendHttpRequest("/api/v3/link-identity", fasthttp.MethodPost, reqDto)
	var response dto.LinkIdentityResDto
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
* @summary 解除绑定用户在身份源下的所有身份信息
* @description 解除绑定用户在某个身份源下的所有身份信息。解绑后，将无法使用执行过解绑操作的身份源登录到对应的 Authing 用户，除非重新绑定身份信息。
* @param requestBody
* @returns UnlinkIdentityResDto
 */
func (client *ManagementClient) UnlinkIdentity(reqDto *dto.UnlinkIdentity) *dto.UnlinkIdentityResDto {
	b, err := client.SendHttpRequest("/api/v3/unlink-identity", fasthttp.MethodPost, reqDto)
	var response dto.UnlinkIdentityResDto
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
* @summary 设置用户 MFA 状态
* @description 设置用户 MFA 状态，即 MFA 触发数据。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) SetUsersMfaStatus(reqDto *dto.SetMfaStatusDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/set-mfa-status", fasthttp.MethodPost, reqDto)
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
	* @summary 获取用户 MFA 状态
	* @description 获取用户 MFA 状态，即 MFA 触发数据。
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @returns GetMapInfoRespDto
*/
func (client *ManagementClient) GetUserMfaStatus(reqDto *dto.GetMfaStatusDto) *dto.GetMapInfoRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-mfa-status", fasthttp.MethodGet, reqDto)
	var response dto.GetMapInfoRespDto
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
	* @summary 获取用户绑定的第三方同步关系
	* @description 如果在 Authing 中的用户进行了上下游同步，此接口可以用于查询出在第三方的关联用户信息
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @returns SyncRelationListRespDto
*/
func (client *ManagementClient) GetUserSyncRelations(reqDto *dto.GetUserSyncRelationsDto) *dto.SyncRelationListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-sync-relations", fasthttp.MethodGet, reqDto)
	var response dto.SyncRelationListRespDto
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
* @summary 删除用户同步关联关系
* @description 如果在 Authing 中的用户进行了上下游同步，此接口可以用于删除某个用户在指定身份源下的关联关系。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteUserSyncRelations(reqDto *dto.DeleteUserSyncRelationReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-user-sync-relations", fasthttp.MethodPost, reqDto)
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
	* @summary 获取公共账号的角色列表
	* @description 通过用户 ID，获取用户角色列表，可以选择所属权限分组 code、选择指定用户 ID 类型等。
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @param namespace 所属权限分组(权限空间)的 code
	* @returns PublicAccountPaginatedRespDto
*/
func (client *ManagementClient) GetPublicAccountRoles(reqDto *dto.GetRolesOfPublicAccountDto) *dto.PublicAccountPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-roles-of-public-account", fasthttp.MethodGet, reqDto)
	var response dto.PublicAccountPaginatedRespDto
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
* @summary 获取角色的公共账号列表
* @description 通过角色 ID，获取用户的公共账号列表。
* @param roleId 角色 ID
* @returns PublicAccountPaginatedRespDto
 */
func (client *ManagementClient) GetPublicAccountsOfRole(reqDto *dto.GetPublicAccountsOfRoleDto) *dto.PublicAccountPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-public-accounts-of-role", fasthttp.MethodGet, reqDto)
	var response dto.PublicAccountPaginatedRespDto
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
* @summary 公共账号绑定批量角色
* @description 公共账号绑定批量角色
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) BindPublicAccountOfRoles(reqDto *dto.SetUserRolesDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/set-public-account-of-roles", fasthttp.MethodPost, reqDto)
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
* @summary 获取分组的公共账号列表
* @description 通过分组 ID，获取用户的公共账号列表。
* @param groupId 分组 ID
* @returns PublicAccountPaginatedRespDto
 */
func (client *ManagementClient) GetPublicAccountsOfGroup(reqDto *dto.GetPublicAccountsOfGroupDto) *dto.PublicAccountPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-public-accounts-of-group", fasthttp.MethodGet, reqDto)
	var response dto.PublicAccountPaginatedRespDto
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
	* @summary 获取公共账号分组列表
	* @description 通过公共账号 ID，获取公共账号分组列表，可以选择指定用户 ID 类型等。
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @returns GroupPaginatedRespDto
*/
func (client *ManagementClient) GetGroupsOfPublicAccount(reqDto *dto.GetGroupsOfPublicAccountDto) *dto.GroupPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-groups-of-public-account", fasthttp.MethodGet, reqDto)
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
* @summary 公共账号添加批量分组
* @description 公共账号通过分组 ID 添加批量分组
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) GetPublicAccountOfGroups(reqDto *dto.SetUserGroupsDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/set-public-account-of-groups", fasthttp.MethodPost, reqDto)
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
* @summary 获取部门的公共账号列表
* @description 通过部门 ID，获取用户的公共账号列表。
* @param departmentId 部门 ID
* @returns PublicAccountPaginatedRespDto
 */
func (client *ManagementClient) GetPublicAccountsOfDepartment(reqDto *dto.GetPublicAccountsOfDepartmentDto) *dto.PublicAccountPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-public-accounts-of-department", fasthttp.MethodGet, reqDto)
	var response dto.PublicAccountPaginatedRespDto
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
	* @summary 获取公共账号的部门列表
	* @description 通过用户 ID，获取用户部门列表，支持分页，可以选择获取自定义数据、选择指定用户 ID 类型、增序或降序等。
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @param page 当前页数，从 1 开始
	* @param limit 每页数目，最大不能超过 50，默认为 10
	* @param withCustomData 是否获取自定义数据
	* @param withDepartmentPaths 是否获取部门路径
	* @param sortBy 排序依据，如 部门创建时间、加入部门时间、部门名称、部门标志符
	* @param orderBy 增序或降序
	* @returns UserDepartmentPaginatedRespDto
*/
func (client *ManagementClient) GetPublicAccountDepartments(reqDto *dto.GetDepartmentsOfPublicAccountDto) *dto.UserDepartmentPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-departments-of-public-account", fasthttp.MethodGet, reqDto)
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
* @summary 设置公共账号所在部门
* @description 设置公共账号所在部门。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) SetPublicAccountOfDepartments(reqDto *dto.SetUserDepartmentsDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/set-public-account-of-departments", fasthttp.MethodPost, reqDto)
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
	* @summary 批量离职用户
	* @description 批量离职用户。离职操作会进行以下操作：
 *
 * - 离职后该成员授权、部门、角色、分组、岗位关系将被删除；
 * - 离职后将保留用户基本信息，同时账号将被禁用，无法登录应用；如果需要彻底删除账号，请调用删除接口。
 *
 * 该操作不可恢复，请谨慎操作！
 *
	* @param requestBody
	* @returns ResignUserRespDto
*/
func (client *ManagementClient) ResignPublicAccountBatch(reqDto *dto.ResignUserBatchReqDto) *dto.ResignUserRespDto {
	b, err := client.SendHttpRequest("/api/v3/resign-public-account-batch", fasthttp.MethodPost, reqDto)
	var response dto.ResignUserRespDto
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
* @summary 获取公共账号的岗位
* @description 获取公共账号的岗位
* @param userId 用户 id
* @returns CreatePostDto
 */
func (client *ManagementClient) GetPostOfPublicUser(reqDto *dto.GetPostOfPublicAccountDto) *dto.CreatePostDto {
	b, err := client.SendHttpRequest("/api/v3/get-post-of-public-account", fasthttp.MethodGet, reqDto)
	var response dto.CreatePostDto
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
* @summary 获取岗位的公共账号列表
* @description 通过岗位 ID，获取用户的公共账号列表。
* @param postId 岗位 ID
* @returns PublicAccountPaginatedRespDto
 */
func (client *ManagementClient) GetPublicAccountsOfPost(reqDto *dto.GetPublicAccountsOfPostDto) *dto.PublicAccountPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-public-accounts-of-post", fasthttp.MethodGet, reqDto)
	var response dto.PublicAccountPaginatedRespDto
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
* @summary 设置公共账号的岗位
* @description 设置公共账号关联的岗位
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) SetPublicAccountOfnPost(reqDto *dto.UserConnectionPostDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/set-public-account-of-post", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 解绑公共账号关联岗位
* @description 解绑公共账号关联岗位
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) UnbindPublicAccountOfPost(reqDto *dto.UserConnectionPostDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/unbind-public-account-of-post", fasthttp.MethodPost, reqDto)
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
* @summary 获取同步任务详情
* @description 获取同步任务详情
* @param syncTaskId 同步任务 ID
* @returns SyncTaskSingleRespDto
 */
func (client *ManagementClient) GetSyncTask(reqDto *dto.GetSyncTaskDto) *dto.SyncTaskSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-sync-task", fasthttp.MethodGet, reqDto)
	var response dto.SyncTaskSingleRespDto
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
* @summary 获取同步任务列表
* @description 获取同步任务列表
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @returns SyncTaskPaginatedRespDto
 */
func (client *ManagementClient) ListSyncTasks(reqDto *dto.ListSyncTasksDto) *dto.SyncTaskPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-sync-tasks", fasthttp.MethodGet, reqDto)
	var response dto.SyncTaskPaginatedRespDto
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
* @summary 创建同步任务
* @description 创建同步任务
* @param requestBody
* @returns SyncTaskPaginatedRespDto
 */
func (client *ManagementClient) CreateSyncTask(reqDto *dto.CreateSyncTaskDto) *dto.SyncTaskPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-sync-task", fasthttp.MethodPost, reqDto)
	var response dto.SyncTaskPaginatedRespDto
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
* @summary 修改同步任务
* @description 修改同步任务
* @param requestBody
* @returns SyncTaskPaginatedRespDto
 */
func (client *ManagementClient) UpdateSyncTask(reqDto *dto.UpdateSyncTaskDto) *dto.SyncTaskPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-sync-task", fasthttp.MethodPost, reqDto)
	var response dto.SyncTaskPaginatedRespDto
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
* @summary 执行同步任务
* @description 执行同步任务
* @param requestBody
* @returns TriggerSyncTaskRespDto
 */
func (client *ManagementClient) TriggerSyncTask(reqDto *dto.TriggerSyncTaskDto) *dto.TriggerSyncTaskRespDto {
	b, err := client.SendHttpRequest("/api/v3/trigger-sync-task", fasthttp.MethodPost, reqDto)
	var response dto.TriggerSyncTaskRespDto
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
* @summary 获取同步作业详情
* @description 获取同步作业详情
* @param syncJobId 同步作业 ID
* @returns SyncJobSingleRespDto
 */
func (client *ManagementClient) GetSyncJob(reqDto *dto.GetSyncJobDto) *dto.SyncJobSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-sync-job", fasthttp.MethodGet, reqDto)
	var response dto.SyncJobSingleRespDto
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
	* @summary 获取同步作业详情
	* @description 获取同步作业详情
	* @param syncTaskId 同步任务 ID
	* @param page 当前页数，从 1 开始
	* @param limit 每页数目，最大不能超过 50，默认为 10
	* @param syncTrigger 同步任务触发类型：
 * - `manually`: 手动触发执行
 * - `timed`: 定时触发
 * - `automatic`: 根据事件自动触发
 *
	* @returns SyncJobPaginatedRespDto
*/
func (client *ManagementClient) ListSyncJobs(reqDto *dto.ListSyncJobsDto) *dto.SyncJobPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-sync-jobs", fasthttp.MethodGet, reqDto)
	var response dto.SyncJobPaginatedRespDto
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
	* @summary 获取同步作业详情
	* @description 获取同步作业详情
	* @param syncJobId 同步作业 ID
	* @param page 当前页数，从 1 开始
	* @param limit 每页数目，最大不能超过 50，默认为 10
	* @param success 根据是否操作成功进行筛选
	* @param action 根据操作类型进行筛选：
 * - `CreateUser`: 创建用户
 * - `UpdateUser`: 修改用户信息
 * - `DeleteUser`: 删除用户
 * - `UpdateUserIdentifier`: 修改用户唯一标志符
 * - `ChangeUserDepartment`: 修改用户部门
 * - `CreateDepartment`: 创建部门
 * - `UpdateDepartment`: 修改部门信息
 * - `DeleteDepartment`: 删除部门
 * - `MoveDepartment`: 移动部门
 * - `UpdateDepartmentLeader`: 同步部门负责人
 * - `CreateGroup`: 创建分组
 * - `UpdateGroup`: 修改分组
 * - `DeleteGroup`: 删除分组
 * - `Updateless`: 无更新
 *
	* @param objectType 操作对象类型:
 * - `department`: 部门
 * - `user`: 用户
 *
	* @returns TriggerSyncTaskRespDto
*/
func (client *ManagementClient) ListSyncJobLogs(reqDto *dto.ListSyncJobLogsDto) *dto.TriggerSyncTaskRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-sync-job-logs", fasthttp.MethodGet, reqDto)
	var response dto.TriggerSyncTaskRespDto
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
	* @summary 获取同步风险操作列表
	* @description 获取同步风险操作列表
	* @param syncTaskId 同步任务 ID
	* @param page 当前页数，从 1 开始
	* @param limit 每页数目，最大不能超过 50，默认为 10
	* @param status 根据执行状态筛选
	* @param objectType 根据操作对象类型，默认获取所有类型的记录：
 * - `department`: 部门
 * - `user`: 用户
 *
	* @returns SyncRiskOperationPaginatedRespDto
*/
func (client *ManagementClient) ListSyncRiskOperations(reqDto *dto.ListSyncRiskOperationsDto) *dto.SyncRiskOperationPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-sync-risk-operations", fasthttp.MethodGet, reqDto)
	var response dto.SyncRiskOperationPaginatedRespDto
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
* @summary 执行同步风险操作
* @description 执行同步风险操作
* @param requestBody
* @returns TriggerSyncRiskOperationsRespDto
 */
func (client *ManagementClient) TriggerSyncRiskOperations(reqDto *dto.TriggerSyncRiskOperationDto) *dto.TriggerSyncRiskOperationsRespDto {
	b, err := client.SendHttpRequest("/api/v3/trigger-sync-risk-operations", fasthttp.MethodPost, reqDto)
	var response dto.TriggerSyncRiskOperationsRespDto
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
* @summary 取消同步风险操作
* @description 取消同步风险操作
* @param requestBody
* @returns CancelSyncRiskOperationsRespDto
 */
func (client *ManagementClient) CancelSyncRiskOperation(reqDto *dto.CancelSyncRiskOperationDto) *dto.CancelSyncRiskOperationsRespDto {
	b, err := client.SendHttpRequest("/api/v3/cancel-sync-risk-operation", fasthttp.MethodPost, reqDto)
	var response dto.CancelSyncRiskOperationsRespDto
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
* @param withCustomData 是否获取自定义数据
* @returns GroupSingleRespDto
 */
func (client *ManagementClient) GetGroup(reqDto *dto.GetGroupDto) *dto.GroupSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-group", fasthttp.MethodGet, reqDto)
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
* @param withMetadata 是否展示元数据内容
* @param withCustomData 是否获取自定义数据
* @param flatCustomData 是否拍平扩展字段
* @returns GroupPaginatedRespDto
 */
func (client *ManagementClient) ListGroups(reqDto *dto.ListGroupsDto) *dto.GroupPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-groups", fasthttp.MethodGet, reqDto)
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
* @summary 获取所有分组
* @description 获取所有分组
* @param fetchMembers 是否获取成员列表
* @param withCustomData 是否获取自定义数据
* @returns GroupListRespDto
 */
func (client *ManagementClient) GetAllGroups(reqDto *dto.GetAllGroupsDto) *dto.GroupListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-all-groups", fasthttp.MethodGet, reqDto)
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
* @summary 创建分组
* @description 创建分组，一个分组必须包含分组名称与唯一标志符 code，且必须为一个合法的英文标志符，如 developers。
* @param requestBody
* @returns GroupSingleRespDto
 */
func (client *ManagementClient) CreateGroup(reqDto *dto.CreateGroupReqDto) *dto.GroupSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-group", fasthttp.MethodPost, reqDto)
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
* @summary 创建或修改分组
* @description 不存在时则创建，存在时则进行更新。
* @param requestBody
* @returns CreateOrUpdateGroupRespDto
 */
func (client *ManagementClient) CreateOrUpdateGroup(reqDto *dto.CreateOrUpdateGroupReqDto) *dto.CreateOrUpdateGroupRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-or-update-group", fasthttp.MethodPost, reqDto)
	var response dto.CreateOrUpdateGroupRespDto
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
func (client *ManagementClient) CreateGroupsBatch(reqDto *dto.CreateGroupBatchReqDto) *dto.GroupListRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-groups-batch", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) UpdateGroup(reqDto *dto.UpdateGroupReqDto) *dto.GroupSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-group", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) DeleteGroupsBatch(reqDto *dto.DeleteGroupsReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-groups-batch", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) AddGroupMembers(reqDto *dto.AddGroupMembersReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/add-group-members", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) RemoveGroupMembers(reqDto *dto.RemoveGroupMembersReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/remove-group-members", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) ListGroupMembers(reqDto *dto.ListGroupMembersDto) *dto.UserPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-group-members", fasthttp.MethodGet, reqDto)
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
* @param namespace 所属权限分组(权限空间)的 Code
* @param resourceType 资源类型
* @returns AuthorizedResourceListRespDto
 */
func (client *ManagementClient) GetGroupAuthorizedResources(reqDto *dto.GetGroupAuthorizedResourcesDto) *dto.AuthorizedResourceListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-group-authorized-resources", fasthttp.MethodGet, reqDto)
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
* @param code 权限分组(权限空间)内角色的唯一标识符
* @param namespace 所属权限分组(权限空间)的 Code
* @returns RoleSingleRespDto
 */
func (client *ManagementClient) GetRole(reqDto *dto.GetRoleDto) *dto.RoleSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-role", fasthttp.MethodGet, reqDto)
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
* @summary 单个角色批量授权
* @description 通过权限分组内角色 code，分配角色，被分配者可以是用户或部门。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) AssignRole(reqDto *dto.AssignRoleDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/assign-role", fasthttp.MethodPost, reqDto)
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
* @summary 批量分配角色
* @description 批量分配角色，被分配者可以是用户，可以是部门
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) AssignRoleBatch(reqDto *dto.AssignRoleBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/assign-role-batch", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) RevokeRole(reqDto *dto.RevokeRoleDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/revoke-role", fasthttp.MethodPost, reqDto)
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
* @summary 批量移除分配的角色
* @description 批量移除分配的角色，被分配者可以是用户，可以是部门
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) RevokeRoleBatch(reqDto *dto.RevokeRoleBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/revoke-role-batch", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) GetRoleAuthorizedResources(reqDto *dto.GetRoleAuthorizedResourcesDto) *dto.RoleAuthorizedResourcePaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-role-authorized-resources", fasthttp.MethodGet, reqDto)
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
func (client *ManagementClient) ListRoleMembers(reqDto *dto.ListRoleMembersDto) *dto.UserPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-role-members", fasthttp.MethodGet, reqDto)
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
func (client *ManagementClient) ListRoleDepartments(reqDto *dto.ListRoleDepartmentsDto) *dto.RoleDepartmentListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-role-departments", fasthttp.MethodGet, reqDto)
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
* @description 通过权限分组（权限空间）内角色 code，创建角色，可以选择权限分组、角色描述、角色名称等。
* @param requestBody
* @returns RoleSingleRespDto
 */
func (client *ManagementClient) CreateRole(reqDto *dto.CreateRoleDto) *dto.RoleSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-role", fasthttp.MethodPost, reqDto)
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
* @description 获取角色列表，支持分页、支持根据权限分组（权限空间）筛选
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param keywords 用于根据角色的 code 或者名称进行模糊搜索，可选。
* @param namespace 所属权限分组(权限空间)的 code
* @returns RolePaginatedRespDto
 */
func (client *ManagementClient) ListRoles(reqDto *dto.ListRolesDto) *dto.RolePaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-roles", fasthttp.MethodGet, reqDto)
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
* @summary 单个权限分组（权限空间）内删除角色
* @description 单个权限分组（权限空间）内删除角色，可以批量删除。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteRolesBatch(reqDto *dto.DeleteRoleDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-roles-batch", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) CreateRolesBatch(reqDto *dto.CreateRolesBatch) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-roles-batch", fasthttp.MethodPost, reqDto)
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
* @description 通过权限分组(权限空间）内角色新旧 Code，修改角色，可以选择角色名称、角色描述等。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) UpdateRole(reqDto *dto.UpdateRoleDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-role", fasthttp.MethodPost, reqDto)
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
* @summary 跨权限分组（空间）删除角色
* @description 跨权限分组（空间）删除角色，可以批量删除。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteRoles(reqDto *dto.DeleteRoleBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/multiple-namespace-delete-roles-batch", fasthttp.MethodPost, reqDto)
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
* @summary 校验角色 Code 或者名称是否可用
* @description 通过用户池 ID、权限空间 Code和角色 Code,或者用户池 ID、权限空间名称和角色名称查询是否可用。
* @param requestBody
* @returns RoleCheckParamsRespDto
 */
func (client *ManagementClient) CheckParamsNamespace(reqDto *dto.CheckRoleParamsDto) *dto.RoleCheckParamsRespDto {
	b, err := client.SendHttpRequest("/api/v3/check-role-params", fasthttp.MethodPost, reqDto)
	var response dto.RoleCheckParamsRespDto
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
* @summary 获取角色授权列表
* @description 获取角色授权列表。
* @param roleCode 角色 code,只能使用字母、数字和 -_，最多 50 字符
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param query 按角色 Code 或者角色名称查询
* @param namespaceCode 权限空间code
* @param targetType 目标类型，接受用户，部门
* @returns RoleListPageRespDto
 */
func (client *ManagementClient) ListRoleAssignments(reqDto *dto.ListRoleAssignmentsDto) *dto.RoleListPageRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-role-assignments", fasthttp.MethodGet, reqDto)
	var response dto.RoleListPageRespDto
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
* @summary 创建管理员角色
* @description 通过角色 code、角色名称进行创建管理员角色，可以选择角色描述
* @param requestBody
* @returns RoleCheckParamsRespDto
 */
func (client *ManagementClient) CreateAdminRole(reqDto *dto.CreateAdminRoleDto) *dto.RoleCheckParamsRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-admin-role", fasthttp.MethodPost, reqDto)
	var response dto.RoleCheckParamsRespDto
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
* @summary 删除管理员自定义角色
* @description 删除管理员自定义角色，支持批量删除。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteAdminRolesBatch(reqDto *dto.DeleteAdminRoleDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-admin-roles", fasthttp.MethodPost, reqDto)
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
* @summary 获取身份源列表
* @description 获取身份源列表，可以指定 租户 ID 筛选。
* @param tenantId 租户 ID
* @param appId 应用 ID
* @returns ExtIdpListPaginatedRespDto
 */
func (client *ManagementClient) ListExtIdp(reqDto *dto.ListExtIdpDto) *dto.ExtIdpListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-ext-idp", fasthttp.MethodGet, reqDto)
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
* @param appId 应用 ID
* @param type 身份源类型
* @returns ExtIdpDetailSingleRespDto
 */
func (client *ManagementClient) GetExtIdp(reqDto *dto.GetExtIdpDto) *dto.ExtIdpDetailSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-ext-idp", fasthttp.MethodGet, reqDto)
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
func (client *ManagementClient) CreateExtIdp(reqDto *dto.CreateExtIdpDto) *dto.ExtIdpSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-ext-idp", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) UpdateExtIdp(reqDto *dto.UpdateExtIdpDto) *dto.ExtIdpSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-ext-idp", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) DeleteExtIdp(reqDto *dto.DeleteExtIdpDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-ext-idp", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) CreateExtIdpConn(reqDto *dto.CreateExtIdpConnDto) *dto.ExtIdpConnDetailSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-ext-idp-conn", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) UpdateExtIdpConn(reqDto *dto.UpdateExtIdpConnDto) *dto.ExtIdpConnDetailSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-ext-idp-conn", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) DeleteExtIdpConn(reqDto *dto.DeleteExtIdpConnDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-ext-idp-conn", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) ChangeExtIdpConnState(reqDto *dto.ChangeExtIdpConnStateDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/change-ext-idp-conn-state", fasthttp.MethodPost, reqDto)
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
* @summary 租户关联身份源
* @description 租户可以关联或取消关联身份源连接。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) ChangeExtIdpConnAssociationState(reqDto *dto.ChangeExtIdpAssociationStateDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/change-ext-idp-conn-association-state", fasthttp.MethodPost, reqDto)
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
* @summary 租户控制台获取身份源列表
* @description 在租户控制台内获取身份源列表，可以根据 应用 ID 筛选。
* @param tenantId 租户 ID
* @param appId 应用 ID
* @param type 身份源类型
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @returns ExtIdpListPaginatedRespDto
 */
func (client *ManagementClient) ListTenantExtIdp(reqDto *dto.ListTenantExtIdpDto) *dto.ExtIdpListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-tenant-ext-idp", fasthttp.MethodGet, reqDto)
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
* @summary 身份源下应用的连接详情
* @description 在身份源详情页获取应用的连接情况
* @param id 身份源 ID
* @param tenantId 租户 ID
* @param appId 应用 ID
* @param type 身份源类型
* @returns ExtIdpListPaginatedRespDto
 */
func (client *ManagementClient) ExtIdpConnStateByApps(reqDto *dto.ExtIdpConnAppsDto) *dto.ExtIdpListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/ext-idp-conn-apps", fasthttp.MethodGet, reqDto)
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
* @summary 获取用户内置字段列表
* @description 获取用户内置的字段列表
* @returns CustomFieldListRespDto
 */
func (client *ManagementClient) GetUserBaseFields() *dto.CustomFieldListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-base-fields", fasthttp.MethodGet, nil)
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
	* @summary 获取用户内置字段列表
	* @description 获取用户内置的字段列表
	* @param targetType 目标对象类型：
 * - `USER`: 用户
 * - `ROLE`: 角色
 * - `GROUP`: 分组
 * - `DEPARTMENT`: 部门
 * ;该接口暂不支持分组(GROUP)
	* @param dataType 字段类型
	* @param tenantId 租户 ID
	* @param page 当前页数，从 1 开始
	* @param limit 每页数目，最大不能超过 50，默认为 10
	* @param userVisible 用户是否可见
	* @param adminVisible 管理员是否可见
	* @param accessControl 访问控制
	* @param keyword 搜索关键词
	* @param lang 搜索语言
	* @returns ListCistomFieldsResDto
*/
func (client *ManagementClient) ListUserBaseFields(reqDto *dto.ListUserBaseFieldsDto) *dto.ListCistomFieldsResDto {
	b, err := client.SendHttpRequest("/api/v3/list-user-base-fields", fasthttp.MethodGet, reqDto)
	var response dto.ListCistomFieldsResDto
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
* @summary 修改用户内置字段配置
* @description 修改用户内置字段配置，内置字段不允许修改数据类型、唯一性。
* @param requestBody
* @returns CustomFieldListRespDto
 */
func (client *ManagementClient) SetUserBaseFields(reqDto *dto.SetUserBaseFieldsReqDto) *dto.CustomFieldListRespDto {
	b, err := client.SendHttpRequest("/api/v3/set-user-base-fields", fasthttp.MethodPost, reqDto)
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
	* @summary 获取自定义字段列表
	* @description 通过主体类型，获取用户、部门或角色的自定义字段列表。
	* @param targetType 目标对象类型：
 * - `USER`: 用户
 * - `ROLE`: 角色
 * - `GROUP`: 分组
 * - `DEPARTMENT`: 部门
 * ;该接口暂不支持分组(GROUP)
	* @param tenantId 租户 ID
	* @returns CustomFieldListRespDto
*/
func (client *ManagementClient) GetCustomFields(reqDto *dto.GetCustomFieldsDto) *dto.CustomFieldListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-custom-fields", fasthttp.MethodGet, reqDto)
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
	* @summary 获取自定义字段列表
	* @description 通过主体类型，获取用户、部门或角色的自定义字段列表。
	* @param targetType 目标对象类型：
 * - `USER`: 用户
 * - `ROLE`: 角色
 * - `GROUP`: 分组
 * - `DEPARTMENT`: 部门
 * ;该接口暂不支持分组(GROUP)
	* @param dataType 字段类型
	* @param tenantId 租户 ID
	* @param page 当前页数，从 1 开始
	* @param limit 每页数目，最大不能超过 50，默认为 10
	* @param userVisible 用户是否可见
	* @param adminVisible 管理员是否可见
	* @param accessControl 访问控制
	* @param keyword 搜索关键词
	* @param lang 搜索语言
	* @returns ListCistomFieldsResDto
*/
func (client *ManagementClient) ListCustFields(reqDto *dto.ListCustomFieldsDto) *dto.ListCistomFieldsResDto {
	b, err := client.SendHttpRequest("/api/v3/list-custom-fields", fasthttp.MethodGet, reqDto)
	var response dto.ListCistomFieldsResDto
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
* @description 创建/修改用户、部门或角色自定义字段定义，如果传入的 key 不存在则创建，存在则更新。
* @param requestBody
* @returns CustomFieldListRespDto
 */
func (client *ManagementClient) SetCustomFields(reqDto *dto.SetCustomFieldsReqDto) *dto.CustomFieldListRespDto {
	b, err := client.SendHttpRequest("/api/v3/set-custom-fields", fasthttp.MethodPost, reqDto)
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
* @summary 删除自定义字段定义
* @description 删除用户、部门或角色自定义字段定义。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteCustomFields(reqDto *dto.DeleteCustomFieldsReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-custom-fields", fasthttp.MethodPost, reqDto)
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
* @summary 设置自定义字段的值
* @description 给用户、角色或部门设置自定义字段的值，如果存在则更新，不存在则创建。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) SetCustomData(reqDto *dto.SetCustomDataReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/set-custom-data", fasthttp.MethodPost, reqDto)
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
	* @param tenantId 租户 ID
	* @param targetType 目标对象类型：
 * - `USER`: 用户
 * - `ROLE`: 角色
 * - `GROUP`: 分组
 * - `DEPARTMENT`: 部门
 *
	* @param targetIdentifier 目标对象的唯一标志符：
 * - 如果是用户，为用户的 ID，如 `6343b98b7cfxxx9366e9b7c`
 * - 如果是角色，为角色的 code，如 `admin`
 * - 如果是分组，为分组的 code，如 `developer`
 * - 如果是部门，为部门的 ID，如 `6343bafc019xxxx889206c4c`
 *
	* @param namespace 所属权限分组的 code，当 targetType 为角色的时候需要填写，否则可以忽略
	* @returns GetCustomDataRespDto
*/
func (client *ManagementClient) GetCustomData(reqDto *dto.GetCustomDataDto) *dto.GetCustomDataRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-custom-data", fasthttp.MethodGet, reqDto)
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
func (client *ManagementClient) CreateResource(reqDto *dto.CreateResourceDto) *dto.ResourceRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-resource", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) CreateResourcesBatch(reqDto *dto.CreateResourcesBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-resources-batch", fasthttp.MethodPost, reqDto)
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
* @param namespace 所属权限分组(权限空间)的 Code
* @returns ResourceRespDto
 */
func (client *ManagementClient) GetResource(reqDto *dto.GetResourceDto) *dto.ResourceRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-resource", fasthttp.MethodGet, reqDto)
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
* @param namespace 所属权限分组(权限空间)的 Code
* @returns ResourceListRespDto
 */
func (client *ManagementClient) GetResourcesBatch(reqDto *dto.GetResourcesBatchDto) *dto.ResourceListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-resources-batch", fasthttp.MethodGet, reqDto)
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
* @summary 分页获取常规资源列表
* @description 根据筛选条件，分页获取常规资源详情列表。
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param keyword 查询条件
* @param namespaceCodeList 权限空间列表
* @returns CommonResourcePaginatedRespDto
 */
func (client *ManagementClient) ListCommonResource(reqDto *dto.ListCommonResourceDto) *dto.CommonResourcePaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-common-resource", fasthttp.MethodGet, reqDto)
	var response dto.CommonResourcePaginatedRespDto
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
* @param namespace 所属权限分组(权限空间)的 Code
* @param type 资源类型
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @returns ResourcePaginatedRespDto
 */
func (client *ManagementClient) ListResources(reqDto *dto.ListResourcesDto) *dto.ResourcePaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-resources", fasthttp.MethodGet, reqDto)
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
func (client *ManagementClient) UpdateResource(reqDto *dto.UpdateResourceDto) *dto.ResourceRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-resource", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) DeleteResource(reqDto *dto.DeleteResourceDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-resource", fasthttp.MethodPost, reqDto)
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
func (client *ManagementClient) DeleteResourcesBatch(reqDto *dto.DeleteResourcesBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-resources-batch", fasthttp.MethodPost, reqDto)
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
* @description 通过资源id批量删除资源
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) BatchDeleteCommonResource(reqDto *dto.DeleteCommonResourcesBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-common-resources-batch", fasthttp.MethodPost, reqDto)
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
* @summary 关联/取消关联应用资源到租户
* @description 通过资源唯一标识以及权限分组，关联或取消关联资源到租户
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) AssociateTenantResource(reqDto *dto.AssociateTenantResourceDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/associate-tenant-resource", fasthttp.MethodPost, reqDto)
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
* @description 创建权限分组，可以设置权限分组名称、Code 和描述信息。
* @param requestBody
* @returns NamespaceRespDto
 */
func (client *ManagementClient) CreateNamespace(reqDto *dto.CreateNamespaceDto) *dto.NamespaceRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-namespace", fasthttp.MethodPost, reqDto)
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
* @description 批量创建权限分组，可以分别设置权限分组名称、Code 和描述信息。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) CreateNamespacesBatch(reqDto *dto.CreateNamespacesBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-namespaces-batch", fasthttp.MethodPost, reqDto)
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
* @description 通过权限分组唯一标志符(Code)，获取权限分组详情。
* @param code 权限分组唯一标志符
* @returns NamespaceRespDto
 */
func (client *ManagementClient) GetNamespace(reqDto *dto.GetNamespaceDto) *dto.NamespaceRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-namespace", fasthttp.MethodGet, reqDto)
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
* @description 分别通过权限分组唯一标志符(Code)，批量获取权限分组详情。
* @param codeList 权限分组 code 列表，批量可以使用逗号分隔
* @returns NamespaceListRespDto
 */
func (client *ManagementClient) GetNamespacesBatch(reqDto *dto.GetNamespacesBatchDto) *dto.NamespaceListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-namespaces-batch", fasthttp.MethodGet, reqDto)
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
* @description 修改权限分组信息，可以修改名称、描述信息以及新的唯一标志符(NewCode)。
* @param requestBody
* @returns UpdateNamespaceRespDto
 */
func (client *ManagementClient) UpdateNamespace(reqDto *dto.UpdateNamespaceDto) *dto.UpdateNamespaceRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-namespace", fasthttp.MethodPost, reqDto)
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
* @description 通过权限分组唯一标志符(Code)，删除权限分组信息。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteNamespace(reqDto *dto.DeleteNamespaceDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-namespace", fasthttp.MethodPost, reqDto)
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
* @description 分别通过权限分组唯一标志符(Code)，批量删除权限分组。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteNamespacesBatch(reqDto *dto.DeleteNamespacesBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-namespaces-batch", fasthttp.MethodPost, reqDto)
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
* @summary 分页获取权限分组列表
* @description 根据筛选条件，分页获取权限分组列表。
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param keywords 搜索权限分组 Code
* @returns NamespaceListPaginatedRespDto
 */
func (client *ManagementClient) ListNamespaces(reqDto *dto.ListNamespacesDto) *dto.NamespaceListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-namespaces", fasthttp.MethodGet, reqDto)
	var response dto.NamespaceListPaginatedRespDto
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
* @summary 分页权限分组下所有的角色列表
* @description 根据筛选条件，分页获取权限分组下所有的角色列表。
* @param code 权限分组唯一标志符
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param keywords 角色 Code 或者名称
* @returns NamespaceRolesListPaginatedRespDto
 */
func (client *ManagementClient) ListNamespaceRoles(reqDto *dto.ListNamespaceRolesDto) *dto.NamespaceRolesListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-namespace-roles", fasthttp.MethodGet, reqDto)
	var response dto.NamespaceRolesListPaginatedRespDto
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
func (client *ManagementClient) AuthorizeResources(reqDto *dto.AuthorizeResourcesDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/authorize-resources", fasthttp.MethodPost, reqDto)
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
	* @param targetType 目标对象类型：
 * - `USER`: 用户
 * - `ROLE`: 角色
 * - `GROUP`: 分组
 * - `DEPARTMENT`: 部门
 *
	* @param targetIdentifier 目标对象的唯一标志符：
 * - 如果是用户，为用户的 ID，如 `6343b98b7cfxxx9366e9b7c`
 * - 如果是角色，为角色的 code，如 `admin`
 * - 如果是分组，为分组的 code，如 `developer`
 * - 如果是部门，为部门的 ID，如 `6343bafc019xxxx889206c4c`
 *
	* @param namespace 所属权限分组(权限空间)的 Code
	* @param resourceType 限定资源类型，如数据、API、按钮、菜单
	* @param resourceList 限定查询的资源列表，如果指定，只会返回所指定的资源列表。
	* @param withDenied 是否获取被拒绝的资源
	* @returns AuthorizedResourcePaginatedRespDto
*/
func (client *ManagementClient) GetAuthorizedResources(reqDto *dto.GetAuthorizedResourcesDto) *dto.AuthorizedResourcePaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-authorized-resources", fasthttp.MethodGet, reqDto)
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
func (client *ManagementClient) IsActionAllowed(reqDto *dto.IsActionAllowedDto) *dto.IsActionAllowedRespDtp {
	b, err := client.SendHttpRequest("/api/v3/is-action-allowed", fasthttp.MethodPost, reqDto)
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

/*
* @summary 获取资源被授权的主体
* @description 获取资源被授权的主体
* @param requestBody
* @returns GetResourceAuthorizedTargetRespDto
 */
func (client *ManagementClient) GetResourceAuthorizedTargets(reqDto *dto.GetResourceAuthorizedTargetsDto) *dto.GetResourceAuthorizedTargetRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-resource-authorized-targets", fasthttp.MethodPost, reqDto)
	var response dto.GetResourceAuthorizedTargetRespDto
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
* @summary 获取用户行为日志
* @description 可以选择请求 ID、客户端 IP、用户 ID、应用 ID、开始时间戳、请求是否成功、分页参数来获取用户行为日志
* @param requestBody
* @returns UserActionLogRespDto
 */
func (client *ManagementClient) GetUserActionLogs(reqDto *dto.GetUserActionLogsDto) *dto.UserActionLogRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-action-logs", fasthttp.MethodPost, reqDto)
	var response dto.UserActionLogRespDto
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
* @summary 获取管理员操作日志
* @description 可以选择请求 ID、客户端 IP、操作类型、资源类型、管理员用户 ID、请求是否成功、开始时间戳、结束时间戳、分页来获取管理员操作日志接口
* @param requestBody
* @returns AdminAuditLogRespDto
 */
func (client *ManagementClient) GetAdminAuditLogs(reqDto *dto.GetAdminAuditLogsDto) *dto.AdminAuditLogRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-admin-audit-logs", fasthttp.MethodPost, reqDto)
	var response dto.AdminAuditLogRespDto
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
* @summary 获取邮件模版列表
* @description 获取邮件模版列表
* @returns GetEmailTemplatesRespDto
 */
func (client *ManagementClient) GetEmailTemplates() *dto.GetEmailTemplatesRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-email-templates", fasthttp.MethodGet, nil)
	var response dto.GetEmailTemplatesRespDto
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
* @summary 修改邮件模版
* @description 修改邮件模版
* @param requestBody
* @returns EmailTemplateSingleItemRespDto
 */
func (client *ManagementClient) UpdateEmailTemplate(reqDto *dto.UpdateEmailTemplateDto) *dto.EmailTemplateSingleItemRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-email-template", fasthttp.MethodPost, reqDto)
	var response dto.EmailTemplateSingleItemRespDto
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
* @summary 预览邮件模版
* @description 预览邮件模版
* @param requestBody
* @returns PreviewEmailTemplateRespDto
 */
func (client *ManagementClient) PreviewEmailTemplate(reqDto *dto.PreviewEmailTemplateDto) *dto.PreviewEmailTemplateRespDto {
	b, err := client.SendHttpRequest("/api/v3/preview-email-template", fasthttp.MethodPost, reqDto)
	var response dto.PreviewEmailTemplateRespDto
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
* @summary 获取第三方邮件服务配置
* @description 获取第三方邮件服务配置
* @returns EmailProviderRespDto
 */
func (client *ManagementClient) GetEmailProvider() *dto.EmailProviderRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-email-provider", fasthttp.MethodGet, nil)
	var response dto.EmailProviderRespDto
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
* @summary 配置第三方邮件服务
* @description 配置第三方邮件服务
* @param requestBody
* @returns EmailProviderRespDto
 */
func (client *ManagementClient) ConfigEmailProvider(reqDto *dto.ConfigEmailProviderDto) *dto.EmailProviderRespDto {
	b, err := client.SendHttpRequest("/api/v3/config-email-provider", fasthttp.MethodPost, reqDto)
	var response dto.EmailProviderRespDto
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
* @summary 获取应用详情
* @description 通过应用 ID，获取应用详情。
* @param appId 应用 ID
* @returns ApplicationSingleRespDto
 */
func (client *ManagementClient) GetApplication(reqDto *dto.GetApplicationDto) *dto.ApplicationSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-application", fasthttp.MethodGet, reqDto)
	var response dto.ApplicationSingleRespDto
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
* @summary 主体授权详情
* @description 主体授权详情
* @param targetId 主体 id
* @param targetType 主体类型
* @param appId 应用 ID
* @returns GetSubjectAuthRespDto
 */
func (client *ManagementClient) DetailAuthSubject(reqDto *dto.GetSubjectAuthDetailDto) *dto.GetSubjectAuthRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-subject-auth-detail", fasthttp.MethodGet, reqDto)
	var response dto.GetSubjectAuthRespDto
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
* @summary 主体授权列表
* @description 主体授权列表
* @param requestBody
* @returns ListApplicationSubjectRespDto
 */
func (client *ManagementClient) ListAuthSubject(reqDto *dto.ListAuthSubjectDto) *dto.ListApplicationSubjectRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-subject-auth", fasthttp.MethodPost, reqDto)
	var response dto.ListApplicationSubjectRespDto
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
* @summary 应用授权列表
* @description 应用授权列表
* @param requestBody
* @returns ListApplicationAuthPaginatedRespDto
 */
func (client *ManagementClient) ListAuthApplication(reqDto *dto.ListApplicationAuthDto) *dto.ListApplicationAuthPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-applications-auth", fasthttp.MethodPost, reqDto)
	var response dto.ListApplicationAuthPaginatedRespDto
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
* @summary 更新授权开关
* @description 更新授权开关
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) EnabledAuth(reqDto *dto.UpdateAuthEnabledDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-auth-enabled", fasthttp.MethodPost, reqDto)
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
* @summary 批量删除应用授权
* @description 批量删除应用授权
* @param authIds 授权 ID
* @returns IsSuccessRespDto
 */
/*
* @summary 获取应用列表
* @description 获取应用列表
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param isIntegrateApp 是否为集成应用
* @param isSelfBuiltApp 是否为自建应用
* @param ssoEnabled 是否开启单点登录
* @param keywords 模糊搜索字符串
* @param all 搜索应用，true：搜索所有应用, 默认为 false
* @returns ApplicationPaginatedRespDto
 */
func (client *ManagementClient) ListApplications(reqDto *dto.ListApplicationsDto) *dto.ApplicationPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-applications", fasthttp.MethodGet, reqDto)
	var response dto.ApplicationPaginatedRespDto
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
* @summary 获取应用简单信息
* @description 通过应用 ID，获取应用简单信息。
* @param appId 应用 ID
* @returns ApplicationSimpleInfoSingleRespDto
 */
func (client *ManagementClient) GetApplicationSimpleInfo(reqDto *dto.GetApplicationSimpleInfoDto) *dto.ApplicationSimpleInfoSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-application-simple-info", fasthttp.MethodGet, reqDto)
	var response dto.ApplicationSimpleInfoSingleRespDto
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
* @summary 获取应用简单信息列表
* @description 获取应用简单信息列表
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param isIntegrateApp 是否为集成应用
* @param isSelfBuiltApp 是否为自建应用
* @param ssoEnabled 是否开启单点登录
* @param keywords 模糊搜索字符串
* @returns ApplicationSimpleInfoPaginatedRespDto
 */
func (client *ManagementClient) ListApplicationSimpleInfo(reqDto *dto.ListApplicationSimpleInfoDto) *dto.ApplicationSimpleInfoPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-application-simple-info", fasthttp.MethodGet, reqDto)
	var response dto.ApplicationSimpleInfoPaginatedRespDto
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
* @summary 创建应用
* @description 创建应用
* @param requestBody
* @returns CreateApplicationRespDto
 */
func (client *ManagementClient) CreateApplication(reqDto *dto.CreateApplicationDto) *dto.CreateApplicationRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-application", fasthttp.MethodPost, reqDto)
	var response dto.CreateApplicationRespDto
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
* @summary 删除应用
* @description 通过应用 ID，删除应用。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteApplication(reqDto *dto.DeleteApplicationDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-application", fasthttp.MethodPost, reqDto)
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
* @summary 获取应用密钥
* @description 获取应用密钥
* @param appId 应用 ID
* @returns GetApplicationSecretRespDto
 */
func (client *ManagementClient) GetApplicationSecret(reqDto *dto.GetApplicationSecretDto) *dto.GetApplicationSecretRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-application-secret", fasthttp.MethodGet, reqDto)
	var response dto.GetApplicationSecretRespDto
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
* @summary 刷新应用密钥
* @description 刷新应用密钥
* @param requestBody
* @returns RefreshApplicationSecretRespDto
 */
func (client *ManagementClient) RefreshApplicationSecret(reqDto *dto.RefreshApplicationSecretDto) *dto.RefreshApplicationSecretRespDto {
	b, err := client.SendHttpRequest("/api/v3/refresh-application-secret", fasthttp.MethodPost, reqDto)
	var response dto.RefreshApplicationSecretRespDto
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
* @summary 获取应用当前登录用户
* @description 获取应用当前处于登录状态的用户
* @param requestBody
* @returns UserPaginatedRespDto
 */
func (client *ManagementClient) ListApplicationActiveUsers(reqDto *dto.ListApplicationActiveUsersDto) *dto.UserPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-application-active-users", fasthttp.MethodPost, reqDto)
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
* @summary 获取应用默认访问授权策略
* @description 获取应用默认访问授权策略
* @param appId 应用 ID
* @returns GetApplicationPermissionStrategyRespDto
 */
func (client *ManagementClient) GetApplicationPermissionStrategy(reqDto *dto.GetApplicationPermissionStrategyDto) *dto.GetApplicationPermissionStrategyRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-application-permission-strategy", fasthttp.MethodGet, reqDto)
	var response dto.GetApplicationPermissionStrategyRespDto
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
* @summary 更新应用默认访问授权策略
* @description 更新应用默认访问授权策略
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) UpdateApplicationPermissionStrategy(reqDto *dto.UpdateApplicationPermissionStrategyDataDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-application-permission-strategy", fasthttp.MethodPost, reqDto)
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
* @summary 授权应用访问权限
* @description 给用户、分组、组织或角色授权应用访问权限，如果用户、分组、组织或角色不存在，则跳过，进行下一步授权，不返回报错
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) AuthorizeApplicationAccess(reqDto *dto.AuthorizeApplicationAccessDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/authorize-application-access", fasthttp.MethodPost, reqDto)
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
* @summary 删除应用访问授权记录
* @description 取消给用户、分组、组织或角色的应用访问权限授权,如果传入数据不存在，则返回数据不报错处理。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) RevokeApplicationAccess(reqDto *dto.RevokeApplicationAccessDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/revoke-application-access", fasthttp.MethodPost, reqDto)
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
* @summary 检测域名是否可用
* @description 检测域名是否可用于创建新应用或更新应用域名
* @param requestBody
* @returns CheckDomainAvailableSecretRespDto
 */
func (client *ManagementClient) CheckDomainAvailable(reqDto *dto.CheckDomainAvailable) *dto.CheckDomainAvailableSecretRespDto {
	b, err := client.SendHttpRequest("/api/v3/check-domain-available", fasthttp.MethodPost, reqDto)
	var response dto.CheckDomainAvailableSecretRespDto
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
* @summary 获取租户应用列表
* @description 获取应用列表，可以指定 租户 ID 筛选。
* @param page 获取应用列表的页码
* @param limit 每页获取的应用数量
* @param keywords 搜索关键字
* @param ssoEnabled 应用是否加入单点登录
* @returns TenantApplicationListPaginatedRespDto
 */
func (client *ManagementClient) ListTenantApplications(reqDto *dto.ListTenantApplicationsDto) *dto.TenantApplicationListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-tenant-applications", fasthttp.MethodGet, reqDto)
	var response dto.TenantApplicationListPaginatedRespDto
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
* @summary 更新应用登录页配置
* @description 通过应用 ID 更新登录页配置。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) UpdateLoginPageConfig(reqDto *dto.UpdateLoginConfigDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-login-page-config", fasthttp.MethodPost, reqDto)
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
* @summary 获取用户池租户配置信息
* @description 根据用户池 ID 获取用户池多租户配置信息
* @returns UserPoolTenantConfigDtoRespDto
 */
func (client *ManagementClient) UserpollTenantConfig() *dto.UserPoolTenantConfigDtoRespDto {
	b, err := client.SendHttpRequest("/api/v3/userpool-tenant-config", fasthttp.MethodGet, nil)
	var response dto.UserPoolTenantConfigDtoRespDto
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
* @summary 更新用户池租户配置信息
* @description 更新用户池多租户配置内登录信息
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) UpdateUserPoolTenantConfig(reqDto *dto.UpdateUserPoolTenantLoginConfigDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-userpool-tenant-config", fasthttp.MethodPost, reqDto)
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
* @summary 更新租户控制台扫码登录状态
* @description 更新租户控制台扫码登录状态
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) UpdateTenantQrCodeState(reqDto *dto.UpdateTenantAppqrcodeState) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-userpool-tenant-appqrcode-state", fasthttp.MethodPost, reqDto)
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
* @summary 设置用户池多租户身份源连接
* @description 设置用户池多租户身份源连接，支持同时设置多个身份源连接，支持设置连接和取消连接
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) ChangeUserpoolTenanExtIdpConnState(reqDto *dto.ChangeUserPoolTenantExtIdpConnDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/change-userpool-tenant-ext-idp-conn-state", fasthttp.MethodPost, reqDto)
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
* @summary 修改应用多因素认证配置
* @description 传入 MFA 认证因素列表进行开启或关闭
* @param requestBody
* @returns MFASettingsRespDto
 */
func (client *ManagementClient) UpdateApplicationMfaSettings(reqDto *dto.UpdateApplicationMfaSettingsDto) *dto.MFASettingsRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-application-mfa-settings", fasthttp.MethodPost, reqDto)
	var response dto.MFASettingsRespDto
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
	* @summary 获取应用下用户 MFA 触发数据
	* @description 获取应用下用户 MFA 触发数据。
	* @param appId 所属应用 ID
	* @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
	* @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 * - `sync_relation`: 用户的外部身份源信息，格式为 `<provier>:<userIdInIdp>`，其中 `<provier>` 为同步身份源类型，如 wechatwork, lark；`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`lark:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
	* @returns GetMapInfoRespDto
*/
func (client *ManagementClient) GetMfaTriggerData(reqDto *dto.GetMfaTriggerDataDto) *dto.GetMapInfoRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-mfa-trigger-data", fasthttp.MethodGet, reqDto)
	var response dto.GetMapInfoRespDto
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
* @summary 创建 ASA 账号
* @description 在某一应用下创建 ASA 账号
* @param requestBody
* @returns AsaAccountSingleRespDto
 */
func (client *ManagementClient) CreateAsaAccount(reqDto *dto.CreateAsaAccountDto) *dto.AsaAccountSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-asa-account", fasthttp.MethodPost, reqDto)
	var response dto.AsaAccountSingleRespDto
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
* @summary 批量创建 ASA 账号
* @description 在某一应用下批量创建 ASA 账号
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) CreateAsaAccountBatch(reqDto *dto.CreateAsaAccountsBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-asa-accounts-batch", fasthttp.MethodPost, reqDto)
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
* @summary 更新 ASA 账号
* @description 更新某个 ASA 账号信息
* @param requestBody
* @returns AsaAccountSingleRespDto
 */
func (client *ManagementClient) UpdateAsaAccount(reqDto *dto.UpdateAsaAccountDto) *dto.AsaAccountSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-asa-account", fasthttp.MethodPost, reqDto)
	var response dto.AsaAccountSingleRespDto
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
* @summary 获取 ASA 账号列表
* @description 分页获取某一应用下的 ASA 账号列表
* @param appId 所属应用 ID
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @returns AsaAccountPaginatedRespDto
 */
func (client *ManagementClient) ListAsaAccount(reqDto *dto.ListAsaAccountsDto) *dto.AsaAccountPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-asa-accounts", fasthttp.MethodGet, reqDto)
	var response dto.AsaAccountPaginatedRespDto
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
* @summary 获取 ASA 账号
* @description 根据 ASA 账号 ID 获取账号详细信息
* @param appId 所属应用 ID
* @param accountId ASA 账号 ID
* @returns AsaAccountSingleRespDto
 */
func (client *ManagementClient) GetAsaAccount(reqDto *dto.GetAsaAccountDto) *dto.AsaAccountSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-asa-account", fasthttp.MethodGet, reqDto)
	var response dto.AsaAccountSingleRespDto
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
* @summary 批量获取 ASA 账号
* @description 根据 ASA 账号 ID 列表批量获取账号详细信息
* @param requestBody
* @returns AsaAccountListRespDto
 */
func (client *ManagementClient) GetAsaAccountBatch(reqDto *dto.GetAsaAccountBatchDto) *dto.AsaAccountListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-asa-accounts-batch", fasthttp.MethodPost, reqDto)
	var response dto.AsaAccountListRespDto
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
* @summary 删除 ASA 账号
* @description 通过 ASA 账号 ID 删除 ASA 账号
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteAsaAccount(reqDto *dto.DeleteAsaAccountDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-asa-account", fasthttp.MethodPost, reqDto)
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
* @summary 批量删除 ASA 账号
* @description 通过 ASA 账号 ID 批量删除 ASA 账号
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteAsaAccountBatch(reqDto *dto.DeleteAsaAccountBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-asa-accounts-batch", fasthttp.MethodPost, reqDto)
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
* @summary 分配 ASA 账号
* @description 分配 ASA 账号给用户、组织、分组或角色
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) AssignAsaAccount(reqDto *dto.AssignAsaAccountsDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/assign-asa-account", fasthttp.MethodPost, reqDto)
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
* @summary 取消分配 ASA 账号
* @description 取消分配给用户、组织、分组或角色的  ASA 账号
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) UnassignAsaAccount(reqDto *dto.AssignAsaAccountsDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/unassign-asa-account", fasthttp.MethodPost, reqDto)
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
* @summary 获取 ASA 账号分配的主体列表
* @description 根据 ASA 账号 ID 分页获取账号被分配的主体列表
* @param appId 所属应用 ID
* @param accountId ASA 账号 ID
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @returns GetAsaAccountAssignedTargetRespDto
 */
func (client *ManagementClient) GetAsaAccountAssignedTargets(reqDto *dto.GetAsaAccountAssignedTargetsDto) *dto.GetAsaAccountAssignedTargetRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-asa-account-assigned-targets", fasthttp.MethodGet, reqDto)
	var response dto.GetAsaAccountAssignedTargetRespDto
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
	* @summary 获取主体被分配的 ASA 账号
	* @description 根据主体类型和标识获取直接分配给主体的 ASA 账号
	* @param appId 所属应用 ID
	* @param targetType 目标对象类型：
 * - `USER`: 用户
 * - `ROLE`: 角色
 * - `GROUP`: 分组
 * - `DEPARTMENT`: 部门
 *
	* @param targetIdentifier 目标对象的唯一标志符：
 * - 如果是用户，为用户的 ID，如 `6343b98b7cfxxx9366e9b7c`
 * - 如果是角色，为角色的 code，如 `admin`
 * - 如果是分组，为分组的 code，如 `developer`
 * - 如果是部门，为部门的 ID，如 `6343bafc019xxxx889206c4c`
 *
	* @returns AsaAccountSingleNullableRespDto
*/
func (client *ManagementClient) GetAssignedAccount(reqDto *dto.GetAssignedAccountDto) *dto.AsaAccountSingleNullableRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-assigned-account", fasthttp.MethodGet, reqDto)
	var response dto.AsaAccountSingleNullableRespDto
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
* @summary 获取安全配置
* @description 无需传参获取安全配置
* @returns SecuritySettingsRespDto
 */
func (client *ManagementClient) GetSecuritySettings() *dto.SecuritySettingsRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-security-settings", fasthttp.MethodGet, nil)
	var response dto.SecuritySettingsRespDto
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
* @summary 修改安全配置
* @description 可选安全域、Authing Token 有效时间（秒）、验证码长度、验证码尝试次数、用户修改邮箱的安全策略、用户修改手机号的安全策略、Cookie 过期时间设置、是否禁止用户注册、频繁注册检测配置、验证码注册后是否要求用户设置密码、未验证的邮箱登录时是否禁止登录并发送认证邮件、用户自助解锁配置、Authing 登录页面是否开启登录账号选择、APP 扫码登录安全配置进行修改安全配置
* @param requestBody
* @returns SecuritySettingsRespDto
 */
func (client *ManagementClient) UpdateSecuritySettings(reqDto *dto.UpdateSecuritySettingsDto) *dto.SecuritySettingsRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-security-settings", fasthttp.MethodPost, reqDto)
	var response dto.SecuritySettingsRespDto
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
* @summary 获取全局多因素认证配置
* @description 无需传参获取全局多因素认证配置
* @returns MFASettingsRespDto
 */
func (client *ManagementClient) GetGlobalMfaSettings() *dto.MFASettingsRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-global-mfa-settings", fasthttp.MethodGet, nil)
	var response dto.MFASettingsRespDto
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
* @summary 修改全局多因素认证配置
* @description 传入 MFA 认证因素列表进行开启,
* @param requestBody
* @returns MFASettingsRespDto
 */
func (client *ManagementClient) UpdateGlobalMfaSettings(reqDto *dto.MFASettingsDto) *dto.MFASettingsRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-global-mfa-settings", fasthttp.MethodPost, reqDto)
	var response dto.MFASettingsRespDto
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
* @summary 创建租户
* @param requestBody
* @returns CreateTenantRespDto
 */
func (client *ManagementClient) CreateTenant(reqDto *dto.CreateTenantDto) *dto.CreateTenantRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-tenant", fasthttp.MethodPost, reqDto)
	var response dto.CreateTenantRespDto
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
* @summary 更新租户数据
* @description 此接口用于更新租户基本信息。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) UpdateTenant(reqDto *dto.UpdateTenantDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-tenant", fasthttp.MethodPost, reqDto)
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
* @summary 删除租户
* @description 此接口用于删除租户。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteTenant(reqDto *dto.DeleteTenantDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-tenant", fasthttp.MethodPost, reqDto)
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
* @summary 获取/搜索租户列表
* @description 此接口用于获取租户列表，支持模糊搜索。
* @param keywords 搜索关键字
* @param withMembersCount 是否增加返回租户成员统计
* @param withAppDetail 增加返回租户下 app 简单信息
* @param withCreatorDetail 增加返回租户下创建者简单信息
* @param withSourceAppDetail 增加返回租户下来源 app 简单信息
* @param page 页码
* @param limit 每页获取的数据量
* @param source 租户来源
* @returns TenantListPaginatedRespDto
 */
func (client *ManagementClient) ListTenants(reqDto *dto.ListTenantsDto) *dto.TenantListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-tenants", fasthttp.MethodGet, reqDto)
	var response dto.TenantListPaginatedRespDto
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
* @summary 获取租户一点点的信息
* @description 根据租户 ID 获取租户一点点的详情
* @param tenantId 租户 ID
* @param withMembersCount 是否增加返回租户成员统计
* @param withAppDetail 增加返回租户关联应用简单信息
* @param withCreatorDetail 增加返回租户下创建者简单信息
* @param withSourceAppDetail 增加返回租户来源应用简单信息
* @returns TenantSingleRespDto
 */
func (client *ManagementClient) GetTenantLittleInfo(reqDto *dto.GetTenantLittleInfoDto) *dto.TenantSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-tenant-little-info", fasthttp.MethodGet, reqDto)
	var response dto.TenantSingleRespDto
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
* @summary 获取租户详情
* @description 根据租户 ID 获取租户详情
* @param tenantId 租户 ID
* @param withMembersCount 是否增加返回租户成员统计
* @param withAppDetail 增加返回租户关联应用简单信息
* @param withCreatorDetail 增加返回租户下创建者简单信息
* @param withSourceAppDetail 增加返回租户来源应用简单信息
* @returns TenantSingleRespDto
 */
func (client *ManagementClient) GetTenant(reqDto *dto.GetTenantDto) *dto.TenantSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-tenant", fasthttp.MethodGet, reqDto)
	var response dto.TenantSingleRespDto
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
* @summary 导入租户
* @description 此接口用于 Excel 导入租户。
* @param requestBody
* @returns ImportTenantRespDto
 */
func (client *ManagementClient) ImportTenant(reqDto *dto.ImportTenantDto) *dto.ImportTenantRespDto {
	b, err := client.SendHttpRequest("/api/v3/import-tenant", fasthttp.MethodPost, reqDto)
	var response dto.ImportTenantRespDto
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
* @summary 导入租户历史
* @description 此接口用于 Excel 导入租户的历史查询。
* @param page 页码
* @param limit 每页获取的数据量
* @returns ImportTenantHistoryRespDto
 */
func (client *ManagementClient) ImportTenantHistory(reqDto *dto.ImportTenantHistoryDto) *dto.ImportTenantHistoryRespDto {
	b, err := client.SendHttpRequest("/api/v3/import-tenant-history", fasthttp.MethodGet, reqDto)
	var response dto.ImportTenantHistoryRespDto
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
* @summary 导入租户通知用户列表
* @description 此接口用于查询导入租户通知用户列表。
* @param importId 导入记录 id
* @param page 页码
* @param limit 每页获取的数据量
* @returns ImportTenantNotifyUserRespDto
 */
func (client *ManagementClient) ImportTenantNotifyUser(reqDto *dto.ImportTenantNotifyUserDto) *dto.ImportTenantNotifyUserRespDto {
	b, err := client.SendHttpRequest("/api/v3/import-tenant-notify-user", fasthttp.MethodGet, reqDto)
	var response dto.ImportTenantNotifyUserRespDto
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
* @summary 导入租户通知邮箱用户
* @description 此接口用于批量发送邮件通知。
* @param requestBody
* @returns SendEmailBatchDataDto
 */
func (client *ManagementClient) SendEmailBatch(reqDto *dto.SendManyTenantEmailDto) *dto.SendEmailBatchDataDto {
	b, err := client.SendHttpRequest("/api/v3/send-email-batch", fasthttp.MethodPost, reqDto)
	var response dto.SendEmailBatchDataDto
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
* @summary 导入租户短信通知用户
* @description 此接口用于批量发送短信通知。
* @param requestBody
* @returns SendSmsBatchRespDto
 */
func (client *ManagementClient) SendSmsBatch(reqDto *dto.SendManyTenantSmsDto) *dto.SendSmsBatchRespDto {
	b, err := client.SendHttpRequest("/api/v3/send-sms-batch", fasthttp.MethodPost, reqDto)
	var response dto.SendSmsBatchRespDto
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
* @summary 获取租户管理员列表
* @description 此接口用于获取租户成员列表，支持模糊搜索。
* @param requestBody
* @returns TenantUserListPaginatedRespDto
 */
func (client *ManagementClient) ListTenantAdmin(reqDto *dto.ListTenantAdminDto) *dto.TenantUserListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-tenant-admin", fasthttp.MethodPost, reqDto)
	var response dto.TenantUserListPaginatedRespDto
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
* @summary 设置租户管理员
* @description 此接口用于根据用户 ID 或租户成员 ID 设置租户管理员。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) SetTenantAdmin(reqDto *dto.RemoveTenantUsersDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/set-tenant-admin", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 取消设置租户管理员
* @description 此接口用于根据用户 ID 或租户成员 ID 取消设置租户管理员。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) DeleteTenantAdmin(reqDto *dto.GetTenantUserDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/delete-tenant-admin", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 批量移除租户成员
* @description 此接口用于根据用户 ID 或租户成员 ID 批量移除租户成员。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) DeleteTenantUser(reqDto *dto.RemoveTenantUsersDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/delete-tenant-user", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 生成一个邀请租户成员的链接
* @description 此接口用于生成一个邀请租户成员的链接。appId 为用户注册成功后要访问的应用 ID
* @param requestBody
* @returns InviteTenantUsersRespDto
 */
func (client *ManagementClient) GenerateInviteTenantUserLink(reqDto *dto.GenerateInviteTenantUserLink) *dto.InviteTenantUsersRespDto {
	b, err := client.SendHttpRequest("/api/v3/generate-invite-tenant-user-link", fasthttp.MethodPost, reqDto)
	var response dto.InviteTenantUsersRespDto
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
* @summary 获取可访问的租户控制台列表
* @description 根据用户 ID 获取可访问的租户控制台列表
* @param keywords 搜索关键字
* @param page 页码
* @param limit 每页获取的数据量
* @returns InviteTenantUserRecordListRespDto
 */
func (client *ManagementClient) ListInviteTennatUserRecords(reqDto *dto.ListInviteTenantUserRecordsDto) *dto.InviteTenantUserRecordListRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-invite-tenant-user-records", fasthttp.MethodGet, reqDto)
	var response dto.InviteTenantUserRecordListRespDto
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
* @summary 获取多租户管理员用户列表
* @description 根据用户池 ID 获取某个用户池内拥有多租户管理权限的用户列表
* @param keywords 搜索关键字
* @param page 页码
* @param limit 每页获取的数据量
* @returns MultipleTenantAdminPaginatedRespDto
 */
func (client *ManagementClient) ListMultipleTenantAdmin(reqDto *dto.ListMultipleTenantAdminsDto) *dto.MultipleTenantAdminPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-multiple-tenant-admins", fasthttp.MethodGet, reqDto)
	var response dto.MultipleTenantAdminPaginatedRespDto
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
* @summary 创建多租户管理员用户
* @description 根据用户 ID 创建一个用户池内拥有多租户管理权限的用户
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) CreateMultipleTenantAdmin(reqDto *dto.CreateMultipleTenantAdminDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/create-multiple-tenant-admin", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 获取多租户管理员用户列表
* @description 根据用户池 ID 获取某个用户池内拥有多租户管理权限的用户列表
* @param userId 用户 ID
* @returns MultipleTenantAdminPaginatedRespDto
 */
func (client *ManagementClient) GetMultipleTenantAdmin(reqDto *dto.GetMultipleTenantAdminDto) *dto.MultipleTenantAdminPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-multiple-tenant-admin", fasthttp.MethodGet, reqDto)
	var response dto.MultipleTenantAdminPaginatedRespDto
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
* @summary 更新多租户管理员用户
* @description 根据用户 ID 更新一个用户池内拥有多租户管理权限的用户的信息
* @param userId 用户 ID
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) UpdateMultipleTenantAdmin(reqDto *dto.UpdateMultipleTenantAdminDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/update-multiple-tenant-admin", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 删除多租户管理员用户
* @description 根据用户 ID 删除一个用户池内拥有多租户管理权限的用户
* @param requestBody
* @param userId 用户 ID（请使用 body 中的 userId）
* @returns CommonResponseDto
 */
func (client *ManagementClient) DeleteMultipleTenantAdmin(reqDto *dto.DeleteMultipleTenantAdminBodyDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/delete-multiple-tenant-admin", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 获取协作管理员用户列表
* @description 根据用户池 ID 获取某个用户池内拥有协作管理员能力的用户列表
* @param keywords 搜索关键字
* @param external 是否外部
* @param page 页码
* @param limit 每页获取的数据量
* @returns TenantCooperatorPaginatedRespDto
 */
func (client *ManagementClient) ListTenantCooperators(reqDto *dto.ListTenantCooperatorsDto) *dto.TenantCooperatorPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-tenant-cooperators", fasthttp.MethodGet, reqDto)
	var response dto.TenantCooperatorPaginatedRespDto
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
* @summary 获取一个协调管理员
* @description 根据用户池 ID 获取某个协调管理员的列表
* @param userId 用户 ID
* @returns TenantCooperatorSingleRespDto
 */
func (client *ManagementClient) GetTenantCooperator(reqDto *dto.GetTenantCooperatorDto) *dto.TenantCooperatorSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-tenant-cooperator", fasthttp.MethodGet, reqDto)
	var response dto.TenantCooperatorSingleRespDto
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
* @summary 获取一个协调管理员拥有的列表
* @description 根据用户池 ID 获取某个协调管理员的列表
* @param userId 用户 ID
* @returns TenantCooperatorSingleRespDto
 */
func (client *ManagementClient) GetTenantCooperatorMenu(reqDto *dto.GetTenantCooperatorMenuDto) *dto.TenantCooperatorSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-tenant-cooperator-menu", fasthttp.MethodGet, reqDto)
	var response dto.TenantCooperatorSingleRespDto
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
* @summary 创建协调管理员
* @description 创建一个协调管理员
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) CreateTenantCooperator(reqDto *dto.CreateTenantCooperatorDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/create-tenant-cooperator", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 更新协调管理员
* @description 根据用户 ID 更新一个协调管理员
* @param userId 用户 ID
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) UpdateTenantCooperator(reqDto *dto.UpdateTenantCooperatorDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/update-tenant-cooperator", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 删除协调管理员
* @description 使用用户 ID 删除一个协调管理员
* @param requestBody
* @param userId 用户 ID（请使用 body 中的 userId）
* @returns CommonResponseDto
 */
func (client *ManagementClient) DeleteTenantCooperator(reqDto *dto.DeleteTenantCooperatorBodyDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/delete-tenant-cooperator", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 获取租户详情
* @description 根据租户 Code 获取租户详情
* @param code 租户 Code
* @returns TenantSingleRespDto
 */
func (client *ManagementClient) GetTenantByCode(reqDto *dto.GetTenantByCodeDto) *dto.TenantSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-tenant-by-code", fasthttp.MethodGet, reqDto)
	var response dto.TenantSingleRespDto
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
* @summary 发送邀请租户用户邮件
* @description 向多个邮箱发送邀请成为租户用户的邮件
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) SendInviteTenantUserEmail(reqDto *dto.SendInviteTenantUserEmailDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/send-invite-tenant-user-email", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 添加租户成员
* @description 根据用户 ID 批量添加租户成员
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) AddTenantUsers(reqDto *dto.AddTenantUsersDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/add-tenant-users", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 批量移除租户成员
* @description 此接口用于根据用户 ID 或租户成员 ID 批量移除租户成员。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) RemoveTenantUsers(reqDto *dto.RemoveTenantUsersDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/remove-tenant-users", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 更新租户成员
* @description 此接口用于根据用户 ID 或租户成员 ID 更新租户成员。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) UpdateTenantUser(reqDto *dto.UpdateTenantUserDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/update-tenant-user", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 创建租户成员
* @description 创建租户成员，邮箱、手机号、用户名必须包含其中一个，邮箱、手机号、用户名、externalId 用户池内唯一，此接口将以管理员身份创建用户因此不需要进行手机号验证码检验等安全检测。
* @param requestBody
* @returns TenantUserDto
 */
func (client *ManagementClient) CreateTenantUser(reqDto *dto.CreateTenantUserReqDto) *dto.TenantUserDto {
	b, err := client.SendHttpRequest("/api/v3/create-tenant-user", fasthttp.MethodPost, reqDto)
	var response dto.TenantUserDto
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
	* @summary 获取/搜索租户成员列表
	* @description
 * 此接口用于获取用户列表，支持模糊搜索，以及通过用户基础字段、用户自定义字段、用户所在部门、用户历史登录应用等维度筛选用户。
 *
 * ### 模糊搜素示例
 *
 * 模糊搜索默认会从 `phone`, `email`, `name`, `username`, `nickname` 五个字段对用户进行模糊搜索，你也可以通过设置 `options.fuzzySearchOn`
 * 决定模糊匹配的字段范围：
 *
 * ```json
 * {
 * "keywords": "北京",
 * "options": {
 * "fuzzySearchOn": [
 * "address"
 * ]
 * }
 * }
 * ```
 *
 * ### 高级搜索示例
 *
 * 你可以通过 `advancedFilter` 进行高级搜索，高级搜索支持通过用户的基础信息、自定义数据、所在部门、用户来源、登录应用、外部身份源信息等维度对用户进行筛选。
 * **且这些筛选条件可以任意组合。**
 *
 * #### 筛选邮箱中包含 `@example.com` 的用户
 *
 * 用户邮箱（`email`）为字符串类型，可以进行模糊搜索：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "email",
 * "operator": "CONTAINS",
 * "value": "@example.com"
 * }
 * ]
 * }
 * ```
 *
 *
 * #### 根据用户登录次数筛选
 *
 * 筛选登录次数大于 10 的用户：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "loginsCount",
 * "operator": "GREATER",
 * "value": 10
 * }
 * ]
 * }
 * ```
 *
 * 筛选登录次数在 10 - 100 次的用户：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "loginsCount",
 * "operator": "BETWEEN",
 * "value": [10, 100]
 * }
 * ]
 * }
 * ```
 *
 * #### 根据用户上次登录时间进行筛选
 *
 * 筛选最近 7 天内登录过的用户：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "lastLoginTime",
 * "operator": "GREATER",
 * "value": new Date(Date.now() - 7 * 24 * 60 * 60 * 1000)
 * }
 * ]
 * }
 * ```
 *
 * 筛选在某一段时间内登录过的用户：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "lastLoginTime",
 * "operator": "BETWEEN",
 * "value": [
 * Date.now() - 14 * 24 * 60 * 60 * 1000,
 * Date.now() - 7 * 24 * 60 * 60 * 1000
 * ]
 * }
 * ]
 * }
 * ```
 *
 * #### 根据用户曾经登录过的应用筛选
 *
 * 筛选出曾经登录过应用 `appId1` 或者 `appId2` 的用户：
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "loggedInApps",
 * "operator": "IN",
 * "value": [
 * "appId1",
 * "appId2"
 * ]
 * }
 * ]
 * }
 * ```
 *
 * #### 根据用户所在部门进行筛选
 *
 * ```json
 * {
 * "advancedFilter": [
 * {
 * "field": "department",
 * "operator": "IN",
 * "value": [
 * {
 * "organizationCode": "steamory",
 * "departmentId": "root",
 * "departmentIdType": "department_id",
 * "includeChildrenDepartments": true
 * }
 * ]
 * }
 * ]
 * }
 * ```
 *
 *
	* @param requestBody
	* @returns TenantUserListPaginatedRespDto
*/
func (client *ManagementClient) ListTenantUsers(reqDto *dto.ListTenantUserDto) *dto.TenantUserListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-tenant-users", fasthttp.MethodPost, reqDto)
	var response dto.TenantUserListPaginatedRespDto
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
* @summary 获取单个租户成员
* @description 根据用户 ID 或租户成员 ID 获取租户成员信息
* @param tenantId 租户 ID
* @param linkUserId 关联的用户池级别的用户 ID
* @param memberId 租户成员 ID
* @returns TenantUserSingleRespDto
 */
func (client *ManagementClient) GetTenantUser(reqDto *dto.GetTenantUserDto) *dto.TenantUserSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-tenant-user", fasthttp.MethodGet, reqDto)
	var response dto.TenantUserSingleRespDto
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
* @summary 租户部门下添加成员
* @description 通过部门 ID、组织 code，添加部门下成员。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) AddTenantDepartmentMembers(reqDto *dto.AddTenantDepartmentMembersReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/add-tenant-department-members", fasthttp.MethodPost, reqDto)
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
* @summary 租户部门下删除成员
* @description 通过部门 ID、组织 code，删除部门下成员。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) RemoveTenantDepartmentMembers(reqDto *dto.RemoveTenantDepartmentMembersReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/remove-tenant-department-members", fasthttp.MethodPost, reqDto)
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
* @summary 创建权限空间
* @description 创建权限空间,可以设置权限空间名称、Code 和描述信息。
* @param requestBody
* @returns CreatePermissionNamespaceResponseDto
 */
func (client *ManagementClient) CreatePermissionNamespace(reqDto *dto.CreatePermissionNamespaceDto) *dto.CreatePermissionNamespaceResponseDto {
	b, err := client.SendHttpRequest("/api/v3/create-permission-namespace", fasthttp.MethodPost, reqDto)
	var response dto.CreatePermissionNamespaceResponseDto
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
* @summary 批量创建权限空间
* @description 批量创建权限空间，可以分别设置权限空间名称、Code 和描述信息。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) CreatePermissionNamespacesBatch(reqDto *dto.CreatePermissionNamespacesBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-permission-namespaces-batch", fasthttp.MethodPost, reqDto)
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
* @summary 获取权限空间详情
* @description 通过权限空间唯一标志符(Code)，获取权限空间详情。
* @param code 权限空间 Code
* @returns GetPermissionNamespaceResponseDto
 */
func (client *ManagementClient) GetPermissionNamespace(reqDto *dto.GetPermissionNamespaceDto) *dto.GetPermissionNamespaceResponseDto {
	b, err := client.SendHttpRequest("/api/v3/get-permission-namespace", fasthttp.MethodGet, reqDto)
	var response dto.GetPermissionNamespaceResponseDto
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
* @summary 批量获取权限空间详情列表
* @description 分别通过权限空间唯一标志符(Code)，获取权限空间详情。
* @param codes 权限空间 code 列表，批量可以使用逗号分隔
* @returns GetPermissionNamespaceListResponseDto
 */
func (client *ManagementClient) GetPermissionNamespacesBatch(reqDto *dto.GetPermissionNamespacesBatchDto) *dto.GetPermissionNamespaceListResponseDto {
	b, err := client.SendHttpRequest("/api/v3/get-permission-namespaces-batch", fasthttp.MethodGet, reqDto)
	var response dto.GetPermissionNamespaceListResponseDto
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
* @summary 分页获取权限空间列表
* @description 分页获取权限空间列表。
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param query 权限空间 name
* @returns PermissionNamespaceListPaginatedRespDto
 */
func (client *ManagementClient) ListPermissionNamespaces(reqDto *dto.ListPermissionNamespacesDto) *dto.PermissionNamespaceListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-permission-namespaces", fasthttp.MethodGet, reqDto)
	var response dto.PermissionNamespaceListPaginatedRespDto
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
* @summary 修改权限空间
* @description 修改权限空间，可以修改权限空间名称、权限空间描述信息以及权限空间新的唯一标志符(Code)。
* @param requestBody
* @returns UpdatePermissionNamespaceResponseDto
 */
func (client *ManagementClient) UpdatePermissionNamespace(reqDto *dto.UpdatePermissionNamespaceDto) *dto.UpdatePermissionNamespaceResponseDto {
	b, err := client.SendHttpRequest("/api/v3/update-permission-namespace", fasthttp.MethodPost, reqDto)
	var response dto.UpdatePermissionNamespaceResponseDto
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
* @summary 删除权限空间
* @description 通过权限空间唯一标志符(Code)，删除权限空间信息。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeletePermissionNamespace(reqDto *dto.DeletePermissionNamespaceDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-permission-namespace", fasthttp.MethodPost, reqDto)
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
* @summary 批量删除权限空间
* @description 分别通过权限空间唯一标志符(Code)，批量删除权限空间信息。
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeletePermissionNamespacesBatch(reqDto *dto.DeletePermissionNamespacesBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-permission-namespaces-batch", fasthttp.MethodPost, reqDto)
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
* @summary 校验权限空间 Code 或者名称是否可用
* @description 通过用户池 ID 和权限空间 Code,或者用户池 ID 和权限空间名称查询是否可用。
* @param requestBody
* @returns PermissionNamespaceCheckExistsRespDto
 */
func (client *ManagementClient) CheckPermissionNamespaceExists(reqDto *dto.CheckPermissionNamespaceExistsDto) *dto.PermissionNamespaceCheckExistsRespDto {
	b, err := client.SendHttpRequest("/api/v3/check-permission-namespace-exists", fasthttp.MethodPost, reqDto)
	var response dto.PermissionNamespaceCheckExistsRespDto
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
* @summary 分页查询权限空间下所有的角色列表
* @description 分页查询权限空间下所有的角色列表，分页获取权限空间下所有的角色列表。
* @param code 权限分组唯一标志符 Code
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param query 角色 Code 或者名称
* @returns PermissionNamespaceRolesListPaginatedRespDto
 */
func (client *ManagementClient) ListPermissionNamespaceRoles(reqDto *dto.ListPermissionNamespaceRolesDto) *dto.PermissionNamespaceRolesListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-permission-namespace-roles", fasthttp.MethodGet, reqDto)
	var response dto.PermissionNamespaceRolesListPaginatedRespDto
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
	* @summary 创建数据资源（推荐、重点）
	* @description
 * ## 描述
 * 该接口用于创建数据资源，当你存在需要被设置权限的数据，根据它们的数据类型，创建对应类型的数据资源，目前我们支持：字符串、数组、树三种类型。
 * ## 注意
 * 请求体中的 `struct` 字段需要根据不同的资源类型传入不同的数据结构，具体请参考下面的示例
 * ## 请求示例
 * ### 创建字符串类型数据资源示例
 * 当你的数据仅用一个字符串就可以表示时，可以使用此类型，例如：一个 API、一个用户 ID 等。
 * 以下是创建一个表示 '/resource/create' API 的数据资源示例：
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "resourceName": "createResource API",
 * "description": "这是 createResource API",
 * "resourceCode": "createResourceAPI",
 * "type": "STRING",
 * "struct": "/resource/create",
 * "actions": ["access"]
 * }
 * ```
 *
 * ### 创建数组类型数据资源示例
 * 当你的数据是一组同类型的数据时，可以使用此类型，例如：一组文档链接、一组门禁卡号等。
 * 以下是创建一个表示一组门禁卡号的数据资源示例：
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "resourceName": "一组门禁卡号",
 * "description": "这是一组门禁卡号",
 * "resourceCode": "accessCardNumber",
 * "type": "ARRAY",
 * "struct": ["accessCardNumber1", "accessCardNumber2", "accessCardNumber3"],
 * "actions": ["get", "update"]
 * }
 * ```
 *
 * ### 创建树类型数据资源示例
 * 当你的数据是具备层级关系时，可以使用此类型，例如：组织架构、文件夹结构等。
 * 以下是创建一个表示公司组织架构的数据资源示例：
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "resourceName": "Authing",
 * "description": "这是 Authing 的组织架构",
 * "resourceCode": "authing",
 * "type": "TREE",
 * "struct": [
 * {
 * "name": "产品",
 * "code": "product",
 * "value": "product",
 * "children": [
 * {
 * "name": "产品经理",
 * "code": "productManager",
 * "value": "pm"
 * },
 * {
 * "name": "设计",
 * "code": "design",
 * "value": "ui"
 * }
 * ]
 * },
 * {
 * "name": "研发",
 * "code": "researchAndDevelopment",
 * "value": "rd"
 * }
 * ],
 * "actions": ["get", "update", "delete"]
 * }
 * ```
 *
	* @param requestBody
	* @returns CreateDataResourceResponseDto
*/
func (client *ManagementClient) CreateDataResource(reqDto *dto.CreateDataResourceDto) *dto.CreateDataResourceResponseDto {
	b, err := client.SendHttpRequest("/api/v3/create-data-resource", fasthttp.MethodPost, reqDto)
	var response dto.CreateDataResourceResponseDto
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
* @summary 创建字符串数据资源
* @description 当你仅需要创建字符串类型数据资源时，可以使用此 API，我们固定了数据资源类型，你无需在传入 `type` 字符段，注意：`struct` 字段只能够传入字符串类型数据。
* @param requestBody
* @returns CreateStringDataResourceResponseDto
 */
func (client *ManagementClient) CreateDataResourceByString(reqDto *dto.CreateStringDataResourceDto) *dto.CreateStringDataResourceResponseDto {
	b, err := client.SendHttpRequest("/api/v3/create-string-data-resource", fasthttp.MethodPost, reqDto)
	var response dto.CreateStringDataResourceResponseDto
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
* @summary 创建数组数据资源
* @description 当你仅需要创建数组类型数据资源时，可以使用此 API，我们固定了数据资源类型，你无需在传入 `type` 字符段，注意：`struct` 字段只能够传入数组类型数据。
* @param requestBody
* @returns CreateArrayDataResourceResponseDto
 */
func (client *ManagementClient) CreateDataResourceByArray(reqDto *dto.CreateArrayDataResourceDto) *dto.CreateArrayDataResourceResponseDto {
	b, err := client.SendHttpRequest("/api/v3/create-array-data-resource", fasthttp.MethodPost, reqDto)
	var response dto.CreateArrayDataResourceResponseDto
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
* @summary 创建树数据资源
* @description 当你仅需要创建树类型数据资源时，可以使用此 API，我们固定了数据资源类型，你无需在传入 `type` 字符段，注意：`struct` 要按照树类型数据资源结构进行传入，请参考示例。
* @param requestBody
* @returns CreateTreeDataResourceResponseDto
 */
func (client *ManagementClient) CreateDataResourceByTree(reqDto *dto.CreateTreeDataResourceDto) *dto.CreateTreeDataResourceResponseDto {
	b, err := client.SendHttpRequest("/api/v3/create-tree-data-resource", fasthttp.MethodPost, reqDto)
	var response dto.CreateTreeDataResourceResponseDto
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
* @summary 获取数据资源列表
* @description 获取数据资源列表,可通过数据资源名称、数据资源 Code 和数据资源所属权限空间 Code 列表进行指定筛选。
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param query 关键字搜索，可以是数据资源名称或者数据资源 Code
* @param namespaceCodes 权限数据所属权限空间 Code 列表
* @returns ListDataResourcesPaginatedRespDto
 */
func (client *ManagementClient) ListDataResources(reqDto *dto.ListDataResourcesDto) *dto.ListDataResourcesPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-data-resources", fasthttp.MethodGet, reqDto)
	var response dto.ListDataResourcesPaginatedRespDto
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
* @summary 获取数据资源信息
* @description 获取数据资源,通过数据资源 ID 查询对应的数据资源信息,包含数据资源名称、数据资源 Code、数据资源类型（TREE、STRING、ARRAY）、数据资源所属权限空间 ID、数据资源所属权限空间 Code 以及数据资源操作列表等基本信息。
* @param namespaceCode 数据资源所属的权限空间 Code
* @param resourceCode 数据资源 Code, 权限空间内唯一
* @returns GetDataResourceResponseDto
 */
func (client *ManagementClient) GetDataResource(reqDto *dto.GetDataResourceDto) *dto.GetDataResourceResponseDto {
	b, err := client.SendHttpRequest("/api/v3/get-data-resource", fasthttp.MethodGet, reqDto)
	var response dto.GetDataResourceResponseDto
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
* @summary 修改数据资源
* @description 修改数据资源,根据权限空间 Code 和数据资源 Code 查询原始信息,只允许修改数据资源名称、描述和数据资源节点。
* @param requestBody
* @returns UpdateDataResourceResponseDto
 */
func (client *ManagementClient) UpdateDataResource(reqDto *dto.UpdateDataResourceDto) *dto.UpdateDataResourceResponseDto {
	b, err := client.SendHttpRequest("/api/v3/update-data-resource", fasthttp.MethodPost, reqDto)
	var response dto.UpdateDataResourceResponseDto
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
* @summary 删除数据资源
* @description 删除数据资源,根据数据资源 ID 删除对应的数据资源信息。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) DeleteDataResource(reqDto *dto.DeleteDataResourceDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/delete-data-resource", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
	* @summary 检查数据资源名称或者 Code 是否可用
	* @description 检查数据资源名称或者 Code 在权限空间内是否有效,通过数据资源名称或者数据资源 Code 以及所属权限空间 Code,判断在指定的权限空间内是否可用。
 *
 * ### 数据资源 Code 有效示例
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "resourceCode": "test"
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 0,
 * "data": {
 * "isValid": "true"
 * }
 * }
 * ```
 *
 * ### 数据资源名称有效示例
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "resourceName": "test"
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 0,
 * "data": {
 * "isValid": "true"
 * }
 * }
 * ```
 *
 * ### 数据资源 Code 无效示例
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "resourceCode": "test"
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 0,
 * "requestId": "934108e5-9fbf-4d24-8da1-c330328abd6c",
 * "data": {
 * "isValid": "false",
 * "message": "data resource code already exist"
 * }
 * }
 * ```
 *
	* @param namespaceCode 数据资源所属的权限空间 Code
	* @param resourceName 数据资源名称, 权限空间内唯一
	* @param resourceCode 数据资源 Code, 权限空间内唯一
	* @returns CheckParamsDataResourceResponseDto
*/
func (client *ManagementClient) CheckDataResourceExists(reqDto *dto.CheckDataResourceExistsDto) *dto.CheckParamsDataResourceResponseDto {
	b, err := client.SendHttpRequest("/api/v3/check-data-resource-exists", fasthttp.MethodGet, reqDto)
	var response dto.CheckParamsDataResourceResponseDto
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
	* @summary 创建数据策略（重点）
	* @description
 * ## 描述
 * 该接口用于创建数据策略，通过数据策略你可以将一组数据资源及其指定的操作进行绑定到一起，共同授权给主体。
 * ## 注意
 * 为了方便使用，`permissions` 字段我们基于路径的方式提供了快捷的写法，如：
 * - 数组、字符串资源：权限空间 code/数据资源 code/数据资源某一 action（如果表示所有操作，则使用`*`替代action）
 * - 树类型资源：权限空间 code/数据资源 code/node code 1/node code 1_1/.../数据资源某一 action
 *
 * ## 请求示例
 * 假设我们要对一名研发人员进行授权，首先创建 3 个数据资源如下：
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "resourceName": "服务器",
 * "resourceCode": "server_2023",
 * "type": "STRING",
 * "struct": "server_2023",
 * "actions": ["read", "write"]
 * }
 * ```
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "resourceName": "研发知识库",
 * "description": "",
 * "resourceCode": "rd_document",
 * "type": "STRING",
 * "struct": "https://www.authing.com/rd_document",
 * "actions": ["read", "write", "share"]
 * }
 * ```
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "resourceName": "研发内部平台菜单",
 * "description": "这是研发使用的内部平台菜单",
 * "resourceCode": "rd_internal_platform",
 * "type": "TREE",
 * "struct": [
 * {
 * "name": "部署",
 * "code": "deploy",
 * "children": [
 * {
 * "name": "生产环境",
 * "code": "prod"
 * },
 * {
 * "name": "测试环境",
 * "code": "test"
 * }
 * ]
 * },
 * {
 * "name": "数据库",
 * "code": "db"
 * "children": [
 * {
 * "name": "查询",
 * "code": "query"
 * },
 * {
 * "name": "导出",
 * "code": "export"
 * }
 * ]
 * }
 * ],
 * "actions": ["access", "execute"]
 * }
 * ```
 * 我们分配一台服务器：server_2023 给他使用，他可以在上面进行任意操作，同时他可以阅读、编辑研发知识库，最后他可以在研发内部平台中进行部署测试环境，但是不能够导出数据库数据。
 * ```json
 * {
 * "policyName": "研发人员策略",
 * "description": "这是一个示例数据策略",
 * "statementList": [
 * {
 * "effect": "ALLOW",
 * "permissions": [
 * "examplePermissionNamespaceCode/server_2023*",
 * "examplePermissionNamespaceCode/rd_document/read",
 * "examplePermissionNamespaceCode/rd_document/write",
 * "examplePermissionNamespaceCode/rd_internal_platform/deploy/test/execute"
 * ]
 * },
 * {
 * "effect": "DENY",
 * "permissions": [
 * "examplePermissionNamespaceCode/rd_internal_platform/db/export/execute"
 * ]
 * }
 * ]
 * }
 * ```
 *
	* @param requestBody
	* @returns CreateDataPolicyResponseDto
*/
func (client *ManagementClient) CreateDataPolicy(reqDto *dto.CreateDataPolicyDto) *dto.CreateDataPolicyResponseDto {
	b, err := client.SendHttpRequest("/api/v3/create-data-policy", fasthttp.MethodPost, reqDto)
	var response dto.CreateDataPolicyResponseDto
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
* @summary 获取数据策略列表
* @description 分页查询数据策略列表，也可通过关键字搜索数据策略名称或者数据策略 Code 进行模糊查找。
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param query 数据策略名称关键字搜索
* @returns ListDataPoliciesPaginatedRespDto
 */
func (client *ManagementClient) ListDataPolices(reqDto *dto.ListDataPoliciesDto) *dto.ListDataPoliciesPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-data-policies", fasthttp.MethodGet, reqDto)
	var response dto.ListDataPoliciesPaginatedRespDto
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
* @summary 获取数据策略简略信息列表
* @description 分页获取数据策略简略信息列表，通过关键字搜索数据策略名称或者数据策略 Code 进行模糊查找出数据策略 ID、数据策略名称和数据策略描述信息。
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param query 数据策略名称关键字搜索
* @returns ListSimpleDataPoliciesPaginatedRespDto
 */
func (client *ManagementClient) ListSimpleDataPolices(reqDto *dto.ListSimpleDataPoliciesDto) *dto.ListSimpleDataPoliciesPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-simple-data-policies", fasthttp.MethodGet, reqDto)
	var response dto.ListSimpleDataPoliciesPaginatedRespDto
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
* @summary 获取数据策略信息
* @description 获取数据策略信息，通过数据策略 ID 获取对应数据策略信息,包含数据策略 ID、数据策略名称、数据策略描述、数据策略下所有的数据权限列表等信息。
* @param policyId 数据策略 ID
* @returns GetDataPolicyResponseDto
 */
func (client *ManagementClient) GetDataPolicy(reqDto *dto.GetDataPolicyDto) *dto.GetDataPolicyResponseDto {
	b, err := client.SendHttpRequest("/api/v3/get-data-policy", fasthttp.MethodGet, reqDto)
	var response dto.GetDataPolicyResponseDto
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
* @summary 修改数据策略
* @description 修改数据策略，通过数据策略名称、数据策略描述和相关的数据资源等字段修改数据策略信息。
* @param requestBody
* @returns UpdateDataPolicyResponseDto
 */
func (client *ManagementClient) UpdateDataPolicy(reqDto *dto.UpdateDataPolicyDto) *dto.UpdateDataPolicyResponseDto {
	b, err := client.SendHttpRequest("/api/v3/update-data-policy", fasthttp.MethodPost, reqDto)
	var response dto.UpdateDataPolicyResponseDto
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
* @summary 删除数据策略
* @description 删除数据策略，通过数据策略 ID 删除对应的策略,同时也删除数据策略和对应的数据资源等关系数据。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) DeleteDataPolicy(reqDto *dto.DeleteDataPolicyDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/delete-data-policy", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 校验数据策略名称是否有效
* @description 通过数据策略名称查询用户池内是否有效。
* @param policyName 数据策略名称，用户池唯一
* @returns CheckParamsDataPolicyResponseDto
 */
func (client *ManagementClient) CheckDataPolicyExists(reqDto *dto.CheckDataPolicyExistsDto) *dto.CheckParamsDataPolicyResponseDto {
	b, err := client.SendHttpRequest("/api/v3/check-data-policy-exists", fasthttp.MethodGet, reqDto)
	var response dto.CheckParamsDataPolicyResponseDto
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
* @summary 获取数据策略下所有的授权主体的信息
* @description 获取数据策略下所有的授权主体的信息，通过授权主体类型、数据策略 ID 和数据资源 ID 查找授权主体列表。
* @param policyId 数据策略 ID
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param query 主体名称
* @param targetType 主体类型,包括 USER、GROUP、ROLE、ORG 四种类型
* @returns ListDataPolicySubjectPaginatedRespDto
 */
func (client *ManagementClient) ListDataPolicyTargets(reqDto *dto.ListDataPolicyTargetsDto) *dto.ListDataPolicySubjectPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-data-policy-targets", fasthttp.MethodGet, reqDto)
	var response dto.ListDataPolicySubjectPaginatedRespDto
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
* @summary 授权数据策略
* @description 数据策略创建主体授权，通过授权主体和数据策略进行相互授权。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) AuthorizeDataPolicies(reqDto *dto.CreateAuthorizeDataPolicyDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/authorize-data-policies", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 撤销数据策略
* @description 删除数据策略相关的主体授权，通过授权主体 ID、授权主体类型和数据策略 ID 进行删除。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) RevokeDataPolicy(reqDto *dto.DeleteAuthorizeDataPolicyDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/revoke-data-policy", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
	* @summary 获取用户权限列表
	* @description
 * ## 描述
 * 该接口用于查询某些用户在某些权限空间的数据资源的权限数据。
 * 我们的鉴权接口针对不同的鉴权场景有多种，区别在于在所处的场景下能够传递的参数列表以及不同形式的出参，**当你需要查询某些用户的所有权限时**可以使用此接口，
 * ## 注意
 * 接口提供了两个数组类型的入参`userIds`和`namespaceCodes`来支持批量查询（注意：namespaceCodes 是可选的）。
 * ## 场景举例
 * 假如你的业务场景是用户登录后能看到他所有可以访问或者进行其他操作的文档、人员信息、设备信息等资源，那么你可以在用户登录后调用此接口查询用户的所有权限。
 * ## 请求示例
 * ### 查询单个用户权限列表示例
 * 注意：在此接口的返回参数中，树类型的数据资源权限返回是按照**路径**的方式返回的
 * - 入参
 *
 * ```json
 * {
 * "userIds": [
 * "6301ceaxxxxxxxxxxx27478"
 * ]
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data": {
 * "userPermissionList": [
 * {
 * "userId": "6301ceaxxxxxxxxxxx27478",
 * "namespaceCode": "examplePermissionNamespace",
 * "resourceList": [
 * {
 * "resourceCode": "strCode",
 * "resourceType": "STRING",
 * "strAuthorize": {
 * "value": "示例字符串资源",
 * "actions": [
 * "read",
 * "post",
 * "get",
 * "write"
 * ]
 * }
 * },
 * {
 * "resourceCode": "arrayCode",
 * "resourceType": "ARRAY",
 * "arrAuthorize": {
 * "values": [
 * "示例数据资源1",
 * "示例数据资源2"
 * ],
 * "actions": [
 * "read",
 * "post",
 * "get",
 * "write"
 * ]
 * }
 * },
 * {
 * "resourceCode": "treeCode",
 * "resourceType": "TREE",
 * "treeAuthorize": {
 * "authList": [
 * {
 * "nodePath": "/treeChildrenCode/treeChildrenCode1",
 * "nodeActions": [
 * "read",
 * "get"
 * ],
 * "nodeName": "treeChildrenName1",
 * "nodeValue": "treeChildrenValue1"
 * },
 * {
 * "nodePath": "/treeChildrenCode/treeChildrenCode2",
 * "nodeActions": [
 * "read",
 * "get"
 * ],
 * "nodeName": "treeChildrenName2",
 * "nodeValue": "treeChildrenValue2"
 * },
 * {
 * "nodePath": "/treeChildrenCode/treeChildrenCode3",
 * "nodeActions": [
 * "read"
 * ],
 * "nodeName": "treeChildrenName3",
 * "nodeValue": "treeChildrenValue3"
 * }
 * ]
 * }
 * }
 * ]
 * }
 * ]
 * }
 * }
 * ```
 *
 * ### 查询多个用户权限列表示例
 *
 * - 入参
 *
 * ```json
 * {
 * "userIds": [
 * "6301ceaxxxxxxxxxxx27478",
 * "6121ceaxxxxxxxxxxx27312"
 * ]
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data": {
 * "userPermissionList": [
 * {
 * "userId": "6301ceaxxxxxxxxxxx27478",
 * "namespaceCode": "examplePermissionNamespace1",
 * "resourceList": [
 * {
 * "resourceCode": "strCode",
 * "resourceType": "STRING",
 * "strAuthorize": {
 * "value": "示例字符串资源",
 * "actions": [
 * "read",
 * "post",
 * "get",
 * "write"
 * ]
 * }
 * }
 * ]
 * },
 * {
 * "userId": "6121ceaxxxxxxxxxxx27312",
 * "namespaceCode": "examplePermissionNamespace2",
 * "resourceList": [
 * {
 * "resourceCode": "arrayCode",
 * "resourceType": "ARRAY",
 * "arrAuthorize": {
 * "values": [
 * "示例数组资源1",
 * "示例数组资源2"
 * ],
 * "actions": [
 * "read",
 * "post",
 * "get",
 * "write"
 * ]
 * }
 * }
 * ]
 * }
 * ]
 * }
 * }
 * ```
 *
 * ### 查询多个用户在多个权限空间下权限列表示例
 *
 * - 入参
 *
 * ```json
 * {
 * "userIds": [
 * "6301ceaxxxxxxxxxxx27478",
 * "6121ceaxxxxxxxxxxx27312"
 * ],
 * "namespaceCodes": [
 * "examplePermissionNamespace1",
 * "examplePermissionNamespace2"
 * ]
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data": {
 * "userPermissionList": [
 * {
 * "userId": "6301ceaxxxxxxxxxxx27478",
 * "namespaceCode": "examplePermissionNamespace1",
 * "resourceList": [
 * {
 * "resourceCode": "strCode1",
 * "resourceType": "STRING",
 * "strAuthorize": {
 * "value": "示例字符串资源",
 * "actions": [
 * "read",
 * "post",
 * "get",
 * "write"
 * ]
 * }
 * }
 * ]
 * },
 * {
 * "userId": "6121ceaxxxxxxxxxxx27312",
 * "namespaceCode": "examplePermissionNamespace2",
 * "resourceList": [
 * {
 * "resourceCode": "arrayCode",
 * "resourceType": "ARRAY",
 * "arrAuthorize": {
 * "values": [
 * "示例数组资源1",
 * "示例数组资源2"
 * ],
 * "actions": [
 * "read",
 * "post",
 * "get",
 * "write"
 * ]
 * }
 * }
 * ]
 * }
 * ]
 * }
 * }
 * ```
 *
	* @param requestBody
	* @returns GetUserPermissionListRespDto
*/
func (client *ManagementClient) GetUserPermissionList(reqDto *dto.GetUserPermissionListDto) *dto.GetUserPermissionListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-permission-list", fasthttp.MethodPost, reqDto)
	var response dto.GetUserPermissionListRespDto
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
	* @summary 判断用户权限（重点）
	* @description
 * ## 描述
 * 当你需要判断用户是否拥有某些资源的指定权限时可以使用此接口
 * ## 注意
 * - 该接口通过传递资源 code 定位对应的数据资源（如果是树类型，则需要传递节点的完整 code 路径）。
 * - 如果你在配置数据策略时配置了**环境属性**的条件判断，那么需要设置参数`judgeConditionEnabled`为`true`（默认为 false），并且通过参数`authEnvParams`传递请求时所处的环境信息（如 IP、设备类型、系统类型等），否则条件判断会不生效，导致数据策略不生效。
 *
 * ## 场景举例
 * 用户在删除某条数据时，需要判断他是否拥有此资源的删除权限，则可以使用此接口。
 *
 * ## 请求示例
 * ### 判断用户对字符串和数组资源权限示例（无条件判断）
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "userId": "63721xxxxxxxxxxxxdde14a3",
 * "action": "get",
 * "resources":["strResourceCode1", "arrayResourceCode1"]
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data": {
 * "checkResultList": [
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "resource": "strResourceCode1",
 * "action": "get",
 * "enabled": true
 * },
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "resource": "arrayResourceCode1",
 * "action": "get",
 * "enabled": true
 * }
 * ]
 * }
 * }
 * ```
 *
 * ### 判断用户对字符串和数组资源权限示例（开启条件判断）
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "userId": "63721xxxxxxxxxxxxdde14a3",
 * "action": "get",
 * "resources": ["strResourceCode1", "arrayResourceCode1"],
 * "judgeConditionEnabled": true,
 * "authEnvParams":{
 * "ip":"110.96.0.0",
 * "city":"北京",
 * "province":"北京",
 * "country":"中国",
 * "deviceType":"PC",
 * "systemType":"ios",
 * "browserType":"IE",
 * "requestDate":"2022-12-26 17:40:00"
 * }
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data": {
 * "checkResultList": [
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "resource": "strResourceCode1",
 * "action": "get",
 * "enabled": false
 * },
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "resource": "arrayResourceCode1",
 * "action": "get",
 * "enabled": false
 * }
 * ]
 * }
 * }
 * ```
 *
 * ### 判断用户对树资源权限示例
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "userId": "63721xxxxxxxxxxxxdde14a3",
 * "action": "get",
 * "resources":["treeResourceCode1/StructCode1/resourceStructChildrenCode1", "treeResourceCode2/StructCode1/resourceStructChildrenCode1"]
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data":{
 * "checkResultList": [{
 * "namespaceCode": "examplePermissionNamespace",
 * "action": "get",
 * "resource": "treeResourceCode1/StructCode1/resourceStructChildrenCode1",
 * "enabled": true
 * },{
 * "namespaceCode": "examplePermissionNamespace",
 * "action": "get",
 * "resource": "treeResourceCode2/StructCode1/resourceStructChildrenCode1",
 * "enabled": true
 * }]
 * }
 * }
 * ```
 *
	* @param requestBody
	* @returns CheckPermissionRespDto
*/
func (client *ManagementClient) CheckPermission(reqDto *dto.CheckPermissionDto) *dto.CheckPermissionRespDto {
	b, err := client.SendHttpRequest("/api/v3/check-permission", fasthttp.MethodPost, reqDto)
	var response dto.CheckPermissionRespDto
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
	* @summary 判断外部用户权限
	* @description
 * ## 描述
 * 当你的用户是外部用户，需要判断其是否拥有某资源的某一权限时，可以使用此接口，通过`externalId`来传递用户的 ID
 *
	* @param requestBody
	* @returns CheckExternalUserPermissionRespDto
*/
func (client *ManagementClient) CheckExternalUserPermission(reqDto *dto.CheckExternalUserPermissionDto) *dto.CheckExternalUserPermissionRespDto {
	b, err := client.SendHttpRequest("/api/v3/check-external-user-permission", fasthttp.MethodPost, reqDto)
	var response dto.CheckExternalUserPermissionRespDto
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
	* @summary 获取用户指定资源权限列表（推荐）
	* @description
 * ## 描述
 * 当你需要查询某一用户拥有指定的资源列表的权限时，可以使用此接口。
 * ## 注意
 * 该接口需要你传递指定的资源 code（如果是树类型资源则需要传递节点的完整 code 路径），此接口性能更佳，推荐使用。
 * ## 请求示例
 * ### 获取用户字符串和数组资源权限示例
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "userId": "63721xxxxxxxxxxxxdde14a3",
 * "resources":["strResourceCode1", "arrayResourceCode1"]
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 *
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data":{
 * "permissionList": [{
 * "namespaceCode": "examplePermissionNamespace",
 * "actions": ["read","get"],
 * "resource": "strResourceCode1"
 * },{
 * "namespaceCode": "examplePermissionNamespace",
 * "actions": ["read","update","delete"],
 * "resource": "arrayResourceCode1"
 * }]
 * }
 * }
 * ```
 *
 * ### 获取用户树资源权限示例
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "userId": "63721xxxxxxxxxxxxdde14a3",
 * "resources":["treeResourceCode1/StructCode1/resourceStructChildrenCode1", "treeResourceCode2/StructCode1/resourceStructChildrenCode1"]
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data":{
 * "permissionList": [{
 * "namespaceCode": "examplePermissionNamespace",
 * "actions": ["read", "update", "delete"],
 * "resource": "treeResourceCode1/StructCode1/resourceStructChildrenCode1"
 * },{
 * "namespaceCode": "examplePermissionNamespace",
 * "actions": ["read", "get", "delete"],
 * "resource": "treeResourceCode2/StructCode1/resourceStructChildrenCode1"
 * }]
 * }
 * }
 * ```
 *
	* @param requestBody
	* @returns GetUserResourcePermissionListRespDto
*/
func (client *ManagementClient) GetUserResourcePermissionList(reqDto *dto.GetUserResourcePermissionListDto) *dto.GetUserResourcePermissionListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-resource-permission-list", fasthttp.MethodPost, reqDto)
	var response dto.GetUserResourcePermissionListRespDto
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
	* @summary 获取资源被授权的用户列表
	* @description
 * ## 描述
 * 当你需要获取拥有指定资源的权限的用户时，可以使用此接口。
 * ## 场景举例
 * 假如你的业务场景是：想看看当前文档能够编辑的用户列表，那么你可以使用此接口。
 * ## 请求示例
 * ### 获取字符串和数组资源被授权的用户列表示例
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "actions": ["get", "update", "read"],
 * "resources":["strResourceCode1", "arrayResourceCode1"]
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data":{
 * "authUserList": [{
 * "resource": "strResourceCode1",
 * "actionAuthList": [{
 * "userIds": ["63721xxxxxxxxxxxxdde14a3"],
 * "action": "get"
 * },{
 * "userIds": ["63721xxxxxxxxxxxxdde14a3"],
 * "action": "update"
 * },{
 * "userIds": ["63721xxxxxxxxxxxxdde14a3"],
 * "action": "read"
 * }]
 * },{
 * "resource": "arrayResourceCode1",
 * "actionAuthList": [{
 * "userIds": ["63721xxxxxxxxxxxxdde14a3"],
 * "action": "get"
 * },{
 * "userIds": ["63721xxxxxxxxxxxxdde14a3"],
 * "action": "update"
 * },{
 * "userIds": ["63721xxxxxxxxxxxxdde14a3"],
 * "action": "read"
 * }]
 * }]
 * }
 * }
 * ```
 *
 * ### 获取树资源被授权的用户列表示例
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "actions": ["get", "update", "delete"],
 * "resources":["treeResourceCode1/StructCode1/resourceStructChildrenCode1", "treeResourceCode2/StructCode1/resourceStructChildrenCode1"]
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data":{
 * "authUserList": [{
 * "resource": "treeResourceCode1/StructCode1/resourceStructChildrenCode1",
 * "actionAuthList": [{
 * "userIds": ["63721xxxxxxxxxxxxdde14a3"],
 * "action": "get"
 * },{
 * "userIds": ["63721xxxxxxxxxxxxdde14a3"],
 * "action": "update"
 * },{
 * "userIds": ["63721xxxxxxxxxxxxdde14a3"],
 * "action": "delete"
 * }]
 * },{
 * "resource": "treeResourceCode2/StructCode1/resourceStructChildrenCode1",
 * "actionAuthList": [{
 * "userIds": ["63721xxxxxxxxxxxxdde14a3"],
 * "action": "get"
 * },{
 * "userIds": ["63721xxxxxxxxxxxxdde14a3"],
 * "action": "update"
 * },{
 * "userIds": ["63721xxxxxxxxxxxxdde14a3"],
 * "action": "delete"
 * }]
 * }]
 * }
 * }
 * ```
 *
	* @param requestBody
	* @returns ListResourceTargetsRespDto
*/
func (client *ManagementClient) ListResourceTargets(reqDto *dto.ListResourceTargetsDto) *dto.ListResourceTargetsRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-resource-targets", fasthttp.MethodPost, reqDto)
	var response dto.ListResourceTargetsRespDto
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
	* @summary 获取用户授权资源的结构列表
	* @description
 * ## 描述
 * 当你需要获取用户拥有某一资源的权限并且需要这个资源的结构信息（树类型资源返回树结构，数组类型资源返回数组、字符串类型返回字符串）则可以使用此接口。
 * ## 注意
 * 由于其他接口对树类型资源的返回参数格式是按照路径的方式返回的，所以我们提供此此接口，按照完整的树结构形式返回。
 * ## 请求示例
 * ### 获取用户授权字符串数据资源示例
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "userId": "63721xxxxxxxxxxxxdde14a3",
 * "resourceCode": "exampleStrResourceCode"
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data":{
 * "namespaceCode": "exampleNamespaceCode",
 * "resourceCode": "exampleStrResourceCode",
 * "resourceType": "STRING",
 * "strResourceAuthAction":{
 * "value": "strTestValue",
 * "actions": ["get","delete"]
 * }
 * }
 * }
 * ```
 *
 *
 * ### 获取用户授权数据数组资源示例
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "userId": "63721xxxxxxxxxxxxdde14a3",
 * "resourceCode": "exampleArrResourceCode"
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data":{
 * "namespaceCode": "exampleNamespaceCode",
 * "resourceCode": "exampleArrResourceCode",
 * "resourceType": "ARRAY",
 * "arrResourceAuthAction":{
 * "values": ["arrTestValue1","arrTestValue2","arrTestValue3"],
 * "actions": ["get","delete"]
 * }
 * }
 * }
 * ```
 *
 *
 * ### 获取用户授权树数据资源示例
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "userId": "63721xxxxxxxxxxxxdde14a3",
 * "resourceCode": "exampleArrResourceCode"
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data":{
 * "namespaceCode": "exampleNamespaceCode",
 * "resourceCode": "exampleArrResourceCode",
 * "resourceType": "TREE",
 * "treeResourceAuthAction":{
 * "nodeAuthActionList":[{
 * "code": "tree11",
 * "name": "tree11",
 * "value": "test11Value",
 * "actions": ["get","delete"],
 * "children": [{
 * "code": "tree111",
 * "name": "tree111",
 * "value": "test111Value",
 * "actions": ["update","read"],
 * }]
 * },{
 * "code": "tree22",
 * "name": "tree22",
 * "value": "test22Value",
 * "actions": ["get","delete"]
 * }]
 * }
 * }
 * }
 * ```
 *
	* @param requestBody
	* @returns GetUserResourceStructRespDto
*/
func (client *ManagementClient) GetUserResourceStruct(reqDto *dto.GetUserResourceStructDto) *dto.GetUserResourceStructRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-resource-struct", fasthttp.MethodPost, reqDto)
	var response dto.GetUserResourceStructRespDto
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
	* @summary 获取外部用户授权资源的结构列表
	* @description
 * ## 描述
 * 当你需要获取外部用户（通过参数`externalId`参数传递外部用户 ID）拥有某一资源的权限并且需要这个资源的结构信息（树类型资源返回树结构，数组类型资源返回数组、字符串类型返回字符串）则可以使用此接口。
 * ## 请求示例
 * ### 获取用户授权字符串数据资源示例
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "externalId": "63721xxxxxxxxxxxxdde14a3",
 * "resourceCode": "exampleStrResourceCode"
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data":{
 * "namespaceCode": "exampleNamespaceCode",
 * "resourceCode": "exampleStrResourceCode",
 * "resourceType": "STRING",
 * "strResourceAuthAction":{
 * "value": "strTestValue",
 * "actions": ["get","delete"]
 * }
 * }
 * }
 * ```
 *
 *
 * ### 获取用户授权数据数组资源示例
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "externalId": "63721xxxxxxxxxxxxdde14a3",
 * "resourceCode": "exampleArrResourceCode"
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data":{
 * "namespaceCode": "exampleNamespaceCode",
 * "resourceCode": "exampleArrResourceCode",
 * "resourceType": "ARRAY",
 * "arrResourceAuthAction":{
 * "values": ["arrTestValue1","arrTestValue2","arrTestValue3"],
 * "actions": ["get","delete"]
 * }
 * }
 * }
 * ```
 *
 *
 * ### 获取用户授权树数据资源示例
 *
 * - 入参
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "externalId": "63721xxxxxxxxxxxxdde14a3",
 * "resourceCode": "exampleArrResourceCode"
 * }
 * ```
 *
 * - 出参
 *
 * ```json
 * {
 * "statusCode": 200,
 * "message": "操作成功",
 * "apiCode": 20001,
 * "data":{
 * "namespaceCode": "exampleNamespaceCode",
 * "resourceCode": "exampleArrResourceCode",
 * "resourceType": "TREE",
 * "treeResourceAuthAction":{
 * "nodeAuthActionList":[{
 * "code": "tree11",
 * "name": "tree11",
 * "value": "test11Value",
 * "actions": ["get","delete"],
 * "children": [{
 * "code": "tree111",
 * "name": "tree111",
 * "value": "test111Value",
 * "actions": ["update","read"],
 * }]
 * },{
 * "code": "tree22",
 * "name": "tree22",
 * "value": "test22Value",
 * "actions": ["get","delete"]
 * }]
 * }
 * }
 * }
 * ```
 *
	* @param requestBody
	* @returns GetExternalUserResourceStructRespDto
*/
func (client *ManagementClient) GetExternalUserResourceStruct(reqDto *dto.GetExternalUserResourceStructDto) *dto.GetExternalUserResourceStructRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-external-user-resource-struct", fasthttp.MethodPost, reqDto)
	var response dto.GetExternalUserResourceStructRespDto
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
	* @summary 判断用户在树资源同层级下的权限（推荐）
	* @description
 * ## 描述
 * 此接口用于判断用户是否拥有某一树资源的**同层级下**部分的节点的某种权限。由于树类型资源比较常用，所以我们基于“判断用户是否拥有资源权限”的业务场景，新增了针对判断树这种类型资源节点权限的接口。
 * ## 注意
 * 我们通过`resource`参数定位到树类型数据资源的某一层级（所以该参数是按照`资源 code/节点 code 路径`格式进行传递的），并通过`resourceNodeCodes`参数定位到是当前曾经的哪些节点。
 * ## 场景举例
 * 假如你的业务场景是：当在一个文件系统中，用户在删除某一文件夹下某些文件，需要判断他是否拥有这些文件的删除权限，则可以使用此接口。
 * ## 请求示例
 * ### 判断用户在树资源同层级权限示例（无条件判断）
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "userId": "63721xxxxxxxxxxxxdde14a3",
 * "action": "read",
 * "resource": "treeResourceCode/structCode",
 * "resourceNodeCodes": ["resourceStructChildrenCode1","resourceStructChildrenCode2","resourceStructChildrenCode3"]
 * }
 * ```
 *
 * ### 判断用户在树资源同层级权限示例（开启条件判断）
 *
 * ```json
 * {
 * "namespaceCode": "examplePermissionNamespace",
 * "userId": "63721xxxxxxxxxxxxdde14a3",
 * "action": "read",
 * "resource": "treeResourceCode/structCode",
 * "resourceNodeCodes": ["resourceStructChildrenCode1","resourceStructChildrenCode2","resourceStructChildrenCode3"],
 * "judgeConditionEnabled": true,
 * "authEnvParams":{
 * "ip":"110.96.0.0",
 * "city":"北京",
 * "province":"北京",
 * "country":"中国",
 * "deviceType":"PC",
 * "systemType":"ios",
 * "browserType":"IE",
 * "requestDate":"2022-12-26 17:40:00"
 * }
 * }
 * ```
 *
	* @param requestBody
	* @returns CheckUserSameLevelPermissionResponseDto
*/
func (client *ManagementClient) CheckUserSameLevelPermission(reqDto *dto.CheckUserSameLevelPermissionDto) *dto.CheckUserSameLevelPermissionResponseDto {
	b, err := client.SendHttpRequest("/api/v3/check-user-same-level-permission", fasthttp.MethodPost, reqDto)
	var response dto.CheckUserSameLevelPermissionResponseDto
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
	* @summary 获取权限视图数据列表
	* @description
 * ## 描述
 * 此接口用于导出菜单：权限管理 -> 数据权限 -> 权限试图列表数据，如果你需要拉取我们数据权限的授权数据（所有用户拥有的所有资源的所有权限），则可以使用此接口。
 *
	* @param requestBody
	* @returns ListPermissionViewResponseDto
*/
func (client *ManagementClient) ListPermissionView(reqDto *dto.ListPermissionViewDto) *dto.ListPermissionViewResponseDto {
	b, err := client.SendHttpRequest("/api/v3/list-permission-view/data", fasthttp.MethodPost, reqDto)
	var response dto.ListPermissionViewResponseDto
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
* @summary 获取套餐详情
* @description 获取当前用户池套餐详情。
* @returns CostGetCurrentPackageRespDto
 */
func (client *ManagementClient) GetCurrentPackageInfo() *dto.CostGetCurrentPackageRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-current-package-info", fasthttp.MethodGet, nil)
	var response dto.CostGetCurrentPackageRespDto
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
* @summary 获取用量详情
* @description 获取当前用户池用量详情。
* @returns CostGetCurrentUsageRespDto
 */
func (client *ManagementClient) GetUsageInfo() *dto.CostGetCurrentUsageRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-usage-info", fasthttp.MethodGet, nil)
	var response dto.CostGetCurrentUsageRespDto
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
* @summary 获取 MAU 使用记录
* @description 获取当前用户池 MAU 使用记录
* @param startTime 起始时间（年月日）
* @param endTime 截止时间（年月日）
* @returns CostGetMauPeriodUsageHistoryRespDto
 */
func (client *ManagementClient) GetMauPeriodUsageHistory(reqDto *dto.GetMauPeriodUsageHistoryDto) *dto.CostGetMauPeriodUsageHistoryRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-mau-period-usage-history", fasthttp.MethodGet, reqDto)
	var response dto.CostGetMauPeriodUsageHistoryRespDto
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
* @summary 获取所有权益
* @description 获取当前用户池所有权益
* @returns CostGetAllRightItemRespDto
 */
func (client *ManagementClient) GetAllRightsItem() *dto.CostGetAllRightItemRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-all-rights-items", fasthttp.MethodGet, nil)
	var response dto.CostGetAllRightItemRespDto
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
* @summary 获取订单列表
* @description 获取当前用户池订单列表
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @returns CostGetOrdersRespDto
 */
func (client *ManagementClient) GetOrders(reqDto *dto.GetOrdersDto) *dto.CostGetOrdersRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-orders", fasthttp.MethodGet, reqDto)
	var response dto.CostGetOrdersRespDto
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
* @summary 获取订单详情
* @description 获取当前用户池订单详情
* @param orderNo 订单号
* @returns CostGetOrderDetailRespDto
 */
func (client *ManagementClient) GetOrderDetail(reqDto *dto.GetOrderDetailDto) *dto.CostGetOrderDetailRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-order-detail", fasthttp.MethodGet, reqDto)
	var response dto.CostGetOrderDetailRespDto
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
* @summary 获取订单支付明细
* @description 获取当前用户池订单支付明细
* @param orderNo 订单号
* @returns CostGetOrderPayDetailRespDto
 */
func (client *ManagementClient) GetOrderPayDetail(reqDto *dto.GetOrderPayDetailDto) *dto.CostGetOrderPayDetailRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-order-pay-detail", fasthttp.MethodGet, reqDto)
	var response dto.CostGetOrderPayDetailRespDto
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
* @summary 创建自定义事件应用
* @description 创建自定义事件应用
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) CreateEventApp(reqDto *dto.CreateEventAppDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/create-event-app", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 获取事件应用列表
* @description 获取事件应用列表
* @returns EventAppPaginatedRespDto
 */
func (client *ManagementClient) ListEventApps() *dto.EventAppPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-event-apps", fasthttp.MethodGet, nil)
	var response dto.EventAppPaginatedRespDto
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
* @summary 获取事件列表
* @description 获取 Authing 服务支持的所有事件列表
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @param app 应用类型
* @returns OpenEventPaginatedRespDto
 */
func (client *ManagementClient) ListEvents(reqDto *dto.ListEventsDto) *dto.OpenEventPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-events", fasthttp.MethodGet, reqDto)
	var response dto.OpenEventPaginatedRespDto
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
* @summary 定义自定义事件
* @description 在 Authing 事件中心定义自定义事件
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) DefineEvent(reqDto *dto.DefineEventDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/define-event", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 推送自定义事件
* @description 向 Authing 事件中心推送自定义事件
* @param requestBody
* @returns PubEventRespDto
 */
func (client *ManagementClient) VerifyEvent(reqDto *dto.PubEventDto) *dto.PubEventRespDto {
	b, err := client.SendHttpRequest("/api/v3/pub-event", fasthttp.MethodPost, reqDto)
	var response dto.PubEventRespDto
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
* @summary 认证端推送自定义事件
* @description 认证端向 Authing 事件中心推送自定义事件
* @param requestBody
* @returns PubEventRespDto
 */
func (client *ManagementClient) PubUserEvent(reqDto *dto.PubEventDto) *dto.PubEventRespDto {
	b, err := client.SendHttpRequest("/api/v3/pub-userEvent", fasthttp.MethodPost, reqDto)
	var response dto.PubEventRespDto
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
* @summary 创建注册白名单
* @description 你需要指定注册白名单类型以及数据，来进行创建
* @param requestBody
* @returns WhitelistListRespDto
 */
func (client *ManagementClient) AddWhitelist(reqDto *dto.AddWhitelistDto) *dto.WhitelistListRespDto {
	b, err := client.SendHttpRequest("/api/v3/add-whitelist", fasthttp.MethodPost, reqDto)
	var response dto.WhitelistListRespDto
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
* @summary 获取注册白名单列表
* @description 获取注册白名单列表，可选页数、分页大小来获取
* @param type 白名单类型
* @returns WhitelistListRespDto
 */
func (client *ManagementClient) ListWhitelists(reqDto *dto.ListWhitelistDto) *dto.WhitelistListRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-whitelist", fasthttp.MethodGet, reqDto)
	var response dto.WhitelistListRespDto
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
* @summary 删除注册白名单
* @description 通过指定多个注册白名单数据,以数组的形式进行注册白名单的删除
* @param requestBody
* @returns IsSuccessDto
 */
func (client *ManagementClient) DeleteWhitelist(reqDto *dto.DeleteWhitelistDto) *dto.IsSuccessDto {
	b, err := client.SendHttpRequest("/api/v3/delete-whitelist", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessDto
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
* @summary 获取 ip 列表
* @description 分页获取 ip 列表
* @param ipType IP 类型
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @returns IpListPaginatedRespDto
 */
func (client *ManagementClient) FindIpList(reqDto *dto.IpListDto) *dto.IpListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/ip-list", fasthttp.MethodGet, reqDto)
	var response dto.IpListPaginatedRespDto
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
* @summary 创建 ip 名单
* @description 创建 ip 名单
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) Add(reqDto *dto.IpListCreateDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/ip-list", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 删除 ip 名单
* @description 删除 ip 名单
* @param id
* @returns CommonResponseDto
 */
/*
* @summary 获取用户列表
* @description 分页获取用户列表
* @param userListType IP 类型
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @returns UserListPaginatedRespDto
 */
func (client *ManagementClient) FindUserList(reqDto *dto.UserListDto) *dto.UserListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/user-list", fasthttp.MethodGet, reqDto)
	var response dto.UserListPaginatedRespDto
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
* @summary 创建用户名单
* @description 创建用户名单
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) AddUser(reqDto *dto.UserListCreateDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/user-list", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 删除用户名单
* @description 删除用户 名单
* @param id
* @returns CommonResponseDto
 */
/*
* @summary 获取风险策略列表
* @description 分页获取风险策略列表
* @param optObject 策略操作对象，目前只有 ip
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @returns UserListPolicyPaginatedRespDto
 */
func (client *ManagementClient) FindRiskListPolicy(reqDto *dto.RiskListPolicyDto) *dto.UserListPolicyPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/risk-list-policy", fasthttp.MethodGet, reqDto)
	var response dto.UserListPolicyPaginatedRespDto
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
* @summary 创建风险策略
* @description 创建风险策略
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) AddRiskListPolicy(reqDto *dto.RiskListPolicyCreateDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/risk-list-policy", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 删除风险策略
* @description 删除风险策略
* @param id
* @returns CommonResponseDto
 */
/*
* @summary 新增设备
* @description 创建新设备。
* @param requestBody
* @returns TerminalInfoRespDto
 */
func (client *ManagementClient) CreateDevice(reqDto *dto.CreateTerminalDto) *dto.TerminalInfoRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-device", fasthttp.MethodPost, reqDto)
	var response dto.TerminalInfoRespDto
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
* @summary 最近登录应用
* @description 根据设备唯一标识获取最近登录的应用列表。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) FindLastLoginAppsByDeviceIds(reqDto *dto.QueryTerminalAppsDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/get-last-login-apps-by-deviceIds", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 新增 verify 设备
* @description 创建 verify 新设备。
* @param requestBody
* @returns TerminalInfoRespDto
 */
func (client *ManagementClient) CreateVerifyDevice(reqDto *dto.CreateTerminalDto) *dto.TerminalInfoRespDto {
	b, err := client.SendHttpRequest("/api/v3/authing-verify/create-device", fasthttp.MethodPost, reqDto)
	var response dto.TerminalInfoRespDto
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
* @summary 创建 Pipeline 函数
* @description 创建 Pipeline 函数
* @param requestBody
* @returns PipelineFunctionSingleRespDto
 */
func (client *ManagementClient) CreatePipelineFunction(reqDto *dto.CreatePipelineFunctionDto) *dto.PipelineFunctionSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-pipeline-function", fasthttp.MethodPost, reqDto)
	var response dto.PipelineFunctionSingleRespDto
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
* @summary 获取 Pipeline 函数详情
* @description 获取 Pipeline 函数详情
* @param funcId Pipeline 函数 ID
* @returns PipelineFunctionSingleRespDto
 */
func (client *ManagementClient) GetPipelineFunction(reqDto *dto.GetPipelineFunctionDto) *dto.PipelineFunctionSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-pipeline-function", fasthttp.MethodGet, reqDto)
	var response dto.PipelineFunctionSingleRespDto
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
* @summary 重新上传 Pipeline 函数
* @description 当 Pipeline 函数上传失败时，重新上传 Pipeline 函数
* @param requestBody
* @returns PipelineFunctionSingleRespDto
 */
func (client *ManagementClient) ReuploadPipelineFunction(reqDto *dto.ReUploadPipelineFunctionDto) *dto.PipelineFunctionSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/reupload-pipeline-function", fasthttp.MethodPost, reqDto)
	var response dto.PipelineFunctionSingleRespDto
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
* @summary 修改 Pipeline 函数
* @description 修改 Pipeline 函数
* @param requestBody
* @returns PipelineFunctionSingleRespDto
 */
func (client *ManagementClient) UpdatePipelineFunction(reqDto *dto.UpdatePipelineFunctionDto) *dto.PipelineFunctionSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-pipeline-function", fasthttp.MethodPost, reqDto)
	var response dto.PipelineFunctionSingleRespDto
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
* @summary 修改 Pipeline 函数顺序
* @description 修改 Pipeline 函数顺序
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) UpdatePipelineOrder(reqDto *dto.UpdatePipelineOrderDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/update-pipeline-order", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 删除 Pipeline 函数
* @description 删除 Pipeline 函数
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) DeletePipelineFunction(reqDto *dto.DeletePipelineFunctionDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/delete-pipeline-function", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
	* @summary 获取 Pipeline 函数列表
	* @description 获取 Pipeline 函数列表
	* @param scene 通过函数的触发场景进行筛选（可选，默认返回所有）：
 * - `PRE_REGISTER`: 注册前
 * - `POST_REGISTER`: 注册后
 * - `PRE_AUTHENTICATION`: 认证前
 * - `POST_AUTHENTICATION`: 认证后
 * - `PRE_OIDC_ID_TOKEN_ISSUED`: OIDC ID Token 签发前
 * - `PRE_OIDC_ACCESS_TOKEN_ISSUED`: OIDC Access Token 签发前
 * - `PRE_COMPLETE_USER_INFO`: 补全用户信息前
 *
	* @returns PipelineFunctionPaginatedRespDto
*/
func (client *ManagementClient) ListPipelineFunctions(reqDto *dto.ListPipelineFunctionsDto) *dto.PipelineFunctionPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-pipeline-functions", fasthttp.MethodGet, reqDto)
	var response dto.PipelineFunctionPaginatedRespDto
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
* @summary 获取 Pipeline 日志
* @description 获取 Pipeline
* @param funcId Pipeline 函数 ID
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @returns PipelineFunctionPaginatedRespDto
 */
func (client *ManagementClient) GetPipelineLogs(reqDto *dto.GetPipelineLogsDto) *dto.PipelineFunctionPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-pipeline-logs", fasthttp.MethodGet, reqDto)
	var response dto.PipelineFunctionPaginatedRespDto
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
* @summary 创建 Webhook
* @description 你需要指定 Webhook 名称、Webhook 回调地址、请求数据格式、用户真实名称来创建 Webhook。还可选是否启用、请求密钥进行创建
* @param requestBody
* @returns CreateWebhookRespDto
 */
func (client *ManagementClient) CreateWebhook(reqDto *dto.CreateWebhookDto) *dto.CreateWebhookRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-webhook", fasthttp.MethodPost, reqDto)
	var response dto.CreateWebhookRespDto
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
* @summary 获取 Webhook 列表
* @description 获取 Webhook 列表，可选页数、分页大小来获取
* @param page 当前页数，从 1 开始
* @param limit 每页数目，最大不能超过 50，默认为 10
* @returns GetWebhooksRespDto
 */
func (client *ManagementClient) ListWebhooks(reqDto *dto.ListWebhooksDto) *dto.GetWebhooksRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-webhooks", fasthttp.MethodGet, reqDto)
	var response dto.GetWebhooksRespDto
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
* @summary 修改 Webhook 配置
* @description 需要指定 webhookId，可选 Webhook 名称、Webhook 回调地址、请求数据格式、用户真实名称、是否启用、请求密钥参数进行修改 webhook
* @param requestBody
* @returns UpdateWebhooksRespDto
 */
func (client *ManagementClient) UpdateWebhook(reqDto *dto.UpdateWebhookDto) *dto.UpdateWebhooksRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-webhook", fasthttp.MethodPost, reqDto)
	var response dto.UpdateWebhooksRespDto
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
* @summary 删除 Webhook
* @description 通过指定多个 webhookId,以数组的形式进行 webhook 的删除,如果 webhookId 不存在,不提示报错
* @param requestBody
* @returns DeleteWebhookRespDto
 */
func (client *ManagementClient) DeleteWebhook(reqDto *dto.DeleteWebhookDto) *dto.DeleteWebhookRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-webhook", fasthttp.MethodPost, reqDto)
	var response dto.DeleteWebhookRespDto
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
* @summary 获取 Webhook 日志
* @description 通过指定 webhookId，可选 page 和 limit 来获取 webhook 日志,如果 webhookId 不存在,不返回报错信息
* @param requestBody
* @returns ListWebhookLogsRespDto
 */
func (client *ManagementClient) GetWebhookLogs(reqDto *dto.ListWebhookLogs) *dto.ListWebhookLogsRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-webhook-logs", fasthttp.MethodPost, reqDto)
	var response dto.ListWebhookLogsRespDto
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
* @summary 手动触发 Webhook 执行
* @description 通过指定 webhookId，可选请求头和请求体进行手动触发 webhook 执行
* @param requestBody
* @returns TriggerWebhookRespDto
 */
func (client *ManagementClient) TriggerWebhook(reqDto *dto.TriggerWebhookDto) *dto.TriggerWebhookRespDto {
	b, err := client.SendHttpRequest("/api/v3/trigger-webhook", fasthttp.MethodPost, reqDto)
	var response dto.TriggerWebhookRespDto
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
* @summary 获取 Webhook 详情
* @description 根据指定的 webhookId 获取 webhook 详情
* @param webhookId Webhook ID
* @returns GetWebhookRespDto
 */
func (client *ManagementClient) GetWebhook(reqDto *dto.GetWebhookDto) *dto.GetWebhookRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-webhook", fasthttp.MethodGet, reqDto)
	var response dto.GetWebhookRespDto
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
* @summary 获取 Webhook 事件列表
* @description 返回事件列表和分类列表
* @returns WebhookEventListRespDto
 */
func (client *ManagementClient) GetWebhookEventList() *dto.WebhookEventListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-webhook-event-list", fasthttp.MethodGet, nil)
	var response dto.WebhookEventListRespDto
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
* @summary 生成 LDAP Server 管理员密码
* @description 生成 LDAP Server 管理员密码
* @returns LdapGetBindPwdRespDto
 */
func (client *ManagementClient) GetBindPwd() *dto.LdapGetBindPwdRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-ldap-server-random-pwd", fasthttp.MethodGet, nil)
	var response dto.LdapGetBindPwdRespDto
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
* @summary 获取 LDAP server 配置信息
* @description 获取 LDAP server 配置信息
* @returns LdapConfigInfoRespDto
 */
func (client *ManagementClient) QueryLdapConfigInfo() *dto.LdapConfigInfoRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-ldap-server-config", fasthttp.MethodGet, nil)
	var response dto.LdapConfigInfoRespDto
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
* @summary 更新 LDAP server 配置信息
* @description 更新 LDAP server 配置信息
* @param requestBody
* @returns LdapOperateRespDto
 */
func (client *ManagementClient) UpdateLdapConfigInfo(reqDto *dto.LdapUpdateDto) *dto.LdapOperateRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-ldap-server-config", fasthttp.MethodPost, reqDto)
	var response dto.LdapOperateRespDto
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
* @summary 初始化/重启 LDAP server
* @description 初始化/重启 LDAP server
* @param requestBody
* @returns LdapOperateRespDto
 */
func (client *ManagementClient) SaveLdapConfigInfo(reqDto *dto.LdapSaveDto) *dto.LdapOperateRespDto {
	b, err := client.SendHttpRequest("/api/v3/enable-ldap-server", fasthttp.MethodPost, reqDto)
	var response dto.LdapOperateRespDto
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
* @summary 关闭 LDAP server 服务，关闭前必须先初始化
* @description 关闭 LDAP server 服务，关闭前必须先初始化
* @param requestBody
* @returns LdapOperateRespDto
 */
func (client *ManagementClient) DisableLdapServer(reqDto *dto.LdapSetEnabledFlagDto) *dto.LdapOperateRespDto {
	b, err := client.SendHttpRequest("/api/v3/disable-ldap-server", fasthttp.MethodPost, reqDto)
	var response dto.LdapOperateRespDto
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
* @summary LDAP server 日志查询
* @description LDAP server 日志查询
* @param type 类型：1 访问日志，2 错误日志
* @param page 当前页,从 1 开始
* @param limit 每页条数
* @param connection 连接标识
* @param operationNumber 操作码
* @param errorCode 错误码
* @param message 消息内容
* @param startTime 开始时间-时间戳
* @param endTime 结束时间-时间戳
* @returns LdapLogRespDto
 */
func (client *ManagementClient) QueryLdapLog(reqDto *dto.GetLdapServerLogDto) *dto.LdapLogRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-ldap-server-log", fasthttp.MethodGet, reqDto)
	var response dto.LdapLogRespDto
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
* @summary LDAP server 根据 DN 查询下一级
* @description LDAP server 根据 DN 查询下一级
* @param page 当前页,从 1 开始
* @param limit 每页条数
* @param dn 当前 DN
* @returns LdapLogRespDto
 */
func (client *ManagementClient) QueryLdapSubEntries(reqDto *dto.GetLdapSubEntriesDto) *dto.LdapLogRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-ldap-sub-entries", fasthttp.MethodGet, reqDto)
	var response dto.LdapLogRespDto
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
* @summary 获取协作管理员 AK/SK 列表
* @description 根据协作管理员 Id 获取协作管理员下所有的 AK/SK 列表
* @param userId 密钥所属用户 ID
* @param tenantId 密钥所属租户 ID
* @param type 密钥类型
* @param status AccessKey 状态，activated：已激活，staging：分级（可轮换），revoked：已撤销
* @returns ListAccessKeyResponseDto
 */
func (client *ManagementClient) GetAccessKeyList(reqDto *dto.ListAccessKeyDto) *dto.ListAccessKeyResponseDto {
	b, err := client.SendHttpRequest("/api/v3/list-access-key", fasthttp.MethodGet, reqDto)
	var response dto.ListAccessKeyResponseDto
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
* @summary 获取协作管理员 AK/SK 详细信息
* @description 获取协作管理员 AK/SK 详细信息,根据协作管理员 ID 和 accessKeyId 获取对应 AK/SK 的详细信息。
* @param userId 用户 ID
* @param accessKeyId accessKeyId
* @returns GetAccessKeyResponseDto
 */
func (client *ManagementClient) GetAccessKey(reqDto *dto.GetAccessKeyDto) *dto.GetAccessKeyResponseDto {
	b, err := client.SendHttpRequest("/api/v3/get-access-key", fasthttp.MethodGet, reqDto)
	var response dto.GetAccessKeyResponseDto
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
* @summary 创建协作管理员的 AK/SK
* @description 创建协作管理员的 AK/SK,根据协作管理员 ID 生成指定的 AK/SK。
* @param requestBody
* @returns CreateAccessKeyResponseDto
 */
func (client *ManagementClient) CreateAccessKey(reqDto *dto.CreateAccessKeyDto) *dto.CreateAccessKeyResponseDto {
	b, err := client.SendHttpRequest("/api/v3/create-access-key", fasthttp.MethodPost, reqDto)
	var response dto.CreateAccessKeyResponseDto
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
* @summary 删除协作管理员的 AK/SK
* @description 删除协作管理员的 AK/SK,根据所选择的 AK/SK 的 accessKeyId 进行指定删除。
* @param requestBody
* @returns CommonResponseDto
 */
func (client *ManagementClient) DeleteAccessKey(reqDto *dto.DeleteAccessKeyDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/delete-access-key", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
* @summary 更新一个管理员 AccessKey
* @description 根据 AccessKeyId 更新一个管理员 AccessKey，目前只支持更新 status，status 支持 activated / revoked
* @param requestBody
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) UpdateAccessKey(reqDto *dto.UpdateAccessKeyDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-access-key", fasthttp.MethodPost, reqDto)
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
* @summary 获取 verify-config 配置
* @description 获取 verify-config 配置
* @returns VerifyConfigInfoRespDto
 */
func (client *ManagementClient) GetVerifyConfig() *dto.VerifyConfigInfoRespDto {
	b, err := client.SendHttpRequest("/api/v3/verify-config", fasthttp.MethodGet, nil)
	var response dto.VerifyConfigInfoRespDto
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
* @summary 修改 verify-config 配置
* @description 修改 verify-config 配置
* @param requestBody
* @returns VerifyConfigDataDto
 */
/*
* @summary 获取 verify-config-app 列表
* @description 获取 verify-config-app 列表
* @param keywords 搜索关键字
* @returns ApplicationDto
 */
func (client *ManagementClient) GetVerifyConfigApp(reqDto *dto.VerifyConfigAppDto) *dto.ApplicationDto {
	b, err := client.SendHttpRequest("/api/v3/verify-config-app", fasthttp.MethodGet, reqDto)
	var response dto.ApplicationDto
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
* @summary 事件发布
* @description 根据事件编码发布一个自定义事件
* @param eventCode 事件编码
* @param body 事件消息
* @returns IsSuccessRespDto
 */
func (client *ManagementClient) PubEvent(eventCode string, data interface{}) *dto.IsSuccessRespDto {
	var reqDto = dto.NewEventReqDto(eventCode, data)
	b, err := client.SendHttpRequest("/api/v3/pub-event", fasthttp.MethodPost, reqDto)
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
* @summary 事件订阅
* @description 根据事件编码订阅一个自定义事件
* @param eventCode 事件编码
* @param onSuccess 成功的消息
* @param onError 失败处理
 */
func (client *ManagementClient) SubEvent(eventCode string, onSuccess func(msg []byte), onError func(err error)) {
	var options = client.options
	defMap := make(map[string]string, 0)
	stringToSign := util.ComposeStringToSign("websocket", "", defMap, defMap)
	token := util.GetAuthorization(client.options.AccessKeyId, client.options.AccessKeySecret, stringToSign)
	// fmt.Println(token)
	if !client.eventHub.CreateManagement(eventCode, options.WssHost, token) {
		return
	}
	client.eventHub.AddReceiver(eventCode, onSuccess, onError)
	// recv message exec corresponding callback function
	go client.eventHub.StartReceive(eventCode)
}

/*
* @summary 事件订阅
* @description 根据事件编码订阅一个自定义事件
* @param eventCode 事件编码
* @param receiver 消息处理器
 */
func (client *ManagementClient) SubEventByReceiver(eventCode string, receiver util.EventReceiver) {
	client.SubEvent(eventCode, receiver.OnSuccess, receiver.OnError)
}
