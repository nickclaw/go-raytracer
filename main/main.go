package main

import (
    "fmt"
    "github.com/nickclaw/raytrace"
)

func main() {
    scene := raytrace.Scene{}
    scene.Build("/Users/nickclaw/Documents/Go/src/github.com/nickclaw/raytrace/resource/scene.obj")
    fmt.Println(scene)
}
