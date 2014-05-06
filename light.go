package raytrace

type Light struct {
    loc Point
    power float64
}

func NewLight(x,y,z,pow float64) Light {
    return Light{
        loc: Point{x,y,z},
        power: pow,
    }
}
