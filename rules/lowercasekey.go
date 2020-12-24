package rules

import (
	"errors"
	"strings"
)

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
	splt := strings.Split(strings.TrimSuffix(str, "\n"), "=")
	if len(splt) > 2 {

	}
	return str, errors.New("[LowercaseKey] This is not key=value string")
}

func (r LowercaseKey) FixString(str string)(string, error) {
	ret, err := r.CheckString(str)
	return ret, err
}