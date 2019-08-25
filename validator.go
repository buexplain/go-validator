package validator

import (
	"fmt"
	"reflect"
)

//结构体校验器
type Validator struct {
	//所有校验规则
	options *Options
	//自定义规则
	custom map[string]FN
}

func New() *Validator {
	tmp := new(Validator)
	tmp.options = &Options{Data: map[string][]*Rule{}}
	tmp.custom = map[string]FN{}
	return tmp
}

func (this Validator) Clone() *Validator {
	tmp := new(Validator)
	tmp.options = this.options.Clone()
	tmp.custom = map[string]FN{}
	for k, v := range this.custom {
		tmp.custom[k] = v
	}
	return tmp
}

func (this *Validator) Custom(name string, fn FN) {
	this.custom[name] = fn
}

func (this *Validator) Rule(field string) *Options {
	this.options.field = field
	return this.options
}

func (this *Validator) Validate(structVar interface{}) (Result, error) {
	reflectType := reflect.TypeOf(structVar)
	reflectValue := reflect.ValueOf(structVar)

	result := Result{}
	errsBag := ErrorBag{}

	//转换结构体指针
	if k := reflectType.Kind(); k == reflect.Ptr {
		reflectType = reflectType.Elem()
		reflectValue = reflectValue.Elem()
	}

	if k := reflectType.Kind(); k != reflect.Struct {
		errsBag["validator"] = []error{fmt.Errorf("The validator parameter must be Struct type")}
		return result, errsBag
	}

	//遍历所有字段
	for i := 0; i < reflectType.NumField(); i++ {
		key := reflectType.Field(i)
		value := reflectValue.Field(i).Interface()
		rules, ok := this.options.Data[key.Name]
		if !ok {
			continue
		}
		for _, rule := range rules {
			//优先使用自定义校验规则
			fn, ok := this.custom[rule.name]
			if !ok {
				fn, ok = rulePool[rule.name]
			}
			//根据规则名称，没有找到相关规则
			if !ok {
				e := fmt.Errorf("not found rule: %s", rule.name)
				if tmp, ok := errsBag[key.Name]; ok {
					errsBag[key.Name] = append(tmp, e)
				} else {
					errsBag[key.Name] = []error{e}
				}
				continue
			}
			//调用校验函数
			if message, err := fn(key.Name, value, rule); err == nil {
				if message != "" {
					if tmp, ok := result[key.Name]; ok {
						result[key.Name] = append(tmp, message)
					} else {
						result[key.Name] = []string{message}
					}
					break
				}
			} else {
				if tmp, ok := errsBag[key.Name]; ok {
					errsBag[key.Name] = append(tmp, err)
				} else {
					errsBag[key.Name] = []error{err}
				}
				break
			}
		}
	}

	if len(errsBag) == 0 {
		return result, nil
	}

	return result, errsBag
}
