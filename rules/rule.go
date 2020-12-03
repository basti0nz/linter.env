package rules

type Rule interface {
	Check(str string)(string,error)
	Fix(str string)(string, error)
}