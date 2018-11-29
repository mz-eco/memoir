package memoir

import (
	"fmt"
	"reflect"
)

type ComponentKind int

const (
	KindDocument      ComponentKind = 1
	KindNameValueList               = 2
	KindKeyValues                   = 3
	KindLabel                       = 4
	KindDataView                    = 5
	KindComponents                  = 6
)

type UIComponents []Component

func (m UIComponents) GetKind() ComponentKind {
	return KindComponents
}

type Component interface {
	GetKind() ComponentKind
}

func ToComponent(def interface{}) Component {

	switch x := def.(type) {
	case Component:
		return x
	case KeyValue:
		return NewKeyValues(x)
	case map[string]interface{}:
		return NewKeyValues(KeyValue(x))
	case UI:
		return x.UI()
	default:
		panic(
			fmt.Sprintf("could not convert type <%s> to Component", reflect.TypeOf(def)))
	}

}

func ToComponents(def ...interface{}) []Component {

	var (
		c = make([]Component, 0)
	)

	for _, v := range def {
		c = append(c, ToComponent(v))
	}

	return c

}
