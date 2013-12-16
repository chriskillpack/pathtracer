package main

import (
  /* "fmt" */
  "image"
  "image/png"
  "os"
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

  for x := 0; x < 256; x++ {
    for y := 0; y < 256; y++ {
      index := y * m.Stride + x * 4
      col := x ^ y
      m.Pix[index+0] = uint8(col);
      m.Pix[index+1] = uint8(col);
      m.Pix[index+2] = uint8(col);
      m.Pix[index+3] = 255;
    }
  }
  png.Encode(img, m)
}
