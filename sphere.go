package main

import (
  "math"
  "pathtracer/vector"
)

type Sphere struct {
  center vector.Vector3
  radius float32
}

// Compute the intersection between a sphere and a ray.
// From: http://www.cs.umbc.edu/~olano/435f02/ray-sphere.html
func (sphere Sphere) Intersect(rayOrigin, rayDirection vector.Vector3) Intersection {
  dst := vector.Sub(rayOrigin, sphere.center)

  a := vector.Dot(rayDirection, rayDirection)
  b := 2 * vector.Dot(rayDirection, dst)
  c := vector.Dot(dst, dst) - (sphere.radius * sphere.radius)
  discrim_sq := float64(b * b - 4 * a * c)
  if (discrim_sq < 0) {
    return Intersection{false, 0, vector.Vector3{}}
  }

  discrim := float32(math.Sqrt(discrim_sq))
  t := float32(0)
  if (math.Abs(discrim_sq) > 1e-2) {
    t = (-b - discrim) / (2 * a)
  } else {
    t = -b / (2 * a)
  }

  normal := sphere.computeNormal(rayOrigin, rayDirection, t)
  return Intersection{true, t, normal}
}

// Compute the normal at the point of intersection.
func (sphere Sphere) computeNormal(rayOrigin, rayDirection vector.Vector3, t float32) vector.Vector3 {
  pointOnSphere := vector.Add(rayOrigin, rayDirection.Scale(t))
  v := vector.Sub(pointOnSphere, sphere.center)
  return vector.Normalize(v)
}
