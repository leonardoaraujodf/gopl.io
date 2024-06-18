package reverse

import (
    "testing"
    "bytes"
)

// TestReverseUTF8 tests the reverseUTF8 function to ensure it correctly reverses UTF-8 encoded strings.
func TestReverseUTF8(t *testing.T) {
    tests := []struct {
        input []byte
        want []byte
    }{
        {[]byte("Go语言"), []byte("言语oG")},
        {[]byte("Hello, 世界"), []byte("界世 ,olleH")},
        // Add more test cases here.  
    }

    for _, tt := range tests {
        reverseUTF8(tt.input)
        if !bytes.Equal(tt.input, tt.want) {
            t.Errorf("reverseUTF8(%q) = %q, want %q", tt.input, tt.input, tt.want)
        }
    }
}
