package memoir

import "encoding/json"

type DocumentType int

const (
	DocHtmlTranslate DocumentType = 0
)

func (m DocumentType) String() string {

	switch m {
	case DocHtmlTranslate:
		return "DocHtmlTranslate"
	}

	return ""
}

type Document struct {
	Kind       ComponentKind
	Name       string
	Type       DocumentType
	Components []Component
}

func (m *Document) Component() Component {
	return m
}

func (m *Document) GetName() string {
	return m.Name
}

func (m *Document) GetKind() ComponentKind {
	return m.Kind
}

func (m *Document) JSON() ([]byte, error) {
	return json.MarshalIndent(m, "", "    ")
}

func (m *Document) Add(components ...Component) *Document {

	m.Components = append(m.Components, components...)

	return m
}

func NewDocument(doc DocumentType, name string, uiList ...UI) *Document {

	return &Document{
		Kind:       KindDocument,
		Name:       name,
		Type:       doc,
		Components: UIToComponent(uiList...),
	}
}
