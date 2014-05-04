package main

import (
    "fmt"
    "github.com/nickclaw/raytrace"
)

func main() {
    face := raytrace.Face{raytrace.Point{0,0,0}, raytrace.Point{1,1,1}, raytrace.Point{2,2,2}}
    fmt.Println(face)
}
