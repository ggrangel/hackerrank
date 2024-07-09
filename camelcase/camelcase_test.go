package camelcase

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCamelCase(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"empty string":   {input: "", want: 0},
		"one word":       {input: "one", want: 1},
		"two words":      {input: "oneTwo", want: 2},
		"only uppercase": {input: "oYSAD", want: 5},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := camelCase(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
