package memoir

import (
	"fmt"
	"testing"
)

func TestDocument_GetKind(t *testing.T) {
	doc := NewDocument(
		DocHtmlTranslate,
		NewKeyValues(KeyValue{
			"StatusCode": 12,
			"Var":        "cccc",
			"OK":         []byte{0xFE, 0xEE},
		}),

		NewDataViewLabel(
			"Value", "请求数据", `{"aa":12}`),
	)

	j, err := doc.JSON()

	fmt.Println(string(j), err)

}
