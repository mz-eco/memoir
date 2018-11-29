package memoir

type NameValue struct {
	Value *UIValue
	Name  string
}

func NewNamedValue(name string, value interface{}) *NameValue {
	return &NameValue{
		Value: NewValue(value),
		Name:  name,
	}
}

type NameValueList struct {
	Kind  ComponentKind
	Name  string
	Items []*NameValue
}

func (m *NameValueList) Component() Component {
	return m
}

func (m *NameValueList) GetName() string {
	return m.Name
}

func (m *NameValueList) GetKind() ComponentKind {
	return m.Kind
}

func NewNameValueList(name string) *NameValueList {
	return &NameValueList{
		Kind:  KindNameValueList,
		Name:  name,
		Items: make([]*NameValue, 0),
	}
}

func (m *NameValueList) Append(name string, value interface{}) {
	m.Items = append(m.Items, NewNamedValue(name, value))
}
