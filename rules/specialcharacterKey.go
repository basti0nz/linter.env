package rules

import (
	"regexp"
)

// specialCharacter
/*
Detects if a line  with an unallowed special characters (characters from a-z, A-Z, 0-9 and _  = #  are allowed):

❌ Wrong
.?%^&**FOO=BAR

❌ Wrong
*FOO=BAR

❌ Wrong
1FOO=BAR

✅ Correct
FOO=BAR

✅ Correct
_FOO=BAR
*/

type SpecialCharacter struct {
	regexpStr *regexp.Regexp
	ErrorCheckMessage string
}



func NewSpecialCharacter() *SpecialCharacter {
	r := SpecialCharacter{}
	r.regexpStr = regexp.MustCompile(`[^\sa-zA-Z0-9_#=]+`)
	r.ErrorCheckMessage = "[SpecialCharacter] Special characters in the line is not allowed"
	return &r
}

func (r *SpecialCharacter) CheckString(str string) (bool, string) {
    return r.regexpStr.MatchString(str), r.ErrorCheckMessage
}

func (r *SpecialCharacter) FixString(str string) string {
	str = r.regexpStr.ReplaceAllString(str, "")
	return str
}
