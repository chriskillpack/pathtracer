package main

import (
  "math"
  "pathtracer/vector"
)

type Sphere struct {
  center vector.Vector3
  radius float32
}

func (sphere Sphere) Intersect(rayOrigin, rayDirection vector.Vector3) (doesIntersect bool, distance float32) {
  /**
   * From: http://www.cs.umbc.edu/~olano/435f02/ray-sphere.html
   */
  dst := vector.Sub(rayOrigin, sphere.center)

  a := vector.Dot(rayDirection, rayDirection)
  b := 2 * vector.Dot(rayDirection, dst)
  c := vector.Dot(dst, dst) - (sphere.radius * sphere.radius)
  discrim_sq := float64(b * b - 4 * a * c)
  if (discrim_sq < 0) {
    return false, 0
  }

  discrim := float32(math.Sqrt(discrim_sq))
  t := float32(0)
  if (math.Abs(discrim_sq) > 1e-2) {
    t = (-b - discrim) / (2 * a)
  } else {
    t = -b / (2 * a)
  }
  return true, t
}
