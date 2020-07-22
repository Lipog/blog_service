package service

//在Request结构体中，应用了两个标签，分别为form和binding
//分蘖代表表单的映射字段名和入参校验的规则内容，主要功能是实现参数半定和参数校验
type CountTagRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name string `form:"name" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name string `form:"name" binding:"required, min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required, min=3, max=10"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
	Name string `form:"name" binding:"min=3,max=100"`
	State uint8 `form:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
