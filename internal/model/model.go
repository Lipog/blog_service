package model

import (
	"blog-service/global"
	"blog-service/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Model struct {
	ID uint32 `gorm:"primary_key" json"id"`
	CreatedBy string `json"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeleteOn uint32 `json:"delete_on"`
	IsDel uint8 `json:"is_del"`
}

//注册回调行为
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=%t&loc=Local",
			databaseSetting.UserName,
			databaseSetting.Password,
			databaseSetting.Host,
			databaseSetting.DBName,
			databaseSetting.ParseTime,
			))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	//这里是新增的注册回调的部分
	db.Callback().Create().Replace("gorm:update_time_stamp",
		updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp",
		updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	//结束
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	return db, nil
}

//新增行为的回调
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}
		//通过scope.FieldByName方法，获取当前是否包含所需要的字段
		//通过Field.IsBlank可以得知该字段的值是否为空
		//若为空，则调用Field.Set方法给该字段设置值
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

//更新行为的回调
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	//通过调用socpe.Get来获取当前设置的标识gorm:update_column的字段属性
	if _, ok := scope.Get("gorm:update_column"); !ok {
		//如果不存在，即没有自定义的update_column，则在更新回调内设置默认字段
		//ModifiedOn的值为当前的时间戳
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

//删除回调的行为
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(
				fmt.Sprintf(
					"UPDATE %v SET %v=%v, %v=%v%v%v",
					scope.QuotedTableName(),
					scope.Quote(deletedOnField.DBName),
					scope.AddToVars(now),
					scope.Quote(isDelField.DBName),
					scope.AddToVars(1),
					addExtraSpaceIfExist(scope.CombinedConditionSql()),
					addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
				)).Exec()
		}
	}


}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
