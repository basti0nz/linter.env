package rules

// KeyWithoutValue
/*
Detects if a line has a key without a value:

❌ Wrong
FOO

✅ Correct
FOO=

✅ Correct
FOO=BAR
 */
