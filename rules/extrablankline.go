package rules

// ExtraBlankLine
/*
Detects if a file contains more than one blank line in a row:

❌ Wrong
A=B
<Blank line>
<Blank line>
FOO=BAR

❌ Wrong
A=B
FOO=BAR
<Blank line>
<Blank line>

✅ Correct
A=B
<Blank line>
FOO=BAR
<Blank line>

✅ Correct
A=B
FOO=BAR
<Blank line>

 */
type ExtraBlankLine struct {}

func (r ExtraBlankLine) CheckFile() {}

func (r ExtraBlankLine) FixFile() {}