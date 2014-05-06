package raytrace

import "math"

type Point [3]float64

/**
 * @constructor
 */
func NewPoint(a,b,c float64) Point {
    return Point{a,b,c}
}

/**
 * Adds two points
 * @param {Point} b
 * @return {Point}
 */
func (a Point) Add(b Point) Point {
    return Point{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

/**
 * Subtracts two points
 * @param {Point} b
 * @return {Point}
 */
func (a Point) Sub(b Point) Point {
    return Point{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

/**
 * Gets the distance between two points
 * @param {Point} b
 * @return {float64}
 */
func (a Point) Distance(b Point) float64 {
    dist := a.Sub(b)
    return math.Sqrt(dist[0] * dist[0] + dist[1] * dist[1] + dist[2] * dist[2])
}

/**
 * Returns the unit vector from one point to another
 * @param {Point} b
 * @return {Vector} unit vector
 */
func (a Point) UnitVectorTo(b Point) Vector {
    return a.VectorTo(b).Norm()
}

/**
 * Returns the vector from one point to another
 * equivalent to Vector.Sub
 * @param {Point} b
 * @return {Vector} vector
 */
func (a Point) VectorTo(b Point) Vector {
    return Vector{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}
