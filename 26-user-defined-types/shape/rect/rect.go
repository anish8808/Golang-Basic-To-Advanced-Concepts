package rect

type Rect struct {
	L, B float32
	A, P float32
}

func (r *Rect) Area() float32 { // when the recivers are pointer (just the reciver will be pointer now it will work as call by refrence)
	r.A = r.L * r.B
	return r.A
}

func (r *Rect) Permiter() float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}
