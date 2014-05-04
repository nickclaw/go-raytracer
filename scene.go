package raytrace

import (
    "fmt"
    "strings"
    "strconv"
    "os"
    "bufio"
)

type Scene struct {
    Vertices []Point
    Faces []Face
    Lights []Point
}

func (s *Scene) Build(file string) bool {
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
            x, _ := strconv.ParseFloat(tokens[1], 64)
            y, _ := strconv.ParseFloat(tokens[2], 64)
            z, _ := strconv.ParseFloat(tokens[3], 64)
            s.Vertices = append(s.Vertices, Point{x,y,z})
        } else if tokens[0] == "f" {
            v1, _ := strconv.Atoi(tokens[1])
            v2, _ := strconv.Atoi(tokens[2])
            v3, _ := strconv.Atoi(tokens[3])
            s.Faces = append(s.Faces, Face{s.Vertices[v1 - 1], s.Vertices[v2 - 1], s.Vertices[v3 - 1]})
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Scanner had an error.", err)
        return false
    }

    return true
}
