package caesar

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCaesarCipher(t *testing.T) {
	tests := map[string]struct {
		plaintext string
		shift     int
		want      string
	}{
		"negative shift": {plaintext: "abc", shift: -1, want: "zab"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := caesarCipher(tc.plaintext, tc.shift)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
