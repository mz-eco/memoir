package memoir

type UILabel struct {
	Kind       ComponentKind
	Name       string
	Title      string
	Components []Component
}

func (m *UILabel) Component() Component {
	return m
}

func (m *UILabel) GetName() string {
	return m.Name
}

func (m *UILabel) GetKind() ComponentKind {
	return m.Kind
}

func Label(name string, v ...interface{}) *UILabel {

	return &UILabel{
		Kind:       KindLabel,
		Name:       name,
		Components: ToComponents(v...),
	}
}
