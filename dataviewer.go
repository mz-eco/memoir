package memoir

type DataView struct {
	Kind  ComponentKind
	Name  string
	Value *Value
}

func (m *DataView) Component() Component {
	return m
}

func (m *DataView) GetName() string {
	return m.Name
}

func (m *DataView) GetKind() ComponentKind {
	return m.Kind
}

func NewDataView(name string, v interface{}) *DataView {

	return &DataView{
		Name:  name,
		Kind:  KindDataView,
		Value: NewValue(v),
	}
}
