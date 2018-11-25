package main

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

// ComputeId returns the first 12 characters from a SHA1 hash of the combined input string(s)
func ComputeId(input ...string) string {
	inputString := strings.Join(input, "")
	sum := fmt.Sprintf("%x", sha1.Sum([]byte(inputString)))
	return sum[0:12]
}
