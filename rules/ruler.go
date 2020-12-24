package rules

import "os"

type StringRuler interface {
	CheckString(str string)(bool, error, string)
	FixString(str string)(bool, error, string)

}
type FileRuler interface {
	CheckFile(file *os.File) error
	FixFile(file *os.File) error
}