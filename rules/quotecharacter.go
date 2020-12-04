package rules

// QuoteCharacter
/*
Detects if a value contains quote characters (' / "):

❌ Wrong
FOO="BAR"

❌ Wrong
FOO='BAR'

❌ Wrong
FOO='B"AR'

✅ Correct
FOO=BAR
 */
