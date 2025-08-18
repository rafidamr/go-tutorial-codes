package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello(name)
	if !want.MatchString(message) || err != nil {
		t.Errorf(`Hello("Gladys") = %q, %v, wants %#q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	message, err := Hello(name)
	if message != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, wants "", error`, message, err)
	}
}
