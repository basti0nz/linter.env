package rules

// EndingBlankLine

/*
Detects if a file doesn't have a blank line at the end:

❌ Wrong
FOO=BAR
🚫↩️

✅ Correct
FOO=BAR

 */

type EndingBlankLine struct {}

func (r EndingBlankLine) CheckFile() {}

func (r EndingBlankLine) FixFile() {}