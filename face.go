package raytrace

type Face [3]Vector

func (r Face) GetNormal() Vector {
    return BuildVector(r[1], r[0]).Cross(BuildVector(r[2], r[0])).Norm()
}

// based off http://www.cs.virginia.edu/~gfx/Courses/2003/ImageSynthesis/papers/Acceleration/Fast%20MinimumStorage%20RayTriangle%20Intersection.pdf
// unculled... because culled didn't work AT ALL
func (r Face) GetIntersection(l Ray) (Vector, float64, bool) {
    e1 := r[1].Sub(r[0])
    e2 := r[2].Sub(r[0])

    pVec := l.Dir.Cross(e2)

    det := e1.Dot(pVec)

    if det > -.000001 && det < .000001 {
        return Vector{}, 0.0, true
    }

    inv_det := 1.0 / det

    tVec := l.Loc.Sub(r[0])
    u := tVec.Dot(pVec) * inv_det
    if u < 0.0 || u > 1.0 {
        return Vector{}, 0.0, true
    }

    qVec := tVec.Cross(e1)
    v := l.Dir.Dot(qVec) * inv_det
    if v < 0.0 || u + v > 1.0 {
        return Vector{}, 0.0, true
    }

    t := e2.Dot(qVec) * inv_det


    return l.Dir.Mult(t).Add(l.Loc), t, false
}

func BuildVector(a,b Vector) Vector {
    return Vector{a[0]-b[0], a[1]-b[1], a[2]-b[2]}
}
