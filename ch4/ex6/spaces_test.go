
package spaces

import (
    "testing"
    "bytes"
)

// TestReverseUTF8 tests the reverseUTF8 function to ensure it correctly reverses UTF-8 encoded strings.
func TestSpaces(t *testing.T) {
    tests := []struct {
        input []byte
        want []byte
    }{
        {[]byte(" Go语\t\n言"), []byte(" Go语  言")},
        {[]byte("Hello, 世界"), []byte("Hello, 世界")},
        // Add more test cases here.  
    }

    for _, tt := range tests {
        spaces(tt.input)
        if !bytes.Equal(tt.input, tt.want) {
            t.Errorf("reverseUTF8(%q) = %q, want %q", tt.input, tt.input, tt.want)
        }
    }
}
