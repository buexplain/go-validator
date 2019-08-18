# validator

一个简单的`Go`结构体校验类

## 示例
```go
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
	v.Custom["test"] = func(field string, value interface{}) (bool, error) {
		return false, nil
	}
	v.Rule("Name", "required", "请输入Name")
	v.Rule("Name", "test", "自定义规则校验失败")
	v.Rule("URL", "required", "请输入URL")
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
```

## 可用的验证规则
* `required` 验证的字段必须存在于输入数据中，而不是空。

## License
[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0.html)
