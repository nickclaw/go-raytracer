package raytrace

import (
    "fmt"
    "strings"
    "strconv"
    "os"
    "bufio"
    "image"
    "image/color"
    "math"
)

/**
 *  Holds all objects in the scene
 */
type Scene struct {
    Vertices []Vector
    Faces []Face
    Lights []Vector
    Camera Viewer
}

/**
 * Reads a .obj file into the scene
 * Can be called multiple times to read in multiple .obj files
 * @param {string} file
 * @return {bool} true if successful
 */
func (s *Scene) Build(file string) bool {
    vOffset := len(s.Vertices) - 1

    // try to open file
    f, err := os.Open(file)
    if err != nil {
        fmt.Println("File " + file + " could not be opened.", err)
        return false
    }
    defer f.Close()

    // scan the file
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        tokens := strings.Split(scanner.Text(), " ")
        if tokens[0] == "v" {
            x, _ := strconv.ParseFloat(tokens[3], 64)
            y, _ := strconv.ParseFloat(tokens[1], 64)
            z, _ := strconv.ParseFloat(tokens[2], 64)
            z *= -1
            s.Vertices = append(s.Vertices, Vector{x,y,z})
            fmt.Println(x,y,z)
        } else if tokens[0] == "f" {
            v1, _ := strconv.Atoi(tokens[3])
            v2, _ := strconv.Atoi(tokens[2])
            v3, _ := strconv.Atoi(tokens[1])
            s.Faces = append(s.Faces, Face{s.Vertices[v1 + vOffset], s.Vertices[v2 + vOffset], s.Vertices[v3 + vOffset]})
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Scanner had an error.", err)
        return false
    }

    return true
}

func (s Scene) Render(image *image.Gray16, scale int) bool {
    // for each pixel of the camera
    for _, ray := range s.Camera.GetRays(scale) {
    //    fmt.Println(ray)
        var col uint16 = 0 // default pixel color

        // for getting closest face
        // var minFace Face
        var minDist float64 = math.MaxFloat64

        // get closest intersecting face
        for _, face := range s.Faces {
            _, dist, err := face.GetIntersection(ray)

            if !err && dist < minDist {
                // minFace = face
                minDist = dist
            }
        }

        // was intersection? render!
        if minDist < math.MaxFloat64 {
            col = 65535 - uint16(65535 * minDist / 1.8)
            image.SetGray16(ray.X, ray.Y, color.Gray16{col})
        }

    }


    return true
}
