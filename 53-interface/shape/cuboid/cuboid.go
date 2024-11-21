package cuboid

//var Global string

type Cuboid struct { // this is unrestricted since starts with Uppercase
	L, B, H float32 // exported or unrestricted fields bcz starts with Uppercase
	a, p    float32 // unexported or restricted fields bcz of lowercase
}

func New(l, b, h float32) *Cuboid {
	return &Cuboid{L: l, B: b, H: h}
}

func (r *Cuboid) Area() float32 {
	r.a = r.L * r.B * r.H
	return r.a
}

func (r *Cuboid) Perimeter() float32 {
	r.p = 4 * (r.L + r.B + r.H)
	return r.p
}

func (r *Cuboid) What() string {
	return "Cuboid"
}
