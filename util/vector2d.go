package util

type Vector2D struct {
	X, Y float64
}

func (v Vector2D) Add(right Vector2D) Vector2D {
	return Vector2D{X: v.X + right.X, Y: v.Y + right.Y}
}

func NewVector(x, y float64) Vector2D {
	return Vector2D{x, y}
}

// dot
// length
// cross
