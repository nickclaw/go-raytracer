package raytrace

// define an interface for all cameras
type Viewer interface {
    GetRays(scale float64) []Ray
    GetDimensions(scale float64) (int, int)
}

type OrthoCamera struct {
    Loc Point
    Dir Vector
    Width, Height float64 // actual, not rendered
}

/**
 * Gets the rays from the camera
 * @param {float64} scale
 * @return {[]Ray}
 */
func (c OrthoCamera) GetRays(scale float64) []Ray {
    rays := []Ray{}
    uOff, vOff := c.Width / 2.0, c.Height / 2.0
    x, y := c.GetDimensions(scale)

    for u := 0; u < x; u++ {
        for v := 0; v < y; v++ {
            rays = append(rays, Ray{
                X: u,
                Y: v,
                Loc: Point{
                    c.Loc[0] + (float64(u) / scale - uOff) * c.Dir[1],
                    c.Loc[1] + (float64(u) / scale - uOff) * c.Dir[0],
                    c.Loc[2] -(float64(v) / scale - vOff),
                },
                Dir: c.Dir,
            })
        }
    }

    return rays
}

/**
 * Returns the dimensions the camera will render
 * @param {float64} scale
 * @return {int} width px
 * @return {int} height px
 */
func (c OrthoCamera) GetDimensions(scale float64) (int, int) {
    return int(c.Width * scale), int(c.Height * scale)
}
