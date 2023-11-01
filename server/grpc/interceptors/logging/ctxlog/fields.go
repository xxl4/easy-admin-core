package ctxlog

type Fields struct {
	value map[string]interface{}
}

func NewFields(key string, value interface{}) *Fields {
	f := &Fields{}
	f.Set(key, value)
	return f
}

func (e *Fields) Set(key string, value interface{}) {
	if e.value == nil {
		e.value = make(map[string]interface{})
	}
	e.value[key] = value
}

func (e *Fields) Values() map[string]interface{} {
	return e.value
}

func (e *Fields) Merge(f *Fields) {
	if len(f.value) > 0 {
		for k, v := range f.value {
			e.Set(k, v)
		}
	}
}
