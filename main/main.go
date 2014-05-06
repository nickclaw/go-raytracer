package main

import (
    "fmt"
    "github.com/nickclaw/raytrace"
    "os"
    "image/jpeg"
)

func main() {
    scene := raytrace.Scene{}
    camera := raytrace.OrthoCamera{
        Loc: raytrace.Point{5, 0, 0},
        Dir: raytrace.Vector{-1, 0, 0},
        Width: 4,
        Height: 4,
    }
    scene.ReadFile("/Users/nickclaw/Documents/Go/src/github.com/nickclaw/raytrace/resource/triangle.obj")
    image := scene.Render(camera, 100)

    file, _ := os.Create("/Users/nickclaw/Documents/Go/src/github.com/nickclaw/raytrace/resource/output.jpg")
    defer file.Close()

    jpeg.Encode(file, &image, nil)

    fmt.Println("Done.")
}
