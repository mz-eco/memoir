package memoir

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type ValueType int

const (
	NullValue     ValueType = 0
	IntValue                = 1
	UIntValue               = 2
	StringValue             = 3
	BytesValue              = 4
	TimeValue               = 5
	DurationValue           = 6
	BooleanValue            = 7
	JsonValue               = 8
	XmlValue                = 9
)

type UIValue struct {
	Type  ValueType
	Value interface{}
}

type Valued interface {
	GetValue() *UIValue
}

type decoder interface {
	Decode(v interface{}) error
}

func parseBytes(bytes []byte, e Encoding) *UIValue {

	var (
		middle = make(map[string]interface{})
	)

	err := e.Decode(bytes, &middle)

	if err != nil {
		return nil
	}

	bytes, err = e.Encode(middle)

	if err != nil {
		return nil
	}

	return &UIValue{
		Type:  e.Type(),
		Value: string(bytes),
	}
}

func NewValueByBytes(bytes []byte) (value *UIValue) {

	defer func() {

		if value == nil {
			value = &UIValue{
				Type:  BytesValue,
				Value: bytes,
			}
		}
	}()

	if len(bytes) == 0 {
		return
	}

	ch := bytes[0]

	switch ch {
	case '{', '[':
		value = parseBytes(bytes, jsonEncoding{})
	case '<':
		value = parseBytes(bytes, xmlEncoding{})
	}

	return
}

func NewValue(value interface{}) *UIValue {

	if value == nil {
		return &UIValue{
			Type:  NullValue,
			Value: nil,
		}
	}

	switch v := value.(type) {
	case int8, int16, int32, int64, int:
		return &UIValue{
			Type:  IntValue,
			Value: reflect.ValueOf(value).Int(),
		}
	case uint, uint8, uint16, uint32, uint64:
		return &UIValue{
			Type:  UIntValue,
			Value: reflect.ValueOf(value).Uint(),
		}
	case string:
		x := NewValueByBytes([]byte(v))

		if x.Type == BytesValue {
			return &UIValue{
				Type:  StringValue,
				Value: v,
			}
		}

		return x
	case []byte:
		return NewValueByBytes(v)
	case bool:
		return &UIValue{
			Type:  BooleanValue,
			Value: v,
		}
	case time.Time:
		return &UIValue{
			Type:  TimeValue,
			Value: v,
		}
	case time.Duration:
		return &UIValue{
			Type:  DurationValue,
			Value: v,
		}
	case *UIValue:
		return v
	default:

		valued, ok := value.(Valued)

		if ok {
			return valued.GetValue()
		}

		stringer, ok := value.(fmt.Stringer)

		if ok {
			return NewValue(stringer.String())
		}
	}

	panic(
		fmt.Sprintf("value type %s not support.", reflect.TypeOf(value)))
}

func (m *UIValue) String() string {
	return fmt.Sprintf("%s", m.Value)
}

func (m *UIValue) View() string {

	switch m.Type {
	case BytesValue:
		return hex.Dump(m.Value.([]byte))
	case StringValue, JsonValue, XmlValue:
		return m.Value.(string)
	case TimeValue:
		p := m.Value.(time.Time)
		return p.Format("2006-01-02 03:04:05")
	case NullValue:
		return "<nil>"
	case IntValue, UIntValue:
		return fmt.Sprintf("%d", m.Value)
	default:
		return fmt.Sprintf("%s", m.Value)
	}
}

func (m *UIValue) Simple() interface{} {

	switch m.Type {
	case BytesValue:
		p := m.Value.([]byte)

		if len(p) > 10 {
			return append(p[:10])
		}
	case StringValue:

		p := m.Value.(string)

		index := strings.Index(p, "\n")

		if index < 0 {
			return p
		}

		if len(p) > 60 {
			return string(p[0:index]) + "..."
		}
	case JsonValue:
		return "<JSON>"
	case XmlValue:
		return "<XML>"
	case TimeValue:
		p := m.Value.(time.Time)

		return p.Format("2006-01-02 03:04:05")
	}

	return m.Value
}
