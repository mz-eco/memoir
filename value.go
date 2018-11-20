package memoir

import (
	"fmt"
	"reflect"
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

type Value struct {
	Type  ValueType
	Value interface{}
}

type Valued interface {
	GetValue() *Value
}

type decoder interface {
	Decode(v interface{}) error
}

func parseBytes(bytes []byte, e Encoding) *Value {

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

	return &Value{
		Type:  e.Type(),
		Value: string(bytes),
	}
}

func NewValueByBytes(bytes []byte) (value *Value) {

	defer func() {

		if value == nil {
			value = &Value{
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

func NewValue(value interface{}) *Value {

	if value == nil {
		return &Value{
			Type:  NullValue,
			Value: nil,
		}
	}

	switch v := value.(type) {
	case int8, int16, int32, int64, int:
		return &Value{
			Type:  IntValue,
			Value: reflect.ValueOf(value).Int(),
		}
	case uint, uint8, uint16, uint32, uint64:
		return &Value{
			Type:  UIntValue,
			Value: reflect.ValueOf(value).Uint(),
		}
	case string:
		x := NewValueByBytes([]byte(v))

		if x.Type == BytesValue {
			return &Value{
				Type:  StringValue,
				Value: v,
			}
		}

		return x
	case []byte:
		return NewValueByBytes(v)
	case bool:
		return &Value{
			Type:  BooleanValue,
			Value: v,
		}
	case time.Time:
		return &Value{
			Type:  TimeValue,
			Value: v,
		}
	case time.Duration:
		return &Value{
			Type:  DurationValue,
			Value: v,
		}
	case *Value:
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

func (m *Value) String() string {
	return fmt.Sprintf("%s", m.Value)
}
