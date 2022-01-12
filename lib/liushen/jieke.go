package liushen

type Jieke struct {
	index Gongwei
	name  string
	shiyi string
}

func NewJieke(index Gongwei, name string, shiyi string) *Jieke {
	return &Jieke{index: index, name: name, shiyi: shiyi}
}
