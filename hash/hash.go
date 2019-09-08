package hash

import (
	"crypto/sha1"
	"fmt"
	"io"
)

// Email function takes an email and returns a hash
func Email(email string) []byte {
	// 1) create new hash
	h := sha1.New()
	// 2) use io package to write value to it (added to byte slice)
	io.WriteString(h, email)
	fmt.Printf("% x \n", h.Sum(nil))
	return h.Sum(nil)
}
