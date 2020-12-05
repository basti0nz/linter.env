package rules

type StringRuler interface {
	CheckString(str string)(bool, error, string)
	FixString(str string)(bool, error, string)

}
type FileRuler interface {
	CheckFile()
	FixFile()
}