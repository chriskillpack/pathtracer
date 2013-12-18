package main

import (
  "image"
  "image/png"
  "os"
  "pathtracer/vector"
)

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

  sphere := Object(Sphere{vector.Vector3{0,0,0}, 2})
  for x := 0; x < 256; x++ {
    for y := 0; y < 256; y++ {
      // Generate a ray
      dx := (float32(x - 128) / 128) * 5
      dy := (float32(y - 128) / 128) * 5
      ray_direction := vector.Normalize(vector.Vector3{dx, dy, 10})
      does_intersect, _ := sphere.Intersect(vector.Vector3{0,0,-10}, ray_direction)
      col := 0
      if does_intersect {
        col = 255
      }

      index := y * m.Stride + x * 4
      m.Pix[index+0] = uint8(col);
      m.Pix[index+1] = uint8(col);
      m.Pix[index+2] = uint8(col);
      m.Pix[index+3] = 255;
    }
  }
  png.Encode(img, m)
}
