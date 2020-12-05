package rules

// IncorrectDelimiter
/*
Detects if a key does not use an underscore to separate words:

❌ Wrong
FOO-BAR=FOOBAR

✅ Correct
FOO_BAR=FOOBAR
 */

const DELIMITER = "_"

type IncorrectDelimiter struct {}

func (r IncorrectDelimiter) CheckString(str string) (bool, error, string){
    return false, nil, str
}

func (r IncorrectDelimiter) FixString(str string)(bool, error, string) {
	return false, nil, str
}
