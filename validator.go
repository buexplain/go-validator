package validator

import (
	"fmt"
	"reflect"
)

type Option struct {
	//规则名称
	Name string
	//校验失败的提示信息
	Message string
}

//结构体校验器
type Validator struct {
	//所有校验规则
	Options map[string][]Option
	//自定义规则
	Custom map[string]FN
}

func New() *Validator {
	tmp := new(Validator)
	tmp.Options = map[string][]Option{}
	tmp.Custom = map[string]FN{}
	return tmp
}

func (this *Validator) Rule(field string, rule string, message string) *Validator {
	r := Option{Name: rule, Message: message}
	if tmp, ok := this.Options[field]; ok {
		this.Options[field] = append(tmp, r)
	} else {
		this.Options[field] = []Option{r}
	}
	return this
}

func (this *Validator) Validate(a interface{}) (Result, error) {
	reflectType := reflect.TypeOf(a)
	reflectValue := reflect.ValueOf(a)

	result := Result{}
	errsBag := ErrorBag{}

	if k := reflectType.Kind(); k != reflect.Struct {
		errsBag["validator"] = []error{fmt.Errorf("The validator parameter must be Struct type")}
		return result, errsBag
	}

	//遍历所有字段
	for i := 0; i < reflectType.NumField(); i++ {
		key := reflectType.Field(i)
		value := reflectValue.Field(i).Interface()
		options, ok := this.Options[key.Name]
		if !ok {
			continue
		}
		for _, option := range options {
			//优先使用自定义校验规则
			fn, ok := this.Custom[option.Name]
			if !ok {
				fn, ok = rules[option.Name]
			}
			//根据规则名称，没有找到相关规则
			if !ok {
				e := fmt.Errorf("not found rule: %s", option.Name)
				if tmp, ok := errsBag[key.Name]; ok {
					errsBag[key.Name] = append(tmp, e)
				} else {
					errsBag[key.Name] = []error{e}
				}
				continue
			}
			//调用校验函数
			if ok, err := fn(key.Name, value); err == nil {
				if !ok {
					if tmp, ok := result[key.Name]; ok {
						result[key.Name] = append(tmp, option.Message)
					} else {
						result[key.Name] = []string{option.Message}
					}
				}
			} else {
				if tmp, ok := errsBag[key.Name]; ok {
					errsBag[key.Name] = append(tmp, err)
				} else {
					errsBag[key.Name] = []error{err}
				}
			}
		}
	}

	if len(errsBag) == 0 {
		return result, nil
	}

	return result, errsBag
}
