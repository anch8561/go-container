package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	// call greetings.Hello with a name
	name := "Andrew"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name) // go uses pass-by-value
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("%v") = %q, %v, want match for %#q, nil`,
			name, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
