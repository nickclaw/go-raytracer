package raytrace

type Face [3]Vector

func (r Face) GetNormal() Vector {
    return BuildVector(r[1], r[0]).Cross(BuildVector(r[2], r[0])).Norm()
}

// based off http://www.cs.virginia.edu/~gfx/Courses/2003/ImageSynthesis/papers/Acceleration/Fast%20MinimumStorage%20RayTriangle%20Intersection.pdf
func (r Face) GetIntersection(l Ray) (Vector, float64, bool) {
    e1 := r[1].Sub(r[0])
    e2 := r[2].Sub(r[0])

    pVec := l.Dir.Cross(e2)

    det := e1.Dot(pVec)

    if det < 0 {
        return Vector{}, 0.0, true
    }

    tVec := l.Loc.Sub(r[0])
    qVec := tVec.Cross(e1)

    u := tVec.Dot(pVec)
    v := l.Dir.Dot(qVec)
    t := e2.Dot(qVec)

    if (u < 0.0 || u > det || v < 0.0 || u + v > det) {
        return Vector{}, 0.0, true
    }

    return Vector{
        u / det,
        v / det,
        t / det,
    }, 1.0, false
}

func BuildVector(a,b Vector) Vector {
    return Vector{a[0]-b[0], a[1]-b[1], a[2]-b[2]}
}
