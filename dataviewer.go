package memoir

type UIDataView struct {
	Kind  ComponentKind
	Name  string
	Value *UIValue
}

func (m *UIDataView) Component() Component {
	return m
}

func (m *UIDataView) GetName() string {
	return m.Name
}

func (m *UIDataView) GetKind() ComponentKind {
	return m.Kind
}

func DataView(v interface{}) *UIDataView {
	return &UIDataView{
		Kind:  KindDataView,
		Value: NewValue(v),
	}
}
