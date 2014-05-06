package main

import "math"

type Point [3]float64

/**
 * @constructor
 */
func NewPoint(a,b,c) Point {
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
 * @return {float64} t
 */
func (a Point) VectorTo(b Point) (Vector, float64) {
    vec := Vector{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
    return vec.Norm(), vec.Len()
}
