package cube

type Cube struct {
	L, B, H float32
}

func (c *Cube) Area() float32 {
	return c.L * c.B * c.H
}

func (c *Cube) Permiter() float32 {
	return 4 * (c.L + c.B + c.H)
}
