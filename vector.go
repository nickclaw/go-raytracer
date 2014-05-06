package raytrace

import "math"

type Vector [3]float64

/**
 * @constructor
 */
func NewVector(a,b,c float64) {
    return Vector{a,b,c}
}

/***** METHODS *****/

/**
 * Returns length of vector
 * @param {float64}
 */
func (a Vector) Length() float64 {
    return math.sqrt(a[0]*a[0] + a[1]*a[1] + a[2]*a[2])
}

/**
 * Creates unit vector
 * @return {Vector}
 */
func (a Vector) Norm() Vector {
    len := 1.0 / a.Length()
    return Vector{a[0]*len, a[1]*len, a[2]*len}
}

/**
 * Multiplies a vector by a constant
 * @param {float64} c
 * @return {Vector}
 */
func (a Vector) Mult(c float64) Vector {
    return Vector{a[0] * c, a[1] * c, a[2] * c}
}

/**
 * Dot prodcuts
 * @param {Vector} b
 * @return {float64}
 */
func (a Vector) Dot(b Vector) float64 {
    return a[0] * b[0] + a[1] * b[1] + a[2] * b[2]
}

/**
 * Returns cross product of two vectors
 * @param {Vector} b
 * @return {Vector}
 */
func (a Vector) Cross(b Vector) Vector {
    return Vector{a[1]*b[2] - a[2]*b[1], a[2]*b[0] - a[0]*b[2], a[0]*b[1] - a[1]*b[0]}
}

/**
 * Checks if one vector is parallel to the other
 * @param {Vector} b
 * @return {bool}
 */
func (a Vector) IsOrthoganal(b Vector) bool {
    return a.Dot(b) == 0
}

/**
 * Does p + v * t
 * @param {Point} origin
 * @param {float64} factor
 * @return {Point}
 */
func (a Vector) ToPoint(b Point, t float64) Point {
    mult := a.Mult(t)
    return Point{b[0] + mult[0], b[1] + mult[1], b[2] + mult[2]}
}
