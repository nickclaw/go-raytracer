package raytrace

import "fmt"

// define an interface for all cameras
type Viewer interface {
    GetRays(scale int) []Ray
}

type OrthoCamera struct {
    Loc Point
    Dir Vector
    Width, Height float64 // actual, not rendered
}

func (c OrthoCamera) GetRays(scale float64) []Ray {
    rays := []Ray{}

    uOff, vOff := c.Width / 2.0, c.Heigth / 2.0

    for u := 0; u < int(c.Width * scale); u++ {
        for v := 0; v < int(c.Height * scale); v++ {
            rays = append(rays, Ray{
                X: u,
                Y: v,
                Loc: Point{
                    c.Loc[0] + (float64(u) / scale - uOff) * c.Dir[1],
                    c.Loc[1] + (float64(u) / scale - uOff) * c.Dir[0],
                    c.Loc[2] + (float64(v) / scale + vOff),
                },
                Dir: c.Dir,
            })
        }
    }

    return rays
}
