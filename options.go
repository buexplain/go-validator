package validator

//配置信息
type Options struct {
	currentField string
	//字段名与校验规则
	Data  map[string][]*Rule
	//字段的别名
	Alias map[string]string
}

func newOptions() *Options {
	tmp := &Options{Data: map[string][]*Rule{}}
	tmp.Alias = map[string]string{}
	return tmp
}

func (this Options) Clone() *Options {
	tmp := newOptions()
	tmp.currentField = this.currentField
	for k, v := range this.Data {
		tmp.Data[k] = make([]*Rule, len(v), len(v))
		for kk, vv := range v {
			tmp.Data[k][kk] = vv.Clone()
		}
	}
	for k, v := range this.Alias {
		tmp.Alias[k] = v
	}
	return tmp
}

func (this *Options) Add(rule string, message ...string) *Options {
	r := newRule(rule, message...)
	if tmp, ok := this.Data[this.currentField]; !ok {
		this.Data[this.currentField] = []*Rule{r}
	} else {
		for k, v := range tmp {
			if v.name == r.name {
				//重复的规则会被覆盖掉
				this.Data[this.currentField][k] = r
				return this
			}
		}
		this.Data[this.currentField] = append(tmp, r)
	}
	return this
}
