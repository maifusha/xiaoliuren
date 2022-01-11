package entity

type Jieke struct {
	name  string
	shiyi string
}

func NewJieke(name string, shiyi string) *Jieke {
	return &Jieke{name: name, shiyi: shiyi}
}
