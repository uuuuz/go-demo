package demo

type Struct001 struct {
	name chain
	age  int
}

type chain struct {
	reporter string
	flag     bool
}

func NewStruct001() *Struct001 {
	return &Struct001{
		name: chain{
			reporter: "marisa",
			flag:     true,
		},
		age: 12,
	}
}
