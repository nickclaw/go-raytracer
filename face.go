package raytrace

type Face [3]Vector

func (r Face) GetNormal() Vector {
    return BuildVector(r[1], r[0]).Cross(BuildVector(r[2], r[0])).Norm()
}

func (r Face) GetIntersection(l Ray) (Vector, float64, bool) {

    N := r.GetNormal()

    /** Find P **/

    // check if ray and plane are parralel
    angle := N.Dot(l.Dir)
    if (angle == 0) {
        return Vector{}, 0.0, true
    }

    // get d?
    d := N.Dot(r[0])
    t := -(N.Dot(l.Loc) + d) / angle

    // check if triangle is behind ray
    if (t < 0) {
        return Vector{}, 0.0, true
    }

    // ray hits triangle plane here
    p := l.Dir.Mult(t).Add(l.Loc)

    // doe point land within triangle?
    e0 := r[1].Sub(r[0])
    vp0 := p.Sub(r[0])
    if (N.Dot(e0.Cross(vp0)) < 0) {
        return Vector{}, 0.0, true
    }

    e1 := r[2].Sub(r[1])
    vp1 := p.Sub(r[1])
    if (N.Dot(e1.Cross(vp1)) < 0) {
        return Vector{}, 0.0, true
    }

    e2 := r[0].Sub(r[2])
    vp2 := p.Sub(r[2])
    if (N.Dot(e2.Cross(vp2)) < 0) {
        return Vector{}, 0.0, true
    }

    // yay
    return p, t, false

}

func BuildVector(a,b Vector) Vector {
    return Vector{a[0]-b[0], a[1]-b[1], a[2]-b[2]}
}
