package main

import (
    "fmt"
    "github.com/nickclaw/raytrace"
)

func main() {
    scene := raytrace.Scene{
        Lights: []raytrace.Point{raytrace.Point{5, 5, 2}},
        Camera: raytrace.OrthoCamera{
            Loc: raytrace.Point{5, 0, 0},
            Dir: raytrace.Vector{-1, 0, 0},
        },
    }

    scene.Build("/Users/nickclaw/Documents/Go/src/github.com/nickclaw/raytrace/resource/scene.obj")
    fmt.Println(scene)
}
