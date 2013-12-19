package main

import (
  "fmt"
  "image"
  "image/png"
  "os"
  "pathtracer/vector"
)

var _ = fmt.Println

func main() {
  /* TODO: Move these into a test file */
  /* v := Vector3{1,2,3} */
  /* fmt.Println(v.Len()) */
  /* v2 := Vector3{0, 1, 0} */
  /* fmt.Println(Dot(v, v2)) */
  /* v3 := Vector3{1, 0, 0} */
  /* fmt.Println(Cross(v2, v3)) */

  m := image.NewRGBA(image.Rect(0, 0, 256, 256))
  img,_ := os.Create("foo.png")
  defer img.Close()

  sphere := SceneObject(Sphere{vector.Vector3{0,0,0}, 2})

  for x := 0; x < 256; x++ {
    for y := 0; y < 256; y++ {
      // Generate a ray
      dx := (float32(x - 128) / 128) * 5
      dy := (float32(y - 128) / 128) * 5
      rayDirection := vector.Normalize(vector.Vector3{dx, dy, 10})
      intersection := sphere.Intersect(vector.Vector3{0,0,-10}, rayDirection)
      var r, g, b uint8
      if intersection.doesIntersect {
        r, g, b = normalToColor(intersection.normal)
      }

      index := y * m.Stride + x * 4
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
