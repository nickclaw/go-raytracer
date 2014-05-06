package raytrace

import (
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "os"
    "image"
    "image/color"
    "math"
)

type Scene struct {
    Faces []Face
    Lights []Light
}

/**
 * Reads a .obj file into the scene
 * @param {string} filename
 * @return {bool} success
 */
func (s *Scene) ReadFile(filename string) bool {
    var vertices []Point

    f, err := os.Open(filename)
    if err != nil {
        fmt.Println("File " + filename + " could not be found.", err)
        return false
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        tokens := strings.Split(scanner.Text(), " ")
        if tokens[0] == "v" {
            x, _ := strconv.ParseFloat(tokens[1], 64)
            y, _ := strconv.ParseFloat(tokens[3], 64)
            z, _ := strconv.ParseFloat(tokens[2], 64)
            vertices = append(vertices, NewPoint(x,y,z))
        } else if tokens[0] == "f" {
            v0, _ := strconv.Atoi(tokens[1])
            v1, _ := strconv.Atoi(tokens[2])
            v2, _ := strconv.Atoi(tokens[3])
            s.Faces = append(s.Faces, NewFace(
                vertices[v0 - 1],
                vertices[v1 - 1],
                vertices[v2 - 1]),
            )
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Scanner had an error.", err)
        return false
    }

    return true
}

/**
 * Renders the scene with the camera
 * @param {Viewer} camera
 * @param {flaot64} scale
 * @return {image}
 */
func (s Scene) Render(camera Viewer, scale float64) image.Gray16 {
    image := createImage(camera, scale)

    for _, ray := range camera.GetRays(scale) {
        pt, face, _, err := closestFace(ray, s.Faces)

        if err {continue}

        col := calculateShading(face, pt, s.Lights)
        image.SetGray16(ray.X, ray.Y, color.Gray16{col})

    }

    return image
}

/**
 * Creates an image with the right dimensions
 * @private
 * @param {Viewer} camera
 * @param {float64} scale
 * return {image}
 */
func createImage(camera Viewer, scale float64) image.Gray16 {
    x, y := camera.GetDimensions(scale)
    i := image.NewGray16(image.Rect(0,0,x,y))

    return *i
}

/**
 * Finds the closest face intersecting with the ray
 * @param {*Ray} r
 * @param {*[]Face} faces
 * @return {Point}
 * @return {Face}
 * @return {float64} distance
 * @return {bool} err
 */
func closestFace(r Ray, faces []Face) (Point, Face, float64, bool) {
    var minFace Face
    var minDist float64 = math.MaxFloat64
    var minPt Point

    for _, face := range faces {
        pt, dist, err := face.GetIntersection(r)

        if !err && dist < minDist {
            minFace, minDist, minPt = face, dist, pt
        }
    }

    if minDist == math.MaxFloat64 {
        return Point{}, Face{}, 0.0, true
    } else {
        return minPt, minFace, minDist, false
    }
}

/**
 * Calculates the appropriate shading at any point
 * @param {Face} face
 * @param {Point} pt
 * @param {[]Light} lights
 * @return {uint16}
 */
func calculateShading(face Face, pt Point, lights []Light) uint16 {
    var col uint16 = 0
    for _, light := range lights {
        n := face.GetNormal().Dot(pt.VectorTo(light.Loc))
        if n > 0.0 {
            col += uint16(n * 10000.0)
        }
    }
    return col
}
