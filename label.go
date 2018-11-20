package memoir

type Label struct {
	Kind       ComponentKind
	Name       string
	Title      string
	Components []Component
}

func (m *Label) Component() Component {
	return m
}

func (m *Label) GetName() string {
	return m.Name
}

func (m *Label) GetKind() ComponentKind {
	return m.Kind
}

func NewLabel(name string, uiList ...UI) *Label {

	return &Label{
		Kind:       KindLabel,
		Name:       name,
		Components: UIToComponent(uiList...),
	}
}
