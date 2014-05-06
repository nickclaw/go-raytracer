package raytrace

type Ray struct {
    X,Y int
    Loc Point
    Dir Vector
}

/**
 * @constructor
 */
func NewRay(x int, y int, loc Point, dir Vector) Ray {
    return Ray{
        X: x,
        Y: y,
        Loc: loc,
        Dir: dir.Norm(),
    }
}


/***** METHODS *****/

/**
 * Extends the ray in space by a factor
 * @param {float64}
 * @return {Point}
 */
func (r Ray) Move(t float64) Point {
    return r.Dir.ToPoint(r.Loc, t)
}
