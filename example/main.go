package main

import (
	"fmt"
	"github.com/buexplain/go-validator"
	"os"
)

type Node struct {
	Name string
	URL  string
}

func main() {
	v := validator.New()
	v.Custom["custom"] = func(field string, value interface{}, rule *validator.Rule) (bool, error) {
		fmt.Println(rule)
		return false, nil
	}
	v.Rule("Name").Add("required", "请输入Name").Add("custom:min=100,max=200", "自定义规则校验失败")
	v.Rule("URL").Add("required", "请输入URL")

	n := &Node{}

	result, err := v.Validate(*n)

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
