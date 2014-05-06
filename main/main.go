package main

import (
    "fmt"
    "github.com/nickclaw/raytrace"
)

func main() {
    scene := raytrace.Scene{}
    scene.ReadFile("/Users/nickclaw/Documents/Go/src/github.com/nickclaw/raytrace/resource/triangle.obj")
    fmt.Println("Done.")
}
