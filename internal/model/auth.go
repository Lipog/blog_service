package model

import "github.com/jinzhu/gorm"

type Auth struct {
	*Model
	AppKey string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

//用于服务端在获取客户传入的app_key和app_secret后进行验证，是否真的存在这样一条数据
func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? and app_secret = ? and is_del = ?",
		a.AppKey, a.AppSecret, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}
	return auth, nil
}