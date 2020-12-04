package rules

// LowercaseKey
/*
Detects if a key has lowercase characters:

❌ Wrong
FOo_BAR=FOOBAR

❌ Wrong
foo_bar=FOOBAR

✅ Correct
FOO_BAR=FOOBAR

 */
