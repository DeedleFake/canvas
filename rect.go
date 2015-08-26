package canvas

// Rectangle represents a rectangle. It is patterned after
// image.Rectangle, but is float64 based.
type Rectangle struct {
	Min, Max Point
}

var ZR = Rectangle{}

func Rect(x1, y1, x2, y2 float64) Rectangle {
	return Rectangle{Point{x1, y1}, Point{x2, y2}}.Canon()
}

func (r Rectangle) Canon() Rectangle {
	if r.Max.X < r.Min.X {
		r.Min.X, r.Max.X = r.Max.X, r.Min.X
	}

	if r.Max.Y < r.Min.Y {
		r.Min.Y, r.Max.Y = r.Max.Y, r.Min.Y
	}

	return r
}

func (r Rectangle) Dx() float64 {
	return r.Max.X - r.Min.X
}

func (r Rectangle) Dy() float64 {
	return r.Max.Y - r.Min.Y
}

// Point represents a point on the cartesian plane. It is patterned
// after image.Point, but is float64 based.
type Point struct {
	X, Y float64
}

var ZP = Point{}

func Pt(x, y float64) Point {
	return Point{x, y}
}
