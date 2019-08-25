# validator

一个简单的`Go`结构体校验类

## 示例
[example](https://github.com/buexplain/go-validator/tree/master/example/main.go)

## 可用的校验规则
* `required` 校验非零值。
* `between` 校验范围 min与max是一段连续的区间。
* `alpha` 校验字母。
* `numeric` 校验数字。
* `alpha_numeric` 校验字母与数字。
* `alpha_numeric_dash` 校验字母与数字，以及破折号和下划线。
* `email` 校验邮箱。
* `password` 校验密码强度，min到max个字符，数字、字母、特殊符号至少两种。
* `in` 校验范围，这个范围不是一段连续的区间。

## License
[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0.html)
