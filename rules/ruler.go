package rules

type StringRuler interface {
	CheckString(str string)(string,error)
	FixString(str string)(string, error)

}
type FileRuler interface {
	CheckFile()
	FixFile()
}