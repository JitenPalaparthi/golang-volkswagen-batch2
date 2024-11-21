package rect

//var Global string

type Rect struct { // this is unrestricted since starts with Uppercase
	L, B float32 // exported or unrestricted fields bcz starts with Uppercase
	a, p float32 // unexported or restricted fields bcz of lowercase
}

func New(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

func (r *Rect) Area() float32 {
	r.a = r.L * r.B
	return r.a
}

func (r *Rect) Perimeter() float32 {
	r.p = 2 * (r.L + r.B)
	return r.p
}

func (r *Rect) What() string {
	return "Rect"
}
