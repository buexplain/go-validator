package main

import (
	"fmt"
	"github.com/buexplain/go-validator"
	"os"
)

type User struct {
	//账号
	Account string
	//密码
	Password string
	//昵称
	Nickname string
	//用户状态
	Status int
}

var validatorObj *validator.Validator

func init()  {
	validatorObj = validator.New()
	validatorObj.Rule("Account").Add("mailbox", "请输入正确的邮箱")
	validatorObj.Rule("Nickname").Add("between:min=1,max=7", "昵称长度必须在一到")
	validatorObj.Rule("Status").Add("between:min=1,max=2", "请选择用户状态")
}

func main() {
	v := validator.New()
	v.Custom("custom", func(field string, value interface{}, rule *validator.Rule) (bool, error) {
		fmt.Println(rule)
		return false, nil
	})

	v.Rule("Name").Add("required", "请输入Name").Add("custom:min=100,max=200", "自定义规则校验失败")
	v.Rule("URL").Add("required", "请输入URL")

	vv := v.Clone()

	vv.Custom("xxx",  func(field string, value interface{}, rule *validator.Rule) (b bool, e error) {
		return true,nil
	})


	user := &User{}

	result, err := v.Validate(*user)

	if err != nil {
		fmt.Println("校验器出错", err)
		os.Exit(1)
	} else {
		if len(result) > 0 {
			fmt.Println("校验失败")
			fmt.Println(result)
			os.Exit(1)
		}
	}

	fmt.Println("校验通过")
}
