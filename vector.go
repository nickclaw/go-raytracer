package raytrace

import "math"

type Vector [3]float64

func (a Vector) Add(b Vector) Vector {
    return Vector{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

func (a Vector) Sub(b Vector) Vector {
    return Vector{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func (a Vector) Mult(c float64) Vector {
    return Vector{a[0] * c, a[1] * c, a[2] * c}
}

func (a Vector) Dot(b Vector) float64 {
    return a[0] * b[0] + a[1] * b[1] + a[2] * b[2]
}

func (a Vector) Len() float64 {
    return math.Sqrt(a[0] * a[0] + a[1] * a[1] + a[2] * a[2])
}

func (a Vector) Dist(b Vector) float64 {
    dist := a.Sub(b)
    return math.Sqrt(dist[0] * dist[0] + dist[1] * dist[1] + dist[2] * dist[2])
}

func (a Vector) Norm() Vector {
    d := 1.0 / a.Len()
    return Vector{a[0] * d, a[1] * d, a[2] * d}
}

func (a Vector) Cross(b Vector) Vector {
    return Vector{a[1]*b[2] - a[2]*b[1], a[2]*b[0] - a[0]*b[2], a[0]*b[1] - a[1]*b[0]}
}
