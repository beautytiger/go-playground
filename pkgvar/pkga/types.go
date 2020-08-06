package pkga

var V1 string = "in pkg 1"

type MyStruct struct {
	Name string
}

func NewStruct() MyStruct {
	return MyStruct{V1}
}
