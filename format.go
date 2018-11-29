package memoir

import (
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/gosuri/uitable"
)

type Formatter struct {
	w Writer
}

func (m Formatter) Next() Formatter {
	return Formatter{
		w: m.w.Next(),
	}
}

func (m Formatter) Write(p []byte) (n int, err error) {
	return m.w.Write(p)
}

func (m Formatter) Line(format string, v ...interface{}) {

	fmt.Fprintln(
		m.w,
		fmt.Sprintf(format, v...),
	)
}

func format(w Formatter, v interface{}) {

	switch x := v.(type) {
	case UI:
		format(w, x.UI())
	case *UILabel:
		w.Line("%s:", x.Name)
		format(w.Next(), x.Components)
	case []Component:
		for _, c := range x {
			format(w, c)
		}
	case UIComponents:
		for _, c := range x {
			format(w, c)
		}
	case *UIKeyValue:
		tbl := uitable.New()

		for k, v := range x.Items {
			tbl.AddRow(k, ":", v.Simple())
		}

		w.Line(tbl.String())
	case *UIDataView:
		w.Line("%s", x.Value.View())
	case *UIDocument:
		w.Line("Document [%s] (%s)", x.Name, x.Type)
		format(w.Next(), x.Components)
	default:
		panic(fmt.Sprintf("%s not support format", reflect.TypeOf(v)))
	}

}

func Format(w io.Writer, v interface{}) {

	var (
		formatter = Formatter{
			w: NewWriter(w),
		}
	)

	format(formatter, v)

}

func Print(v interface{}) {
	Format(os.Stdout, v)
}
