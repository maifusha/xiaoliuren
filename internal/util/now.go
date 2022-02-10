package util

import "time"

type Now struct {
	t time.Time
	f string
}

func NewNow() *Now {
	return &Now{
		t: time.Now(),
		f: "2006-01-02 15:04:05",
	}
}

func (n *Now) Fmt(f string) *Now {
	n.f = f
	return n
}

func (n *Now) String() string {
	return n.t.Format(n.f)
}

func (n *Now) Stamp() int64 {
	return n.t.Unix()
}
