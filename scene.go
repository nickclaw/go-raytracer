package raytrace

import (
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "os"
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
func (s Scene) ReadFile(filename string) bool {
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

    fmt.Println(vertices)

    if err := scanner.Err(); err != nil {
        fmt.Println("Scanner had an error.", err)
        return false
    }

    return true
}

func (s Scene) Render() bool {
    return true
}
