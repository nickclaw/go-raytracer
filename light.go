package raytrace

type Light struct {
    Loc Point
    Power float64
}

func NewLight(x,y,z,pow float64) Light {
    return Light{
        Loc: Point{x,y,z},
        Power: pow,
    }
}
