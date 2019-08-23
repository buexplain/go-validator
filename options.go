package validator


//配置信息
type Options struct {
	field string
	Data  map[string][]*Rule
}

func (this Options) Clone() *Options {
	tmp := new(Options)
	tmp.field = this.field
	tmp.Data = make(map[string][]*Rule)
	for k, v := range this.Data {
		tmp.Data[k] = make([]*Rule, len(v), len(v))
		for kk, vv := range v {
			tmp.Data[k][kk] = vv.Clone()
		}
	}
	return tmp
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

