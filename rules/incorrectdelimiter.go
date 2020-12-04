package rules

// IncorrectDelimiter
/*
Detects if a key does not use an underscore to separate words:

❌ Wrong
FOO-BAR=FOOBAR

✅ Correct
FOO_BAR=FOOBAR
 */

type IncorrectDelimiter struct {}

func (r IncorrectDelimiter) CheckString(str string) (string,error){
    return str, nil
}

func (r IncorrectDelimiter) FixString(str string)(string, error) {
	return str, nil
}
