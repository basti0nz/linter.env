package rules

/* UnorderedKey

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