package main

import (
  "fmt"
  "image"
  "image/png"
  "os"
  "pathtracer/vector"
)

var _ = fmt.Println

const (
  ImageWidth int = 256
  ImageHeight = 256
)

func main() {
  m := image.NewRGBA(image.Rect(0, 0, 256, 256))
  img,_ := os.Create("foo.png")
  defer img.Close()

  sceneObjects := []SceneObject{
    Sphere{vector.Vector3{0,0,0}, 2},
    Plane{vector.Vector3{0,1,0}, -3},
  }
  // sceneObjects = append(sceneObjects, foo)

  for i := 0; i < ImageHeight; i++ {
    for j := 0; j < ImageWidth; j++ {
      // Generate a ray
      npx := float32(j - ImageWidth/2) / float32(ImageWidth/2)
      npy := float32(ImageHeight/2 - i) / float32(ImageHeight/2)
      rayDirection := vector.Normalize(vector.Vector3{npx * 5, npy * 5, 10})

      ray := Ray{vector.Vector3{0,0,-10}, rayDirection}
      var closestIntersection Intersection = DefaultIntersection
      for _, object := range sceneObjects {
        intersection := object.Intersect(ray)
        if intersection.doesIntersect {
          if intersection.distance < closestIntersection.distance {
            closestIntersection = intersection
          }
        }
      }

      var r, g, b uint8
      if closestIntersection.doesIntersect {
        r, g, b = normalToColor(closestIntersection.normal)
      }

      index := i * m.Stride + j * 4
      m.Pix[index+0] = r
      m.Pix[index+1] = g
      m.Pix[index+2] = b
      m.Pix[index+3] = 255;
    }
  }
  png.Encode(img, m)
}

// Convert a normal into a color. Normal components are in the range [-1,1] and
// are converted to colors in the range [0,254].
func normalToColor(normal vector.Vector3) (r, g, b uint8) {
  r = colorToUint8((normal.X + 1) * 127)
  g = colorToUint8((normal.Y + 1) * 127)
  b = colorToUint8((normal.Z + 1) * 127)
  return
}

// Convert a floating point value to a clamped value that fits in the [0,255]
// range of a uint8.
func colorToUint8(x float32) uint8 {
  if (x > 255) {
    return 255
  } else if (x < 0) {
    return 0
  } else {
    return uint8(x)
  }
}
