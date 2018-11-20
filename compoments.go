package memoir

type ComponentKind int

const (
	KindDocument      ComponentKind = 1
	KindNameValueList               = 2
	KindKeyValues                   = 3
	KindLabel                       = 4
	KindDataView                    = 5
)

type Component interface {
	GetName() string
	GetKind() ComponentKind
}
