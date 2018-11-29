package memoir

type KeyValue map[string]interface{}

type UIKeyValue struct {
	Kind  ComponentKind
	Items map[string]*UIValue
}

func (m *UIKeyValue) Component() Component {
	return m
}

func (m *UIKeyValue) GetKind() ComponentKind {
	return m.Kind
}

func NewKeyValues(values KeyValue) *UIKeyValue {

	var (
		kvs = make(map[string]*UIValue, 0)
	)

	if values != nil {
		for key, value := range values {
			kvs[key] = NewValue(value)
		}
	}
	return &UIKeyValue{
		Kind:  KindKeyValues,
		Items: kvs,
	}
}

func (m *UIKeyValue) Set(key string, v interface{}) {
	m.Items[key] = NewValue(v)
}
