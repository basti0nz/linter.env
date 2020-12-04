package rules

// SpaceCharacter
/*
Detects lines with a whitespace around equal sign character =:

❌ Wrong
FOO =BAR

❌ Wrong
FOO= BAR

❌ Wrong
FOO = BAR

✅ Correct
FOO=BAR

 */
