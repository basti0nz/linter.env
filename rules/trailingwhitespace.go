package rules

// TrailingWhitespace
/*
Detects if a line has a trailing whitespace:

❌ Wrong
FOO=BAR<Trailing whitespace>

✅ Correct
FOO=BAR


import (
	"fmt"
	"strings"
	"regexp"
)


	re := regexp.MustCompile(`\r?\n`)
	space := regexp.MustCompile(`\s+`)
	x = re.ReplaceAllString(x, "")
	x = space.ReplaceAllString(x, "")
 */
