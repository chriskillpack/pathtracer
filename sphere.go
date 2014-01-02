package main

import (
  "math"
  "pathtracer/vector"
)

type Sphere struct {
  center vector.Vector3
  radius float32

  material Material
}

// Compute the intersection between a sphere and a ray.
// From: http://www.cs.umbc.edu/~olano/435f02/ray-sphere.html
func (sphere Sphere) Intersect(ray Ray) Intersection {
  dst := vector.Sub(ray.origin, sphere.center)

  a := vector.Dot(ray.direction, ray.direction)
  b := 2 * vector.Dot(ray.direction, dst)
  c := vector.Dot(dst, dst) - (sphere.radius * sphere.radius)
  discrim_sq := float64(b * b - 4 * a * c)
  if (discrim_sq < 0) {
    return Intersection{false, 0, vector.Vector3{}}
  }

  discrim := float32(math.Sqrt(discrim_sq))
  var t float32
  if (math.Abs(discrim_sq) > 1e-2) {
    t = (-b - discrim) / (2 * a)
  } else {
    t = -b / (2 * a)
  }

  normal := sphere.computeNormal(ray, t)
  return Intersection{true, t, normal}
}

func (sphere Sphere) Material() Material {
  return sphere.material
}

// Compute the normal at the point of intersection.
func (sphere Sphere) computeNormal(ray Ray, t float32) vector.Vector3 {
  pointOnSphere := vector.Add(ray.origin, vector.Scale(ray.direction, t))
  v := vector.Sub(pointOnSphere, sphere.center)
  return vector.Normalize(v)
}
