package raytrace

// define an interface for all cameras
type Viewer interface {
    GetRays(scale int) []Ray
}

type OrthoCamera struct {
    Loc Point
    Dir Vector
    width, height int // actual, not rendered
}

func (c OrthoCamera) GetRays(scale int) []Ray {
    rays := []Ray{}
    for u := 0; u < scale * c.width; u++ {
        for v := 0; v < scale * c.height; v++ {
            rays = append(rays, Ray{Loc: Point{c.Dir[1] * float64(u), c.Dir[0] * float64(u) , float64(v)}, Dir: c.Dir})
        }
    }
    return rays
}
