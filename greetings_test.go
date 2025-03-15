package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Mia"
	want := regexp.MustCompile(`\b` + name + `\b`)
	got, err := Hello("Mia")
	if !want.MatchString(got) || err != nil {
		t.Fatalf(`Hello("Mia") = %q, %v, want match for %#q, nil`, got, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	got, err := Hello("")
	if got != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want match "", error`, got, err)
	}
}
