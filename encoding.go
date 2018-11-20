package memoir

import (
	"encoding/json"
	"encoding/xml"
)

type Encoding interface {
	Type() ValueType
	Encode(v interface{}) ([]byte, error)
	Decode(bytes []byte, v interface{}) error
}

type jsonEncoding struct {
}

func (jsonEncoding) Type() ValueType {
	return JsonValue
}

func (jsonEncoding) Encode(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "    ")
}

func (jsonEncoding) Decode(bytes []byte, v interface{}) error {
	return json.Unmarshal(bytes, v)
}

type xmlEncoding struct {
}

func (xmlEncoding) Type() ValueType {
	return XmlValue
}

func (xmlEncoding) Encode(v interface{}) ([]byte, error) {
	return xml.MarshalIndent(v, "", "    ")
}

func (xmlEncoding) Decode(bytes []byte, v interface{}) error {
	return xml.Unmarshal(bytes, v)
}
