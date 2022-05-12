package util

import "time"

type Now struct {
	t   time.Time
	fmt string
}

func NewNow() *Now {
	return &Now{
		t:   time.Now(),
		fmt: "2006-01-02 15:04:05",
	}
}

func (n *Now) SetFmt(f string) *Now {
	n.fmt = f
	return n
}

func (n *Now) String() string {
	return n.t.Format(n.fmt)
}

func (n *Now) Stamp() int64 {
	return n.t.Unix()
}
