package spaces

import (
    "unicode"
    "unicode/utf8"
)

func spaces(s []byte) []byte {
    for i := 0; i < len(s); {
        r, size := utf8.DecodeRune(s[i:])
        if unicode.IsSpace(r) {
            copy(s[i+1:],s[i+size:])
            s[i] = ' ';
            // Shrink s removing the rune size but adding 1 for the space in ASCII
            s = s[:len(s)-size+1]
            i++
        } else {
            i += size
        }
    }

    return s
}
