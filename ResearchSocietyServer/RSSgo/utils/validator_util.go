package utils

import (
	"log"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type ValidatorUtil struct {
}

func (ValidatorUtil) GetError(err error, r interface{}) string {
	defer func() {
		err := recover() //内置函数，可以捕捉到函数异常
		if err != nil {
			//这里是打印错误，还可以进行报警处理，例如微信，邮箱通知
			log.Println("err错误信息：", err)
		}
	}()
	errs := err.(validator.ValidationErrors)
	s := reflect.TypeOf(r)
	for _, fieldError := range errs {
		filed, _ := s.FieldByName(fieldError.Field())
		// 获取错误信息
		errText := filed.Tag.Get(fieldError.Tag() + "_msg")
		if errText != "" {
			return errText
		}
		return fieldError.Field() + "格式错误"
	}
	return ""
}
