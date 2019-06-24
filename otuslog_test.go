package otuslog

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

var testCases = []OtusEvent{
	&HwSubmitted{1, "code?", "my homework"},
	&HwAccepted{2, 4},
	&HwSubmitted{3, "code?", "link 1"},
	&HwAccepted{4, 5},
	&HwMy{563, "untitled"},
	&HwSubmitted{5, "code?", "homework"},
	&HwAccepted{6, 7},
	&HwSubmitted{7, "code?", "link 2"},
	&HwAccepted{8, 3},
}

func TestInterface(t *testing.T) {

	output := bytes.NewBuffer([]byte{})

	data := bytes.NewBuffer([]byte{})
	fmt.Fprintf(data, "%s accepted 1 5\n", time.Now().Format("2006-01-02"))

	LogOtusEvent(&HwAccepted{1, 5}, output)

	if data.String() != output.String() {
		t.Errorf("Accepted Event error:\n\tEXPECTED %v\n\tGOT %v", data.String(), output.String())
	}

	output.Reset()
	data.Reset()
	fmt.Fprintf(data, "%s submitted 2 \"homework\"\n", time.Now().Format("2006-01-02"))

	LogOtusEvent(&HwSubmitted{2, "", "homework"}, output)

	if data.String() != output.String() {
		t.Errorf("Submitted Event error:\n\tEXPECTED %v\n\tGOT %v", data.String(), output.String())
	}

}

func BenchmarkLogOtusEvent(b *testing.B) {

	buf := bytes.NewBuffer([]byte{})

	for _, testCase := range testCases {
		LogOtusEvent(testCase, buf)
	}

}

func BenchmarkLogOtusEvent2(b *testing.B) {

	buf := bytes.NewBuffer([]byte{})

	for _, testCase := range testCases {
		LogOtusEvent2(testCase, buf)
	}

}
