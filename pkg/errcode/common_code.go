package errcode

var (
	Success = NewError(0, "成功")
	ServerError = NewError(10000000, "服务器内部错误")
	InvalidParams = NewError(1000001, "入参错误")
	NotFound = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist = NewError(10000003, "鉴权失败，找不到对应的APPKey和AppSecret")
	UnauthorizedTokenError = NewError(10000004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout = NewError(10000005, "Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "Token生成失败")
	TooManyRequests = NewError(10000007, "请求过多")
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail = NewError(20010004, "删除标签失败")
	ErrorCountTagFail = NewError(20010005, "统计标签失败")
	ErrorGetArticleFail = NewError(2002001, "获取单个文章失败")
	ErrorGetArticlesFail = NewError(2002002, "获取多个文章失败")
	ErrorCreateArticleFail = NewError(2002003, "创建文章失败")
	ErrorUpdateArticleFail = NewError(2002004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(2002005, "删除文章失败")
)