package validator

import (
	"fmt"
	"reflect"
)

//配置信息
type Options struct {
	field string
	Data  map[string][]*Rule
}

func (this *Options) Add(rule string, message string) *Options {
	r := NewRule(rule, message)
	if tmp, ok := this.Data[this.field]; ok {
		this.Data[this.field] = append(tmp, r)
	} else {
		this.Data[this.field] = []*Rule{r}
	}
	return this
}

//结构体校验器
type Validator struct {
	//所有校验规则
	Options *Options
	//自定义规则
	Custom map[string]FN
}

func New() *Validator {
	tmp := new(Validator)
	tmp.Options = &Options{Data: map[string][]*Rule{}}
	tmp.Custom = map[string]FN{}
	return tmp
}

func (this *Validator) Rule(field string) *Options {
	this.Options.field = field
	return this.Options
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
		rules, ok := this.Options.Data[key.Name]
		if !ok {
			continue
		}
		for _, rule := range rules {
			//优先使用自定义校验规则
			fn, ok := this.Custom[rule.Name]
			if !ok {
				fn, ok = rulePool[rule.Name]
			}
			//根据规则名称，没有找到相关规则
			if !ok {
				e := fmt.Errorf("not found rule: %s", rule.Name)
				if tmp, ok := errsBag[key.Name]; ok {
					errsBag[key.Name] = append(tmp, e)
				} else {
					errsBag[key.Name] = []error{e}
				}
				continue
			}
			//调用校验函数
			if ok, err := fn(key.Name, value, rule); err == nil {
				if !ok {
					if tmp, ok := result[key.Name]; ok {
						result[key.Name] = append(tmp, rule.Message)
					} else {
						result[key.Name] = []string{rule.Message}
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
