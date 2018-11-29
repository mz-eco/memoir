package memoir

import (
	"bytes"
	"io"
	"strings"
)

type Writer struct {
	level int
	w     io.Writer
	head  bool
}

func (m Writer) Next() Writer {

	return Writer{
		level: m.level + 1,
		w:     m.w,
		head:  true,
	}
}

func NewWriter(w io.Writer) Writer {
	return Writer{
		w:     w,
		level: 0,
		head:  true,
	}
}

func (m Writer) Write(p []byte) (n int, err error) {

	var (
		ws int
	)

	if m.head {
		m.head = false
		ws, err = m.w.Write([]byte(strings.Repeat(" ", 4*m.level)))

		if err != nil {
			return ws, err
		}
	}

	for len(p) > 0 {
		index := bytes.IndexByte(p, '\n')

		if index < 0 {
			ws, err = m.w.Write(p)
			return n + ws, err
		} else {
			ws, err = m.w.Write(p[0 : index+1])

			if err != nil {
				break
			}
			n += ws
			p = p[index+1:]

			if len(p) > 0 {
				ws, err = m.w.Write([]byte(strings.Repeat(" ", 4*m.level)))

				if err != nil {
					break
				}

				n += ws
			}

		}

	}

	n += ws
	return
}
