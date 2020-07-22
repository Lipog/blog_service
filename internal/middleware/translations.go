package middleware
//国际化处理
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/gin-gonic/gin/binding"
//	"github.com/go-playground/locales/en"
//	"github.com/go-playground/locales/zh"
//	"github.com/go-playground/locales/zh_Hant_TW"
//	ut "github.com/go-playground/universal-translator"
//	"gopkg.in/go-playground/validator.v8"
//)
//
//func Translations() *gin.HandlerFunc {
//	return func(c *gin.Context) {
//		uni := ut.New(en.New(),zh.New(),zh_Hant_TW.New())
//		locale := c.GetHeader("locale")
//		trans, _ := uni.GetTranslator(locale)
//		v, ok := binding.Validator.Engine().(*validator.Validate)
//		if ok {
//			switch locale {
//			case "zh":
//				_ = zh_tra
//			}
//		}
//	}
//}
