package raytrace

type Face struct{
    V0, V1, V2 Point
    EV1, EV2 Vector
}

/**
 * @constructor
 */
func NewFace(v0,v1,v2 Point) Face {
    return Face{
        V0: v0,
        V1: v1,
        V2: v2,
        EV1: v1.VectorTo(v0),
        EV2: v2.VectorTo(v0),
    }
}

/**
 * Finds the intersection between a Face and a Ray
 * based off of http://www.cs.virginia.edu/~gfx/Courses/2003/ImageSynthesis/papers/Acceleration/Fast
 * @param {Ray} r
 * @return {Point} intersection
 * @return {float64} distance
 * @return {bool} was error
 */
func (f Face) GetIntersection(r Ray) (Point, float64, bool) {
    pVec := r.Dir.Cross(f.EV2)
    det := f.EV2.Dot(pVec)

    if det > -.000001 && det < .000001 {
        return Point{}, 0.0, true
    }

    inv_det := 1.0 / det
    tVec := r.Loc.VectorTo(f.V0)

    u := tVec.Dot(pVec) * inv_det
    if u < 0.0 || u > 1.0 {
        return Point{}, 0.0, true
    }

    qVec := tVec.Cross(f.EV1)
    v := r.Dir.Dot(qVec)
    if v < 0.0 || u + v > 1.0 {
        return Point{}, 0.0, true
    }

    t := f.EV2.Dot(qVec)

    return r.Move(t), t, false
}
