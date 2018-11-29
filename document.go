package memoir

import "encoding/json"

type DocumentType int

const (
	DocHttpTranslate DocumentType = 0
)

func (m DocumentType) String() string {

	switch m {
	case DocHttpTranslate:
		return "DocHttpTranslate"
	}

	return ""
}

type UIDocument struct {
	Kind       ComponentKind
	Name       string
	Type       DocumentType
	Components []Component
}

func (m *UIDocument) Component() Component {
	return m
}

func (m *UIDocument) GetName() string {
	return m.Name
}

func (m *UIDocument) GetKind() ComponentKind {
	return m.Kind
}

func (m *UIDocument) JSON() ([]byte, error) {
	return json.MarshalIndent(m, "", "    ")
}

func (m *UIDocument) Add(components ...Component) *UIDocument {

	m.Components = append(m.Components, components...)

	return m
}

func NewDocument(doc DocumentType, name string, v ...interface{}) *UIDocument {

	return &UIDocument{
		Kind:       KindDocument,
		Name:       name,
		Type:       doc,
		Components: ToComponents(v...),
	}
}
