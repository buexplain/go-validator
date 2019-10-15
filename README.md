# validator

一个简单的`Go`结构体校验类

## 示例
[example](https://github.com/buexplain/go-validator/tree/master/example/main.go)

## 可用的校验规则
* `required` 校验非零值。
* `between` 校验范围 min与max是一段连续的区间。
* `alpha` 校验字母。
* `numeric` 校验数字。
* `positive_numeric` 校验正正数。
* `alpha_numeric` 校验字母与数字。
* `alpha_numeric_dash` 校验字母与数字，以及破折号和下划线。
* `email` 校验邮箱。
* `password` 校验密码强度，min到max个字符，数字、字母、特殊符号至少两种。
* `in` 校验范围，这个范围不是一段连续的区间，类似于白名单校验，支持校验`slice`，与传递`split=\,`后，切割成`slice`再校验。

## 规则参数的输入与解析
```text
规则名称:参数名1=参数值&参数名2=参数值1,参数值2
```
如果要输入字符 `:` `=` `&` `,`那么可以使用反斜杠进行转义。

示例如下

输入：
```text
ruleName:param1=1&param2=get,post&param3=a,b,\,,\&,\=,c,:
```
解析：
```text
map[param1:[1] param2:[get post] param3:[a b , & = c :]]
```

## License
[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0.html)
