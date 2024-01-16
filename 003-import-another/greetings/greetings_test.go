package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calss greetings.hello with a name checking
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v want "", error`, msg, err)
	}
}
