// Package redact provides functionality to redact potentially sensitive flags,
// to prevent their contents from being exposed via expvars or debugging
// endpoints.
package redact

import (
	"flag"
	"fmt"
	"strings"
)

// RedactFlags iterates through the given FlagSet and arguments slice, and for
// each non-bool flag which has the string "[SENSITIVE]" in its usage, redacts
// the value of the flag from the given arguments.
func RedactFlags(set *flag.FlagSet, args []string) {
	set.VisitAll(func(f *flag.Flag) {
		if _, ok := f.Value.(boolFlag); ok {
			return
		}

		if strings.Contains(f.Usage, sensitive) {
			short := fmt.Sprintf("-%s", f.Name)
			long := fmt.Sprintf("-%s=", f.Name)

			for i, v := range args {
				if strings.HasPrefix(v, long) {
					args[i] = long + redacted
				} else if strings.HasPrefix(v, short) && i < len(args)-1 {
					args[i+1] = redacted
				}
			}
		}
	})
}

// oh hey there guts of flag package how are yooouuuu
type boolFlag interface {
	flag.Value
	IsBoolFlag() bool
}

const (
	sensitive = "[SENSITIVE]"
	redacted  = "[REDACTED]"
)
