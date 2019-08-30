package validator

import (
	"fmt"
	"reflect"
)

//结构体校验器
type Validator struct {
	//所有的字段与其对应的校验的规则的集合
	options *Options
	//自定义规则
	custom map[string]FN
}

func New() *Validator {
	tmp := new(Validator)
	tmp.options = newOptions()
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

//添加自定义规则
func (this *Validator) Custom(ruleName string, fn FN) {
	this.custom[ruleName] = fn
}

//给字段添加规则
func (this *Validator) Rule(field string, alias ...string) *Options {
	this.options.currentField = field
	if len(alias) > 0 {
		this.options.Alias[field] = alias[0]
	}else {
		this.options.Alias[field] = field
	}
	return this.options
}

//校验一个结构体
func (this *Validator) Validate(structVar interface{}) (*Result, error) {
	reflectType := reflect.TypeOf(structVar)
	reflectValue := reflect.ValueOf(structVar)

	result := newResult()
	errsBag := newErrorBag()

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
		alias, ok := this.options.Alias[key.Name]
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
				err := fmt.Errorf("not found rule: %s", rule.name)
				errsBag.Add(key.Name, err)
				continue
			}
			//调用校验函数
			if message, err := fn(key.Name, value, rule); err == nil {
				if message != "" {
					result.Add(alias, message)
					break
				}
			} else {
				errsBag.Add(key.Name, err)
				break
			}
		}
	}

	if len(errsBag) == 0 {
		return result, nil
	}

	return result, errsBag
}
