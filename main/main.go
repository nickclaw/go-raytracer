package main

import (
    "fmt"
    "os"
    "github.com/nickclaw/raytrace"
    "image"
    "image/jpeg"
)

func main() {
    width, height, scale := 3, 3, 400
    image := image.NewGray16(image.Rect(0,0,width*scale, height*scale))

    scene := raytrace.Scene{
        Camera: raytrace.OrthoCamera{
            Loc: raytrace.Vector{5,0,0},
            Dir: raytrace.Vector{-1,0,0},
            Width: width,
            Height: height,
        },
        Lights: []raytrace.Vector{
            raytrace.Vector{-10,2,1},
            raytrace.Vector{-10,-5, -3},
        },
    }
    scene.Build("/Users/nickclaw/Documents/Go/src/github.com/nickclaw/raytrace/resource/both.obj")
    scene.Render(image, scale)

    output, _ := os.Create("/Users/nickclaw/Documents/Go/src/github.com/nickclaw/raytrace/resource/output.jpg")
    defer output.Close()

    jpeg.Encode(output, image, nil)
    fmt.Println("Done.")
}
