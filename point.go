package raytrace

import "math"

type Point [3]float64

func (a Point) Add(b Point) Point {
    return Point{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

func (a Point) Sub(b Point) Point {
    return Point{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func (a Point) Mult(c float64) Point {
    return Point{a[0] * c, a[1] * c, a[2] * c}
}

func (a Point) Dot(b Point) float64 {
    return a[0] * b[0] + a[1] * b[1] + a[2] * b[2]
}

func (a Point) Len() float64 {
    return math.Sqrt(a[0] * a[0] + a[1] * a[1] + a[2] * a[2])
}

func (a Point) Dist(b Point) float64 {
    dist := a.Sub(b)
    return math.Sqrt(dist[0] * dist[0] + dist[1] * dist[1] + dist[2] * dist[2])
}

func (a Point) Norm() Point {
    d := 1.0 / a.Len()
    return Point{a[0] * d, a[1] * d, a[2] * d}
}

func (a Point) Cross(b Point) Point {
    return Point{a[1]*b[2] - a[2]*b[1], a[2]*b[0] - a[0]*b[2], a[0]*b[1] - a[1]*b[0]}
}
