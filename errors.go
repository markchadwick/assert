package assert

import (
	"bytes"
	"fmt"
	"reflect"
)

type Error struct {
	preambles []string
	values    []interface{}
}

func NewError() *Error {
	return &Error{
		preambles: make([]string, 0),
		values:    make([]interface{}, 0),
	}
}

func (e *Error) Add(preamble string, value interface{}) *Error {
	if value != nil {
		preamble = fmt.Sprintf("%s %T: ", preamble, value)
	} else {
		preamble = fmt.Sprintf("%s: ", preamble)
	}

	e.preambles = append(e.preambles, preamble)
	e.values = append(e.values, value)

	return e
}

func (e *Error) Error() string {
	longest := 0
	for _, pre := range e.preambles {
		if len(pre) > longest {
			longest = len(pre)
		}
	}
	buf := new(bytes.Buffer)
	lenFmt := fmt.Sprintf("%%%ds", longest)
	num := len(e.preambles)

	for i, pre := range e.preambles {
		fmt.Fprintf(buf, lenFmt, pre)
		fmt.Fprintf(buf, e.valueStr(e.values[i]))
		if i < num-1 {
			fmt.Fprint(buf, "\n")
		}
	}
	return buf.String()
}

func (e *Error) valueStr(i interface{}) string {
	switch t := i.(type) {
	case string:
		return fmt.Sprintf(`"%s"`, t)
	case reflect.Value:
		return e.valueStr(t.Interface())
	}
	return fmt.Sprintf("%v", i)
}

func errNotEqual(exp, act interface{}) error {
	return NewError().
		Add("expected", exp).
		Add("received", act)
}

func errEqual(exp, act interface{}) error {
	return NewError().
		Add("expected not", exp).
		Add("received", act)
}
