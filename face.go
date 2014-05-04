package raytrace

type Face [3]Point

func (r Face) GetNormal() Vector {
    return BuildVector(r[1], r[0]).Cross(BuildVector(r[2], r[0])).Norm()
}

// made from http://www.lighthouse3d.com/tutorials/maths/ray-triangle-intersection/
func (r Face) GetIntersection(l Ray) (Point, float64, bool){
    e1 := BuildVector(r[1], r[0]) // triangle edge 1
    e2 := BuildVector(r[2], r[0]) // triangle edge 2
    h := l.Dir.Cross(e2)
    a := e1.Dot(h)

    if a == 0.0 {
        return Point{}, 0.0, true
    }

    f := 1 / a
    s := BuildVector(l.Loc, r[0])
    u := f * s.Dot(h)

    if u < 0.0 || u > 1.0 {
        return Point{}, 0.0, true
    }

    q := s.Cross(e1)
    v := f * l.Dir.Dot(q)

    if v < 0.0 || u + v > 1.0 {
        return Point{}, 0.0, true
    }

    t := f * e2.Dot(q)

    if (t > 0) {
        return Point{l.Loc[0] + t * l.Dir[0], l.Loc[1] + t * l.Dir[1], l.Loc[2] + t * l.Dir[2]}, t, false
    } else {
        return Point{}, t, true
    }
}

func BuildVector(a,b Point) Vector {
    return Vector{a[0]-b[0], a[1]-b[1], a[2]-b[2]}
}
