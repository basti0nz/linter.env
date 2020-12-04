package rules

type Rule interface {
	CheckString(str string)(string,error)
	FixString(str string)(string, error)
	CheckFile()
	FixFile()
}