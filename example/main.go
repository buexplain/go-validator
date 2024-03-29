package main

import (
	"fmt"
	"github.com/buexplain/go-validator"
	"os"
)

type User struct {
	//自增id
	ID int
	//账号
	Account string
	//密码
	Password string
	//昵称
	Nickname string
	//用户状态 1=禁止 2=允许
	Status int
	//用户角色 admin=普通管理员 superAdmin=超级管理员
	Role string
}

//全局的校验类，只读，不支持并发写
var validatorObj *validator.Validator

func init() {
	validatorObj = validator.New()
	//通过自定义规则，来检查用户账号是否唯一
	validatorObj.Custom("account_unique", func(field string, value interface{}, rule *validator.Rule, structVar interface{}) (s string, e error) {
		fmt.Printf("account_unique: field=%s value=%s rule= %+v\n", field, value, rule)
		return "", nil
	})

	validatorObj.Field("ID").Rule("required", "请输入用户ID").Rule("positive_numeric", "错误的用户ID")

	validatorObj.Field("Account").
		Rule("email", "请输入邮箱", "请输入正确的邮箱").
		Rule("account_unique", "该账号已经存在，请重新输入")

	validatorObj.Field("Password").
		Rule("password:min=8&max=64",
			"请输入密码",
			"密码格式有误，请输入8位以上64位以下的密码",
			"密码格式有误，请输入数字、字母、符号",
			"密码格式有误，数字、字母、符号至少两种",
		)

	validatorObj.Field("Nickname", "昵称").Rule("between:min=1&max=7", "请输入昵称", "昵称字符在1~7个之间")

	validatorObj.Field("Status", "用户状态").Rule("between:min=1&max=2", "请选择用户状态", "错误的用户状态值")

	validatorObj.Field("Role", "角色").Rule(`in:in=admin,superAdmin,user&split=\,`, "请选择用户身份", "错误的用户身份值")
}

func createUser() {
	user := &User{}
	user.ID = 1
	user.Account = "x@x.x"
	user.Password = `123~！#￥%……&*（）——+~!@#$%^&*()_+，。/；‘、】【,./;'[]\'_:"><"`
	user.Nickname = "123我爱你"
	user.Status = 1
	user.Role = "admin,superAdmin"

	result, err := validatorObj.Validate(*user)

	if err != nil {
		fmt.Println("校验器出错\n", err)
		os.Exit(1)
	} else {
		if !result.IsEmpty() {
			fmt.Println("校验失败", result.String())
			fmt.Println(result)
			os.Exit(1)
		}
	}

	fmt.Println("校验通过，create user")
}

func updateUser() {
	user := &User{}
	user.ID = 1
	user.Account = "x@x.x"
	user.Nickname = "123我爱你"
	user.Password = ""
	user.Status = 1
	user.Role = "admin,superAdmin"

	//重写 password 规则，改为密码有值，则校验，无值，则通过
	v := validatorObj.Clone()
	v.Custom("password", func(field string, value interface{}, rule *validator.Rule, structVar interface{}) (s string, e error) {
		str, _ := value.(string)
		if str == "" {
			return "", nil
		}
		//从规则池中，拿出 password 规则，继续使用
		return validator.Pool("password")(field, value, rule, structVar)
	})

	result, err := v.Validate(user)

	if err != nil {
		fmt.Println("校验器出错")
		fmt.Println(err)
		os.Exit(1)
	} else {
		if !result.IsEmpty() {
			fmt.Println("校验失败", result.String())
			os.Exit(1)
		}

	}

	fmt.Println("校验通过 update user")
}

func main() {
	createUser()
	updateUser()
}
