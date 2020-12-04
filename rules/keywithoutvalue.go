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
type KeyWithoutValue struct {}

func (r KeyWithoutValue) CheckString(str string) (string,error){
    return str, nil
}

func (r KeyWithoutValue) FixString(str string)(string, error) {
	return str, nil
}
