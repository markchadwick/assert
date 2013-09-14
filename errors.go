package assert

import (
	"bytes"
	"errors"
	"fmt"
)

func errNotEqual(exp, act interface{}) error {
	expPre := preamble("expected", exp)
	actPre := preamble("received", act)

	length := len(expPre)
	if len(actPre) > length {
		length = len(actPre)
	}
	lenFmt := fmt.Sprintf("%%-%ds", length)

	b := new(bytes.Buffer)
	fmt.Fprintf(b, lenFmt, expPre)
	fmt.Fprintf(b, "%v\n", exp)
	fmt.Fprintf(b, lenFmt, actPre)
	fmt.Fprintf(b, "%v", act)
	return errors.New(b.String())
}

func errEqual(exp, act interface{}) error {
	expPre := preamble("expected not", exp)
	actPre := preamble("received", act)

	length := len(expPre)
	if len(actPre) > length {
		length = len(actPre)
	}
	lenFmt := fmt.Sprintf("%%-%ds", length)

	b := new(bytes.Buffer)
	fmt.Fprintf(b, lenFmt, expPre)
	fmt.Fprintf(b, "%v\n", exp)
	fmt.Fprintf(b, lenFmt, actPre)
	fmt.Fprintf(b, "%v", act)
	return errors.New(b.String())
}

func preamble(msg string, v interface{}) string {
	if v == nil {
		return fmt.Sprintf("%s: ", msg)
	}
	return fmt.Sprintf("%s %T: ", msg, v)
}
