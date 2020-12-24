package rules

import (
	"regexp"
)

// LeadingCharacter
/*
Detects if a line starts with an unallowed character (characters from A to Z and _ (underscore) are allowed):

❌ Wrong
 FOO=BAR

❌ Wrong
.FOO=BAR

❌ Wrong
*FOO=BAR

❌ Wrong
1FOO=BAR

✅ Correct
FOO=BAR

✅ Correct
_FOO=BAR
*/



type LeadingCharacter struct {
	regexpStr *regexp.Regexp
	ErrorCheckMessage string
}

func NewLeadingCharacter() *LeadingCharacter {
	r := LeadingCharacter{}
	r.regexpStr = regexp.MustCompile(`^[\d\.\s\W]+`)
	r.ErrorCheckMessage = "[LeadingCharacter] The line starts with unallowed character"
	return &r
}

func (r *LeadingCharacter) CheckString(str string) (bool, string) {
    return r.regexpStr.MatchString(str), r.ErrorCheckMessage
}

func (r *LeadingCharacter) FixString(str string) string {
	str = r.regexpStr.ReplaceAllString(str, "")
	return str
}
