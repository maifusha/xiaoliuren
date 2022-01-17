package liushen

type Jieke struct {
	Index   Gongwei `json:"index"`
	Name    string  `json:"name"`
	Jixiong string  `json:"jixiong"`
	Shiyi   string  `json:"shiyi"`
}

func NewJieke(index Gongwei, name string, jixiong string, shiyi string) *Jieke {
	return &Jieke{Index: index, Name: name, Jixiong: jixiong, Shiyi: shiyi}
}
