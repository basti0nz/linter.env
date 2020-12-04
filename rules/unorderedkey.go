package rules

/* UnorderedKey
Detects if a key is not alphabetically ordered:

❌ Wrong
FOO=BAR
BAR=FOO

✅ Correct
BAR=FOO
FOO=BAR

❌ Wrong
FOO=BAR
BAR=FOO

✅ Correct
FOO=BAR

BAR=FOO

✅ Correct
FOO=BAR
# dotenv-linter:off LowercaseKey
bar=FOO
 */