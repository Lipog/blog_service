package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettingS
	AppSetting *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS //数据库初始化所需要的参数
	Logger *logger.Logger //用于日志组件的初始化
	JWTSetting *setting.JWTSettingS
)
