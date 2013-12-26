package main

import (
  "fmt"
  "image"
  "image/png"
  "math"
  "math/rand"
  "os"
  "pathtracer/vector"
)

var _ = fmt.Println

const (
  ImageWidth int = 256
  ImageHeight = 256
)

var (
  WorldUp = vector.Vector3{0,1,0}

  EyePosition = vector.Vector3{0,15,-10}
  EyeLookAtTarget = vector.Vector3{0,15,0}
  EyePlaneDist = float32(1.0)
)

func main() {
  m := image.NewRGBA(image.Rect(0, 0, ImageWidth, ImageHeight))
  img,_ := os.Create("foo.png")
  defer img.Close()

  sceneObjects := []SceneObject{
    Sphere{vector.Vector3{2,10,2}, 2.5},
    Plane{vector.Vector3{-1,0,0}, 10}, // Left
    Plane{vector.Vector3{0,0,1}, 20},  // Back
    Plane{vector.Vector3{1,0,0}, 10}, // Right
    Plane{vector.Vector3{0,1,0}, 0},  // Bottom
    Plane{vector.Vector3{0,-1,0}, -30},  // Top
  }
  // sceneObjects = append(sceneObjects, foo)

  // Compute the eye-space to world-space transformation basis vectors. We use
  // these basis vectors to generate target points for ray generation. The
  // whole process is a little long-winded but working up from first principles.
  lookVector := vector.Sub(EyeLookAtTarget, EyePosition).Normalize()
  xAxis, yAxis, _ := coordinateFrame(lookVector)

  for i := 0; i < ImageHeight; i++ {
    for j := 0; j < ImageWidth; j++ {
      // Compute the 2D position of this pixel on the view plane. The view plane
      // is centered around the view vector and extends [-1,1] in both axis.
      npx := float32(j - ImageWidth/2) / float32(ImageWidth/2)
      npy := float32(ImageHeight/2 - i) / float32(ImageHeight/2)

      // Scale the world-space basis vectors of the view plane to generate
      // offset vectors from the center point of the view plane.
      xAxisOffset := vector.Scale(xAxis, npx)
      yAxisOffset := vector.Scale(yAxis, npy)
      // Move down the view vector from the eye to generate the center point
      // on the view plane in world-space.
      midPoint := vector.Add(EyePosition, vector.Scale(lookVector, EyePlaneDist))

      // Add the offset vectors on to the center point to generate the ray
      // target in world-space.
      target := vector.Add(vector.Add(midPoint, xAxisOffset), yAxisOffset)

      // Finally compute a ray direction from the eye to the target.
      rayDirection := vector.Sub(target, EyePosition).Normalize()

      ray := Ray{EyePosition, rayDirection}
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

// Generate a random direction on a hemisphere oriented around the input normal.
// The random directions have a cosine-weighted distribution.
// See formula 35 in http://people.cs.kuleuven.be/~philip.dutre/GI/TotalCompendium.pdf
func generateHemisphereDirection(normal vector.Vector3) vector.Vector3 {
  r1 := 2 * math.Pi * rand.Float64()
  r2 := rand.Float64()
  r2s := math.Sqrt(1-r2)

  // Compute a direction in the unit hemisphere
  x := float32(math.Cos(r1) * r2s)
  y := float32(math.Sin(r1) * r2s)
  z := float32(math.Sqrt(r2))

  // Compute a coordinate frame around the normal and then linearly combine
  // the basis vectors using the direction components as weights to transform
  // from unit hemisphere coordinate system to the normal's coordinate system.
  u, v, w := coordinateFrame(normal)
  t := vector.Add(vector.Scale(u, x), vector.Scale(v, y))
  return vector.Add(t, vector.Scale(w, z))
}

// Given an input vector n this function generates a coordinate frame with the
// input vector as the Z axis.
func coordinateFrame(n vector.Vector3) (u, v, w vector.Vector3) {
  if (math.Abs(float64(vector.Dot(n, WorldUp))) > 0.9) {
    // Input too close to world up, use another world up.
    u = vector.Cross(vector.Vector3{0,0,-1}, n)
  } else {
    u = vector.Cross(WorldUp, n).Normalize()
  }

  v = vector.Cross(n, u).Normalize()
  w = vector.Normalize(n)

  return
}
