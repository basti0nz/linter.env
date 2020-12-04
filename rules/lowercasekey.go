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

type LowercaseKey struct {}

func (r LowercaseKey) CheckString(str string) (string,error){
	return str, nil
}

func (r LowercaseKey) FixString(str string)(string, error) {
	return str, nil
}