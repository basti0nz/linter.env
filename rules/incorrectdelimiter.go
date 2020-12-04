package rules

// IncorrectDelimiter
/*
Detects if a key does not use an underscore to separate words:

❌ Wrong
FOO-BAR=FOOBAR

✅ Correct
FOO_BAR=FOOBAR
 */
