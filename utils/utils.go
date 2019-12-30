// Package utils includes utilities shared among backends.
package utils

import (
	"regexp"
)

// This regex is probably ugly but it works so far.

// CodeBlockRegex is the regexp used for matching codeblocks.
var CodeBlockRegex = regexp.MustCompile("^```(?:\\S*)\n?((?:.|\\s)*)\n?```$")

// StripCodeBlocks strips leading and trailing codeblock characters in the input source.
func StripCodeBlock(src string) string {
	if len(src) == 0 {
		return ""
	}

	match := CodeBlockRegex.FindStringSubmatch(src)

	// If no codeblocks matched just return the input as is.
	if match == nil {
		return src
	}

	return match[1]
}
