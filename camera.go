package raytrace

import "fmt"

// define an interface for all cameras
type Viewer interface {
    GetRays(scale int) []Ray
}

type OrthoCamera struct {
    Loc Vector
    Dir Vector
    Width, Height int // actual, not rendered
}

func (c OrthoCamera) GetRays(scale int) []Ray {
    rays := []Ray{}
    xMult, yMult := 1.0 / float64(scale), 1.0 / float64(scale)
    xOff, yOff := -float64(c.Width) / 2, -float64(c.Height) / 2

    fmt.Println("mults", xMult, yMult)
    fmt.Println("offs", xOff, yOff)
    fmt.Println("uvs", scale * c.Width, scale * c.Height)

    for u := 0; u < scale * c.Width; u++ {
        for v := 0; v < scale * c.Height; v++ {
            rays = append(rays, Ray{
                X : u,
                Y : v,
                Loc: Vector{
                    c.Loc[0] + (float64(u) * xMult + xOff) * c.Dir[1],
                    c.Loc[1] + (float64(u) * yMult + xOff) * c.Dir[0],
                    c.Loc[2] + (float64(v) * yMult + yOff),
                },
                Dir: c.Dir,
            })
        }
    }

    return rays
}
