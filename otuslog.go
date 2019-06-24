package otuslog

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

type HwAccepted struct {
	Id    int
	Grade int
}

type HwSubmitted struct {
	Id      int
	Code    string
	Comment string
}

type HwMy struct {
	Id      int
	Comment string
}

func (e *HwSubmitted) Log() string {
	return fmt.Sprintf("%d \"%s\"", e.Id, e.Comment)
}

func (e *HwAccepted) Log() string {
	return fmt.Sprintf("%d %d", e.Id, e.Grade)
}

func (e *HwMy) Log() string {
	return fmt.Sprintf("%d \"%s\"", e.Id, e.Comment)
}

type OtusEvent interface {
	Log() string
}

// 2019-01-01 submitted 3456 "please take a look at my homework"
// 2019-01-01 accepted 3456 4

func LogOtusEvent(e OtusEvent, w io.Writer) {

	logstr := bytes.NewBufferString(time.Now().Format("2006-01-02 "))

	switch e.(type) {
	case *HwAccepted:
		logstr.WriteString("accepted ")
	case *HwSubmitted:
		logstr.WriteString("submitted ")
	default:
		logstr.WriteString("unknown event ")
	}

	logstr.WriteString(e.Log())
	logstr.WriteString("\n")

	w.Write(logstr.Bytes())
}

func LogOtusEvent2(e OtusEvent, w io.Writer) {

	var tp string

	switch e.(type) {
	case *HwAccepted:
		tp = "accepted"
	case *HwSubmitted:
		tp = "submitted"
	default:
		tp = "unknown event"
	}

	fmt.Fprintf(w, "%s %s %s\n", time.Now().Format("2006-01-02"), tp, e.Log())
}
