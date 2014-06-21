package redact

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func Example() {
	flag.Parse()                           // parse the args to get the values
	RedactFlags(flag.CommandLine, os.Args) // redact them from the runtime
}

func TestRedactFlags(t *testing.T) {
	set := flag.NewFlagSet("f", flag.PanicOnError)
	set.String("secret", "", "a secret [SENSITIVE]")
	set.String("other_secret", "", "another secret [SENSITIVE]")
	set.String("not_secret", "", "totally OK")
	set.Bool("bool", false, "what")

	args := []string{
		"-secret=whee",
		"-other_secret",
		"yay",
		"-not_secret=ok",
	}
	RedactFlags(set, args)

	expected := []string{
		"-secret=[REDACTED]",
		"-other_secret",
		"[REDACTED]",
		"-not_secret=ok",
	}

	if !reflect.DeepEqual(args, expected) {
		t.Errorf("Was %#v, but expected %#v", args, expected)
	}
}
