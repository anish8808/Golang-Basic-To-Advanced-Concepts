package square

type Square struct {
	S float32
}

func (s1 *Square) Area() float32 {
	return s1.S * s1.S
}
func (s *Square) Permiter() float32 {
	return 2 * (s.S)
}
