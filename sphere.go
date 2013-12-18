package main

import (
  "math"
  "pathtracer/vector"
)

type Sphere struct {
  center vector.Vector3
  radius float32
}

func (sphere Sphere) Intersect(ray_origin, ray_direction vector.Vector3) float32 {
  /**
   * From: http://www.cs.umbc.edu/~olano/435f02/ray-sphere.html
   */
  dst := vector.Sub(ray_origin, sphere.center)

  a := vector.Dot(ray_direction, ray_direction)
  b := 2 * vector.Dot(ray_direction, dst)
  c := vector.Dot(dst, dst) - (sphere.radius * sphere.radius)
  discrim_sq := float64(b * b - 4 * a * c)
  if (discrim_sq < 0) {
    return math.MaxFloat32
  }

  discrim := float32(math.Sqrt(discrim_sq))
  t := float32(0)
  if (math.Abs(discrim_sq) > 1e-2) {
    t = (-b - discrim) / (2 * a)
  } else {
    t = -b / (2 * a)
  }
  return t
}
