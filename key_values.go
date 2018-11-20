package memoir

type KeyValue map[string]interface{}

type KeyValues struct {
	Name  string
	Kind  ComponentKind
	Items map[string]*Value
}

func (m *KeyValues) Component() Component {
	return m
}

func (m *KeyValues) GetName() string {
	return m.Name
}

func (m *KeyValues) GetKind() ComponentKind {
	return m.Kind
}

func NewKeyValues(name string, values KeyValue) *KeyValues {

	var (
		kvs = make(map[string]*Value, 0)
	)

	if values != nil {
		for key, value := range values {
			kvs[key] = NewValue(value)
		}
	}
	return &KeyValues{
		Name:  name,
		Kind:  KindKeyValues,
		Items: kvs,
	}
}

func (m *KeyValues) Set(key string, v interface{}) {
	m.Items[key] = NewValue(v)
}
